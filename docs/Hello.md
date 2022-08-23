
# Hello Service

The `Hello` application is a simple web server that listen to hello requests and keep a count of each name "greeted".

## Build

### Requirements

- [Go (Golang) version 1.14+](https://golang.org/dl/)
- [Docker](https://docs.docker.com/get-docker/)

Optional:

- [JQ](https://stedolan.github.io/jq/download/)


### Building GO binary
```shell
cd hello
CGO_ENABLED=0 go build -v -ldflags="-X main.version=1.2.3" -o hello
```

### Building Docker Image

```shell
export VERSION="1.2.3"
docker build -t hello:v${VERSION} --build-arg VERSION=${VERSION} .
```

Test the image using:

```shell
docker run -ti --rm hello:v${VERSION} -v
```

### Public Image

We maintain a public image of the application on Github at [nytimes/hello:v1.0.0](https://hub.docker.com/r/nytimes/hello)

## Usage

```shell
./hello -h

Usage of ./hello:
  -port="8080": the port to listen on, default to 8080
  -v=false: Show version and quit
```

## Service API

1. `GET /hello/:name` which responds with `Hello, <name>!`.
2. `GET /healthz` which would serve as a health check endpoint, and responds with some system stats/metrics of your choice
3. `GET /counts` which returns a JSON response with the counts of how many times each name has been called, in the format below. No need for formatting or capitalizing the names, just return exactly what was posted.

    [
     {"name": "alice", "count": 2},
     {"name": "bob", "count": 1}
    ]

1. `DELETE /counts` Which resets the data so there are no counts (empty counts array).
2. Any undefined route requested should return an HTTP 404 not found status.

## Example

### Running the app in Docker

You can run the app using your local Docker image that you built before, sharing the port 8080:

```shell
docker run -ti --rm -p 8080:8080 hello:v${VERSION}
```

Or you can run the app using the public image. The image is tagged with version `v1.0.0` and hosted in DockerHub. Set your `VERSION` variable to use it:

```shell
export VERSION="1.0.0"
docker run -ti --rm -p 8080:8080 nytimes/hello:v${VERSION}
```

### Check Health

The `/healthz` endpoint will return a `200` code with on `OK` status and the uptime of the application.

```shell
curl -s http://localhost:8080/healthz

{"uptime":"2m31.629924755s","status":"OK"}
```

### Hello a person

You can greet `spongebob` and `patrick` by calling the `/hello` endpoint. You can greet the same person multiple times:

```shell
curl -s http://localhost:8080/hello/spongebob
Hello, spongebob!

curl -s http://localhost:8080/hello/patrick
Hello, patrick!

curl -s http://localhost:8080/hello/spongebob
Hello, spongebob!
```

### Count of hellos

You can get the count of past hellos by calling the `/counts` API endpoint:

```shell
 curl -s http://localhost:8080/counts|jq '.'
[
  {
    "name": "spongebob",
    "count": 2
  },
  {
    "name": "patrick",
    "count": 1
  }
]
```

### Reset the counts

You can reset the counts by calling the `/counts` endpoint with the `DELETE` action:

```shell
curl -X DELETE -s http://localhost:8080/counts

curl -s http://localhost:8080/counts
[]
```
