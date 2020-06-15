# Choose lightweight Go image
FROM golang:1.13.5 as build

RUN mkdir /app
ADD . /app
WORKDIR /app

# Build bot binary for linux
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./...

# Run our bot within lightweight alpine linux
FROM alpine:latest AS production

COPY --from=build /app .
CMD ["./main"]