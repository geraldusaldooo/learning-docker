# Running multiple containers using Docker Compose


## Preparing database
![Database](https://github.com/geraldusaldooo/learning-docker/blob/b6a0a0a4e837683eed045fb25c45c96270b45da2/my-docker-compose/wise-word-compose/01-created-new-db.png)



## Create existing project that serve http and connect to postgres database.

Adding path '/' and '/ping' as a response
![Path](https://github.com/geraldusaldooo/learning-docker/blob/b6a0a0a4e837683eed045fb25c45c96270b45da2/my-docker-compose/wise-word-compose/02-connecting-paths.png)



## Creating new docker file
![Dockerfile](https://github.com/geraldusaldooo/learning-docker/blob/b6a0a0a4e837683eed045fb25c45c96270b45da2/my-docker-compose/wise-word-compose/03-created-dockerfile.png)



## Creating docker-compose.yml
![Docker Compose](https://github.com/geraldusaldooo/learning-docker/blob/b6a0a0a4e837683eed045fb25c45c96270b45da2/my-docker-compose/wise-word-compose/04-created-docker-compose-database.png)

![Docker Compose 2](https://github.com/geraldusaldooo/learning-docker/blob/b6a0a0a4e837683eed045fb25c45c96270b45da2/my-docker-compose/wise-word-compose/04-created-docker-compose-go.png)



## Response local host
![/](https://github.com/geraldusaldooo/learning-docker/blob/b6a0a0a4e837683eed045fb25c45c96270b45da2/my-docker-compose/wise-word-compose/05-response-localhost-9911.png)

![/ping](https://github.com/geraldusaldooo/learning-docker/blob/b6a0a0a4e837683eed045fb25c45c96270b45da2/my-docker-compose/wise-word-compose/06-response-localhost-ping.png)



## Connection log
![Log](https://github.com/geraldusaldooo/learning-docker/blob/b6a0a0a4e837683eed045fb25c45c96270b45da2/my-docker-compose/wise-word-compose/07-connection-log.png)