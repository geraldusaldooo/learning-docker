### Volume and Container

## In this case we are running postgres database through 2 different container using the same volume and see the difference



# Docker run with declaration of postgres username, password, volume and port.
![Docker run](https://github.com/geraldusaldooo/learning-docker/blob/main/my-container-volume/12-running-pg-command.png)



# Checking if the server, named my-postgres-geraldusaldooo, is running
![Server running](https://github.com/geraldusaldooo/learning-docker/blob/main/my-container-volume/13-server-running.png)



# Checking if the new volume has been created, under the name of my-pg-volume-geraldusaldooo
![Volume created](https://github.com/geraldusaldooo/learning-docker/blob/main/my-container-volume/14-new-volume-created.png)



# Creating new server in pgadmin and connecing it wih the Hostname, Port, Username and Password stated in the first command
![New server](https://github.com/geraldusaldooo/learning-docker/blob/main/my-container-volume/15-creating-new-server.png)



# Creating a table in pgadmin to check for changes later when we change container
![New table](https://github.com/geraldusaldooo/learning-docker/blob/main/my-container-volume/16-creating-table-in-pgadmin.png)



# Stopping and deleting the old container
![Old container](https://github.com/geraldusaldooo/learning-docker/blob/main/my-container-volume/17-stop-delete-old-container.png)



# Using the cmd prompt to create the 'newer' version of the container with v2 label
![New container](https://github.com/geraldusaldooo/learning-docker/blob/main/my-container-volume/19-new-container-running.png)



# Check if the new container is running
![Running v2 container](https://github.com/geraldusaldooo/learning-docker/blob/main/my-container-volume/19-new-container-running.png)



# The result is that the v2 has the table created in v1
![Table](https://github.com/geraldusaldooo/learning-docker/blob/main/my-container-volume/20-new-table-same-content.png)



## So even if the container is changed, if the container is still using the same volume, the data stored in the volume will carry on.