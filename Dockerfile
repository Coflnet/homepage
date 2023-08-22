# build stage
FROM golang:1.21-bookworm as builder

WORKDIR /app
# ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .

RUN go mod download

# FROM build_base AS server_builder
COPY . .
RUN CGO_ENABLED=0 go build ./cmd/homepage/main.go


# final stage
# add scratch when dev is done
FROM scratch
COPY --from=builder /app/main /app/

EXPOSE 9658
ENTRYPOINT ["/app/main"]
