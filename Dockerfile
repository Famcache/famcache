FROM golang:1.21.6 AS builder

WORKDIR /famcache

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux make build

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /famcache

COPY --from=builder /famcache/bin/server .

CMD [ "./server" ]
