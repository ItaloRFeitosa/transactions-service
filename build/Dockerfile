FROM golang:1.21-alpine3.18 AS build

RUN apk add --no-cache git make

WORKDIR /app

COPY go.mod go.sum  ./
RUN go mod download

COPY . .

RUN make build

FROM alpine:3.18

COPY --from=build /app/out/api /app/api
RUN apk add ca-certificates

EXPOSE 8080

ENTRYPOINT ["/app/api"]