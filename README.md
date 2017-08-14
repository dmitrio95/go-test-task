# go-test-task
This repository is intended to contain a code for a test task. This is a simple implementation of a web service trying to follow REST principles that offers an information on the server's network interfaces.

# Build instructions
## Build a Go app
The following commands will download the project to your `$GOPATH`, build and install it to your Go binaries directory.
```
go get github.com/dmitrio95/go-test-task
cd "$GOPATH"/src/github.com/dmitrio95/go-test-task
./build.sh
```
After launching it will listen on 8080 port by default, but you can override this by `--address` command line option:
```
./go-test-task --address :12345
```

## Build a Docker image
You can also build a Docker image for this service using a `build_docker.sh` script:
```
go get github.com/dmitrio95/go-test-task
cd "$GOPATH"/src/github.com/dmitrio95/go-test-task
./build_docker.sh
```
Depending on your system's settings `build_docker.sh` may require root privileges. In that case make sure you pass a proper `GOPATH` environment variable to that script:
```
sudo GOPATH=/path/to/go/workspace/ ./build_docker.sh
```
You can then launch that image:
```
docker run -p 12345:8080 go-test-task-image
```
where you can replace `12345` with any desired port number. Then you can connect to this service, for example, using a web browser:
```
firefox example.com:12345
```
