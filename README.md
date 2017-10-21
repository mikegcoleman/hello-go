# hello-go

A simple Docker image that creates a web server to display the container name.

To build this image use Golang-builder as follows:

1. Clone the repo to your local machine

2. If you forked the repo, update the canonical import path in `hello-go.go` to match your local github repo

  `package main // import "github.com/mikegcoleman/hello-go"`
  
3. Build the image with [Golang builder](https://github.com/CenturyLinkLabs/golang-builder)

  ```
  docker run --rm \
      -v $(pwd):/src \
      -v /var/run/docker.sock:/var/run/docker.sock \
      centurylink/golang-builder
  ```
    
