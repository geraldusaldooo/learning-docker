# learning-docker


## Task is running existing docker image [welcome-to-docker](https://hub.docker.com/r/docker/welcome-to-docker) as a container

### First off is to install [Docker Desktop](https://docs.docker.com/desktop/install/windows-install/)

After installing Docker Desktop it should look like this.

![Installed docker](run-existing-docker-image/0-installed-docker.png)


### Running the existing container (welcome-to-docker) using command prompt or git bash.

![How to run the container](run-existing-docker-image/1-syntax-running-docker-image.png)


Then when opening localhost with the specified port (8080) it should look like the image below.

![Running in localhost](run-existing-docker-image/2-running-localhost.png)


### Looking at the logs of running server

![Logs of running server](run-existing-docker-image/3-running-container-logs.png)


### Now executing command inside the Docker container, for example a cat function

![Executing command in Docker exec](run-existing-docker-image/4-executing-command-in-container.png)


### To stop the container we can use the stop button located on the right corner of Docker Desktop

![Stopping server through Docker](run-existing-docker-image/5-stopping-server.png)


### What the logs look like during the shutting down process

![Log when shutting down](run-existing-docker-image/6-log-after-shutting-down.png)
