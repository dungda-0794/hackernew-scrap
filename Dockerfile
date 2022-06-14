FROM golang:1.18.3-alpine3.15 as base

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash openssh

# for develop
FROM base as dev

# run air command in the directory where our code will live
WORKDIR /hackernew-scrap
RUN chmod a+x /hackernew-scrap

COPY go.mod go.sum ./
RUN go mod download

# install the air binary so we get live code-reloading when we save files
RUN go install github.com/cosmtrek/air@latest

