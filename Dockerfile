FROM golang:latest
RUN mkdir -p /Joeys
WORKDIR /
COPY go.mod .
COPY go.sum .
COPY Joeys/* /Joeys/
RUN go mod download
RUN go build -o app Joeys/main.go
CMD ["./app"]