# BUILD stage (uses native archetecture)
FROM golang:1.18 AS build-base

RUN apt-get update && \
    apt-get --allow-unauthenticated install -y  git tree && \
    go env

WORKDIR /base/padel-field
ADD . /base/padel-field/

ENV GOPATH /go
RUN make compile

# FINAL stage
FROM ubuntu/postgres as run-padelfield

RUN apt-get update
RUN apt-get install -y apt-transport-https  \
    ca-certificates curl gnupg lsb-release

WORKDIR /padel-field

COPY --from=build-base /base/padel-field/ /padel-field
COPY --from=build-base /base/padel-field/bin/padelField /padel-field/bin/padelField

ENV READY="http://0.0.0.0:8080/"
EXPOSE 8080

ENTRYPOINT ["/padel-field/bin/padelField", "-r", "-s"]
