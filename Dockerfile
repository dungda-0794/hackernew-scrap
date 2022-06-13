# syntax=docker/dockerfile:1

FROM golang:1.17 as base

# for develop
FROM base as dev

# install the air binary so we get live code-reloading when we save files
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# run air command in the directory where our code will live
WORKDIR /hackernew-scrap

RUN go mod tidy

CMD ["air"]
