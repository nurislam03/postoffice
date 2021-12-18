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
1. Server accepts post request at `/callback` endpoint from `tester_server` and then It publishes them to RabbitMQ.
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



<p align="right">(<a href="#top">back to top</a>)</p>