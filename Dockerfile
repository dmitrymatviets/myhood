#### Setting base build image
FROM golang:1.14-alpine3.11 AS builder
# Variables
ENV APPDIR $GOPATH/src/myhood
ENV ARTIFACT /build/myhood
ENV UI /build/ui
# Installing deps
RUN apk --update --no-cache add git
# Making directory for building app from source
RUN mkdir -p ${APPDIR}
# Setting created directory as our working directory
WORKDIR ${APPDIR}
# Copying source code
COPY . .
COPY ./ui ${UI}
RUN ls -la
RUN go mod download
RUN go mod vendor
# Building application
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -ldflags "-s -w" -o ${ARTIFACT}
RUN ls -la


### Setting base image
FROM alpine:3.11

# Variables
ENV ARTIFACT /build
ENV BINDIR /usr/local/bin

RUN ulimit -l
# Setting up timezone
RUN apk --update --no-cache add tzdata git curl

RUN cp /usr/share/zoneinfo/Europe/Moscow /etc/localtime && \
    echo "Europe/Moscow" > /etc/timezone && \
    date
# Copying artifact from build image
COPY --from=builder ${ARTIFACT} ${BINDIR}

WORKDIR ${BINDIR}

RUN ls -la

# Setting up container entrypoint and command
ENTRYPOINT [ "myhood" ]



