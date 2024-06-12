# Project 2: Running Compose V2

This project aims to run docker compose.


### Creating database server
![Server](https://github.com/geraldusaldooo/learning-docker/blob/main/project-my-docker-compose/Image/18-database-created.png)


### Creating golang project

Creating a golang project that serves http containing:
- Path '/ping' that returns "Pong" if successful and "Ouch!" if failed
- Print console log "Ping successful" if successful, "Ping failed" if failed
- Record data in database
- Stars on port 77

![Golang project](https://github.com/geraldusaldooo/learning-docker/blob/main/project-my-docker-compose/Image/19-create-golang-for-ping.png)


### Creating a docker file
![Dockerfile](https://github.com/geraldusaldooo/learning-docker/blob/main/project-my-docker-compose/Image/20-dockerfile-compose.png)


### Build Docker Compose file
![Docker Compose](https://github.com/geraldusaldooo/learning-docker/blob/main/project-my-docker-compose/Image/21-docker-compose.png)


### Building compose
![Build compose](https://github.com/geraldusaldooo/learning-docker/blob/main/project-my-docker-compose/Image/22-build-compose.png)


### Log Compose
![Log Compose](https://github.com/geraldusaldooo/learning-docker/blob/main/project-my-docker-compose/Image/23-log-compose.png)


### Checking connection to localhost

It prints the text I have written in the Golang file.

![Localhost response](https://github.com/geraldusaldooo/learning-docker/blob/main/project-my-docker-compose/Image/24-request-localhost.png)


### Checking connection of /ping
![/ping response](https://github.com/geraldusaldooo/learning-docker/blob/main/project-my-docker-compose/Image/25-request-ping.png)

