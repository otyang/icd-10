FROM golang:alpine as build-stage
WORKDIR /build
COPY . .  
RUN go mod download  
#RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 build  -o ma ./cmd/icd
WORKDIR /build/cmd/icd/
RUN go build && ls
EXPOSE 3000
CMD ["./icd"]