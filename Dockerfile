FROM golang:latest
RUN mkdir -p /Joeys
COPY go.mod /Joeys/
COPY go.sum /Joeys/
COPY main.go /Joeys/
COPY pkg /Joeys/pkg/
WORKDIR /Joeys
RUN go mod download
RUN go build -o app main.go
CMD ["/Joeys/app"]