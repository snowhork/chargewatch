FROM golang:1.15 AS builder

WORKDIR /app

ENV GO111MODULE=on
ENV GOAVERSION=v3.2.6

RUN go get -u goa.design/goa/v3/...@${GOAVERSION}

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . /app
RUN design/goagen.sh

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o main ./cmd/chargewatch

#####
#####

FROM alpine:latest
COPY --from=builder /app/main /bin/main
RUN apk --no-cache add ca-certificates
EXPOSE 8088
CMD ["/bin/main"]
