# http-server
A simple app to create, delete and list users using Mux router.

Tech stack :
- App has been written in Golang
- Used Docker to containerize the app
- Used Docker Compose to demonstrate multi container deployments
- Written K8s components to orchestrate the containers
- Added helm charts for easy deployments of K8 clusters

## Installing and running the app locally

```
  git clone https://github.com/niranjan-n1/http-server.git
  
  cd http-server
  
  go run main.go
```

# Deploying the app as a Docker container

### Building the Docker Image locally

```
  docker build . --tag <imagename>:<tag>  
```

### Pulling the image from docker.io registry

```
  docker pull niranjann1/httpserver:1.1 
```

### Verify the image 

```
  docker images
```
### Running the image as container

```
docker run -d -p 8081:8081 --name <container-name> httpserver:1.1
```

### Verify that the container is running 

```
  docker ps
```
### Running mutiple instances of the application using docker-compose

```
  docker-compose up -d 
```

## Deployment using Helm

### Install Helm CLI

```
  curl -LO https://get.helm.sh/helm-v3.9.0-linux-amd64.tar.gz
  tar -C /tmp/ -zxvf helm-v3.9.0-linux-amd64.tar.gz
  rm helm-v3.9.0-linux-amd64.tar.gz
  mv /tmp/linux-amd64/helm /usr/local/bin/helm
  chmod +x /usr/local/bin/helm
```

### Installing the chart

```
  helm install <chart-name> <chart>
```
Eg :

```
  helm install my-server helm 
```

  
