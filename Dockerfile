# Choose lightweight Go image
FROM golang:1.22.2 as build

RUN mkdir /app
ADD . /app
WORKDIR /app

# Build bot binary for linux
RUN CGO_ENABLED=0 GOOS=linux go build

# Run our bot within lightweight alpine linux
FROM alpine:latest AS production

COPY --from=build /app .
CMD ["./cacao"]