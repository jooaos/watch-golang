# About

This project leverages **Go (Golang)**, **Air** and **Docker Compose** to enable hot reloading while maintaining the benefits of containerization with lightweight images.

# How it works?

In this setup, Docker Compose configures and runs your Go application within an isolated environment. Every time a code change is detected, Air automatically recompiles the binary and the Docker reloads the application with this binary, not requiring a manual container restart.

## Advantages

1. Fully configured development environment using Docker.
2. Hot reloading during development to streamline workflows.
3. Easy scalability: Add new services to your application with Docker Compose.

# Prerequisites

## Docker compose

Ensure that [Docker Compose](https://docs.docker.com/compose/compose-file/) is installed on your system with version **>= 2.22.0**. This is required to take advantage of the [watch mode](https://docs.docker.com/compose/how-tos/file-watch/)


## Air

[Air](https://github.com/air-verse/air) is a tool that handles automatic recompilation of Go binaries whenever code changes are detected.


## Make (no required)

Although not strictly required, we recommend installing **Make** to simplify project usage

> If your prefer not use it, you can manually copy and paste the commands from the Makefile into your terminal


# Step by step

1. Start the application with Air:
```
make air-up service=example
```
2. Start the containers:
```
make up
```
3. Make a code change to the ```example``` service and observe the logs to accert the automatic reload.
