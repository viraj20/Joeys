FROM golang:latest
RUN mkdir -p /Joeys
COPY go.mod /Joeys/
COPY go.sum /Joeys/
WORKDIR /Joeys
RUN go mod download
COPY main.go /Joeys/
COPY pkg /Joeys/pkg/
RUN go build -o app main.go
CMD ["/Joeys/app"]