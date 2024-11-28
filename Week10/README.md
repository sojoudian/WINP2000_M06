
# Simple Web Server in Go with Docker Integration

---

## Purpose of the Project
The purpose of this project is to teach the following concepts:
1. Building a simple web server in Go.
2. Serving files over the web using a file server.
3. Creating a basic `index.html` for the web interface.
4. Dockerizing a simple Go application.
5. Creating a Docker image, running it locally, tagging the image, and pushing it to Docker Hub.
6. Using the Docker image anywhere by pulling it from Docker Hub.

---

## Project Structure
- **`main.go`**: Contains the Go code for the web server.
- **`index.html`**: A basic HTML webpage that will be served by the Go server.
- **`Dockerfile`**: Dockerfile for containerizing the Go application.
- **`go.mod`**: Dependency management file for Go modules.

---

## Step-by-Step Guide

### 1. Building the Web Server in Go
The web server is written in Go and serves files from a directory. The main components include:
- `http.FileServer`: Serves static files from a directory.
- `http.Handle`: Maps a path to the file server.
- `http.ListenAndServe`: Starts the server.

#### Example Code (main.go):
```go
package main

import (
    "net/http"
    "log"
)

func main() {
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/", fs)

    log.Println("Server is running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

---

### 2. Creating the `index.html` File
The `index.html` file serves as the primary web page. It includes:
- A responsive design for mobile and desktop viewing.
- Sections for "About Me," "Projects," and "Contact."

Place the `index.html` file in a directory named `static/` to be served by the Go server.

---

### 3. Running the Go Server Locally
1. Install Go on your system.
2. Place the `index.html` in a `static/` folder.
3. Run the following command in the terminal:
   ```bash
   go run main.go
   ```
4. Open `http://localhost:8080` in your browser to see the `index.html` file.

---

### 4. Dockerizing the Go Application
The Dockerfile contains the instructions to build and run the Go application in a container.

#### Dockerfile:
```dockerfile
# Use official Golang image for building the app
FROM golang:1.20 as builder
WORKDIR /app

# Copy Go modules and app code
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the Go binary
RUN go build -o server main.go

# Use a lightweight base image for the runtime
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/server .
COPY ./static ./static

# Expose port 8080 and run the app
EXPOSE 8080
CMD ["./server"]
```

#### Steps to Build and Run the Docker Image
1. Build the Docker image:
   ```bash
   docker build -t my-go-app .
   ```
2. Run the image locally:
   ```bash
   docker run -p 8080:8080 my-go-app
   ```

---

### 5. Pushing the Docker Image to Docker Hub
#### Prerequisites
- A Docker Hub account.
- Login to Docker Hub:
  ```bash
  docker login
  ```

#### Steps
1. Tag the Docker image:
   ```bash
   docker tag my-go-app <your-dockerhub-username>/my-go-app
   ```
2. Push the image to Docker Hub:
   ```bash
   docker push <your-dockerhub-username>/my-go-app
   ```

---

### 6. Using the Docker Image Anywhere
You can pull the Docker image on any system and run it without needing the source code.

#### Steps:
1. Pull the Docker image:
   ```bash
   docker pull <your-dockerhub-username>/my-go-app
   ```
2. Run the image:
   ```bash
   docker run -p 8080:8080 <your-dockerhub-username>/my-go-app
   ```

---

## Conclusion
This project covers the fundamentals of building a web server in Go, creating a Dockerized application, and deploying it using Docker Hub. The steps provide a practical understanding of containerization and web server deployment.


# Docker commands
 ```bash
 docker build -t dw10:0.1 .
 docker run dw10:0.1
 docker run -p 80:80 dw10:0.1
 docker run -p 8033:80 dw10:0.1
 docker login
 docker tag dw10:0.1 maziar/dw10:0.1
 docker push maziar/dw10:0.1
```
