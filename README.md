# About

The project consist in use golang + air + docker compose to make hot reloads and keep using the container concept using light images.

# How it works?

Here, the docker compose configures and execute the your go application in an isolated environmente, and in each code change, the air detects changes and recompile the binary automatic, without the necessity of restart your container manually.

## Advantages

1. Dev environment configured and using Docker,
2. Hot reload to your code during the development,
3. Facility to scale up your application with new services using docker compose

# Pré-requisito

## Docker compose

It's necessary that [Docker Compose](https://docs.docker.com/compose/compose-file/) installed in yout host has a version >= 2.22.0, just in this way we can use the [watch mode](https://docs.docker.com/compose/how-tos/file-watch/)


## Air

The air is responsible for recompile the binaries that will be injected inside the containers, for the installation, check [here]()


## Make (no required)

To facility the usability of project, we recommend the make installation

> If you don't want use it, you can access Makefile, copy and paste the command in your terminal


# Step by step

1. 
