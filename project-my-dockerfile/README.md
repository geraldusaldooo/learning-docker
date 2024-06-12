# Project 1: Build Dockerfile to Docker Image V2

This project aims to build dockerfile to docker image from golang.


### Creating folder containing AUTHORS.md, LINKS.md, README.md, a Dockerfile and golang files.
![Folder](https://github.com/geraldusaldooo/learning-docker/blob/main/project-my-dockerfile/Image/01-creating-folder.png)


### Creating golang project

Creating a golang project that serves http containing:
- Path '/' that prints string of html of written text, in this case I use "Connection succesful"
- The same path prints "Ouch!" in console log
- Stars on port 77

![Golang project](https://github.com/geraldusaldooo/learning-docker/blob/main/project-my-dockerfile/Image/02-golang-maingo.png)


### Re-writing AUTHORS.md file and LINKS.md file
![Author](https://github.com/geraldusaldooo/learning-docker/blob/main/project-my-dockerfile/Image/03-name-author.png)
![Link](https://github.com/geraldusaldooo/learning-docker/blob/main/project-my-dockerfile/Image/04-link-github.png)


### Running go mod init and go mod tidy
![Go mod](https://github.com/geraldusaldooo/learning-docker/blob/main/project-my-dockerfile/Image/07-go-mod.png)


### Creating a docker file

Creating a docker file that contains:
- Base image of golang:1.21
- Setting WORKDIR to /myapp
- Run command "go version" after WORKDIR
- Copy "AUTHORS.md" to image before building
- Copy "LINKS.md" to image before building
- Make golang build output with name "my-go-app"

![Dockerfile](https://github.com/geraldusaldooo/learning-docker/blob/main/project-my-dockerfile/Image/06-dockerfile.png)


### Changing go.mod file to required base image of golang
![Go mod edit](https://github.com/geraldusaldooo/learning-docker/blob/main/project-my-dockerfile/Image/05-go-mod-edit.png)


### Building image
![Build image](https://github.com/geraldusaldooo/learning-docker/blob/main/project-my-dockerfile/Image/08-build-image.png)


### Image has been created in Docker Desktop
![Image](https://github.com/geraldusaldooo/learning-docker/blob/main/project-my-dockerfile/Image/09-image-created.png)


### Creating container

Container needs to be created under these conditions:
- Name "go-app"
- Exposed to host port 5555

![Container creation](https://github.com/geraldusaldooo/learning-docker/blob/main/project-my-dockerfile/Image/10-creating-container.png)


### Log of the container during startup
![Log startup](https://github.com/geraldusaldooo/learning-docker/blob/main/project-my-dockerfile/Image/11-log-container.png)


### Checking connection to localhost

It prints the text I have written in the Golang file.

![Localhost response](https://github.com/geraldusaldooo/learning-docker/blob/main/project-my-dockerfile/Image/12-connection-local-host.png)


### Log during request

Prints the string "Ouch!" in console log

![Log during request](https://github.com/geraldusaldooo/learning-docker/blob/main/project-my-dockerfile/Image/13-connection-log.png)