FROM golang:alpine as build-stage 
WORKDIR /build
COPY . .
RUN go mod download 

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 build  -o main ./cmd/icd


EXPOSE 3000  
CMD [ "cmd/icd/main" ]


##
## Run the tests in the container
##
WORKDIR /
FROM build-stage AS run-test-stage
COPY -from=build-stage 
RUN go test -v build/cmd/icd/*.go





bin
.dockerignore
Dockerfile
docker-compose.yml
dev.Dockerfile
dev.docker-compose.yml
.env