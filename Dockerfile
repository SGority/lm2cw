FROM golang:1.13.7-buster
WORKDIR /app

COPY go.mod .
RUN go mod download
RUN go get github.com/GeertJohan/go.rice/rice
COPY . .

RUN go build -v -o /bin/api /app/cmd/api

FROM alpine
COPY --from=0 /bin/api /bin/api
RUN mkdir templates
COPY templates/* templates/
ENTRYPOINT ["api"]
