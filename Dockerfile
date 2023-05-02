FROM golang:alpine as build-stage
WORKDIR /build
COPY . .  
RUN go mod download
WORKDIR /build/cmd/icd/
RUN go build


 
##
## Deploy the application binary into a lean image
##
FROM debian:buster-slim AS runner
WORKDIR /app 
# lets copy the database, config file, and icd binary
RUN mkdir uploads
COPY --from=build-stage /build/cmd/icd/icd_codes_db.sqlite  /build/cmd/icd/.example.env /build/cmd/icd/icd ./
EXPOSE 3000 
CMD ["./icd", "-configFile", "./.example.env"]
 