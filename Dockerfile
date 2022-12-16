# BUILD stage (using go image to build)
FROM golang:1.18 AS build-base

RUN apt-get update && \
    apt-get --allow-unauthenticated install -y  git tree && \
    go env

WORKDIR /padel-field
# copying the entire repo inside /padel-field/
ADD . /padel-field/

ENV GOPATH /go
# compiling
RUN GOOS=linux GOARCH=amd64 go build -o ./bin/padelField

# FINAL stage
FROM alpine:3.16.2 as run-padelfield

WORKDIR /app
RUN apk --no-cache add libc6-compat tzdata && \
    apk add bash && apk update


# binary
COPY --from=build-base /padel-field/bin/padelField /app/bin/run

# other resources
COPY --from=build-base /padel-field/api/templates/* /app/api/templates/
COPY --from=build-base /padel-field/resources/* /app/resources/

# db
ENV ADDR=172.20.0.5:5432
ENV DB_USER=root
ENV PASSWORD=root
ENV DATABASE=padelfield

ENV GIN_MODE=release

EXPOSE 8080

ENTRYPOINT ["/app/bin/run"]
# -r: restart the entire DB
# -s: start gin gonic sv
CMD ["-r", "-s"]

