FROM golang:alpine
COPY main.go .
RUN go build -o demo main.go
CMD ["./demo"]