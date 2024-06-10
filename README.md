# learning-docker

## Learning Docker by using existing [docker image](https://hub.docker.com/r/docker/welcome-to-docker)




### Installing [Docker Desktop](https://docs.docker.com/desktop/install/windows-install/)

![Installed docker](run-existing-docker-image/0-installed-docker.png)

Installed Docker Desktop should look like this.



### Running Docker image

![Docker Command](run-existing-docker-image/1-syntax-running-docker-image.png)

-p 8080:80 means host is at port 8080, local is at host 80
--name shows what is the name of the container



### Checking localhost on the specified port

![Docker web localhost](run-existing-docker-image/2-running-localhost.png)



### Showing the logs of the current running container

![Running container log](run-existing-docker-image/3-running-container-logs.png)



### Executing command cat to show the files inside

![Cat command](run-existing-docker-image/4-executing-command-in-container.png)



### Stopping server using the stop button on the top right corner

![Stop button](run-existing-docker-image/5-stopping-server.png)



### Logs during shutting down

![Shutting down log](run-existing-docker-image/6-log-after-shutting-down.png)



## Building Docker file from GO file




### Building dockerfile

![Dockerfile](run-existing-docker-image/7-building-docker-file.png)



### Naming the container and changing the host port

![Name and port](run-existing-docker-image/8-running-new-container.png)



### Log when running the dockerfile

![Log](run-existing-docker-image/9-running-container-log-dockerimg.png)



### Running the localhost port, writing dependant on the string written in main.go file

![Locahost](run-existing-docker-image/10-running-localhost-web.png)



### When the web is opened, it will show another log

![Log when web is opened](run-existing-docker-image/11-running-web-log.png)
