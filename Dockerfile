FROM golang:alpine as build-stage
WORKDIR /build
COPY . .  
RUN go mod download 
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /build/cmd/icd/main.go

##
## Run the tests in the container
##
FROM build-stage AS run-test-stage
RUN go test -v ./...

##
## Deploy the application binary into a lean image
##
FROM debian:buster-slim
WORKDIR /app 
COPY --from=builder /build/cmd/icd/main .
COPY --from=builder /build/icd_db . 
EXPOSE 3000 
ENTRYPOINT ["./main"]