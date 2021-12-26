# Postoffice

<!-- GETTING STARTED -->
To get a local copy of this project up and running follow these simple steps.


### Technology used
- [X] Go or golang
- [X] Docker
- [X] RabbitMQ
- [X] PostgreSQL

### Uses
This Project has two main part
- [x] Server
    - [x] RabbitMQ Publisher
    - [x] Cron Job Scheduler
- [x] Worker
    - [x] RabbitMQ Consumer
    - [x] Stores Object on PostgreSQL

### Environment Setup
1. In the source code you should get a file named `.example.env`.
2. Copy everything from that file and past it by creating a file called `.env`

If you need any changes change accordingly.

### Run Docker
1. In `docker-compose.yml` file, two services are `rabbitmq` and `postgres`
2. To run them follow the command:
   ```sh
   $ docker-compose up -d
   ```
3. Wait a bit. Now you should have both `rabbitmq` and `postgres` is ready.

### Dependency
1. To trac dependency create a `go.mod` file using the command below:
   ```sh
   $ go mod init 
   ```
2. If you need to modify vendored packages use the command below:
   ```sh
   $ go mod vendor 
   ```
3. Use the command below to cleans up unused dependencies or adds missing dependencies:
   ```sh
   $ go mod tidy 
   ```
### Available Commands
1. To see all available commands use:
   ```sh
   $ go run main.go
   ```

### Migration
1. Use the command for migration
   ```sh
   $ go run main.go migrate
   ```

### Server
1. Server accepts post request at `/callback` endpoint from `tester_server` (`Note`: necessary information about `tester_server` is given in the last part of this readme file.) and then It publishes them to RabbitMQ.
2. A cron job scheduler is running which deletes records which are not seen in last 30s.
3. To run use this command
   ```sh
   $ go run main.go serve 
   ```
### Worker
1. Worker will consume message form RabbitMQ and call a get request to `tester_server` using `tester_client` to get object details.
2. It stores those details to postgresql database.
3. To run the worker use the command:
   ```sh
   $ go run main.go worker "<n>"
   ```
Here `n` = `number of worker(s)` you want to consume message from the RabbitMQ. Just replace the `<n>` from the command.


### Tester_Server

Write a rest-service that listens on localhost:9090 for POST requests on /callback. Run the go service attached to this task. It will send requests to your service
at a fixed interval of 5 seconds. The request body will look like this:
`{
    "object_ids": [1,2,3,4,5,6]
}`
The amount of IDs varies with each request. Expect up to 200 IDs.
Every ID is linked to an object whose details can be fetched from the provided service. Our service listens on localhost:9010/objects/:id and returns the following response:
`{
    "id": <id>,
    "online": true|false 
}`
Note that this endpoint has an unpredictable response time between 300ms and 4s!
Your task is to request the object information for every incoming object_id and filter the objects by their "online" status.

Store all objects in a PostgreSQL database along with a timestamp when the object was last seen.
Let your service delete objects in the database when they have not been received for more than 30 seconds.

**Important**: due to business constraints we are not allowed to miss any callback to our service. Write code in such a way that all errors are properly recovered
and that the endpoint is always available. Optimize for very high throughput
so that this service could work in production.


#### Source Code of `tester_server.go`

`package main

import (
    "bytes"
    "fmt"
    "log"
    "math/rand"
    "net/http"
    "os"
    "os/signal"
    "strconv"
    "strings"
    "time"
)

func main() {
    rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	go func() {
		client := &http.Client{Timeout: 1 * time.Second}

		for {
			time.Sleep(5 * time.Second)

			ids := make([]string, rng.Int31n(200))
			for i := range ids {
				ids[i] = strconv.Itoa(rng.Int() % 100)
			}
			body := bytes.NewBuffer([]byte(fmt.Sprintf(`{"object_ids":[%s]}`, strings.Join(ids, ","))))
			resp, err := client.Post("http://localhost:9090/callback", "application/json", body)
			if err != nil {
				fmt.Println(err)
				continue
			}
			_ = resp.Body.Close()
		}
	}()

	http.HandleFunc("/objects/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Duration(rng.Int63n(4000)+300) * time.Millisecond)

		idRaw := strings.TrimPrefix(r.URL.Path, "/objects/")
		log.Println(idRaw)
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		w.Write([]byte(fmt.Sprintf(`{"id":%d,"online":%v}`, id, id%2 == 0)))
	})
	go func() { _ = http.ListenAndServe(":9010", nil) }()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig

	fmt.Println("closing")
}
`

<p align="right">(<a href="#top">back to top</a>)</p>