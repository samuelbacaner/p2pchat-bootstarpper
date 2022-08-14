FROM golang:1.18-alpine as builder
RUN mkdir /build
ADD *.go go.mod go.sum /build/
WORKDIR /build
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -o p2pchat-bootstrapper .

FROM alpine:3.16.2
COPY --from=builder /build/p2pchat-bootstrapper .

# executable
ENTRYPOINT [ "./p2pchat-bootstrapper" ]
# arguments that can be overridden
CMD [ "3", "300" ]