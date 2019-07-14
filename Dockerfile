FROM golang as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/doggos

# ------------------------------------------------------------------------------

FROM scratch
COPY --from=builder /app/doggos /app/
EXPOSE 7000
ENTRYPOINT ["/app/doggos"]