# GoTube

Example project from ["Bootstrapping Microservices with Docker, Kubernetes, and Terraform"](https://www.manning.com/books/bootstrapping-microservices-with-docker-kubernetes-and-terraform) book.

Based on Go + Fiber.

## Microservices

At this moment (incomplete) project consists of 5 microservices, database and message broker:

1. Video streaming service. Gets video id from query param, search for video path in *database*, get video stream from *video storage* service and returns to client.
2. Video storage service. Uses Azure Blob storage for videos storage and returns stream by path for *video streaming* service.
3. History service. Receives "viewed" events from *video streaming* service by post request and puts it into *database*.
4. MongoDB for storing data.
5. RabbitMQ for indirect communication.
6. Recommendation service. Receives "viewed" events from *video streaming* by RabbitMQ and simply logs them.
7. Metadata service. Provides video metadata info.

## Requirements

* docker
* docker-compose

### Optional

* make

## How to start?

Use `make up` and `make down` to start and stop containers.

Use `make dev` for development with support of live reload by `air`.

## See also

* NodeJS + Express version: https://github.com/capcom6/node-tube
* Python + Flask version: https://github.com/capcom6/py-tube
