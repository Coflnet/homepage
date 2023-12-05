# build stage
FROM registry.suse.com/bci/golang:1.21 as builder

WORKDIR /app
# ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .

RUN go mod download

# FROM build_base AS server_builder
COPY . .
RUN CGO_ENABLED=0 go build ./cmd/homepage/main.go


FROM registry.suse.com/bci/bci-micro:15.5

COPY --from=builder /app/main /app/
COPY --from=builder /app/internal/views /app/internal/views
COPY --from=builder /app/static /app/static
COPY --from=builder /app/config.yaml /app/config.yaml
COPY --from=builder /app/active.en.toml /app/active.en.toml
COPY --from=builder /app/active.de.toml /app/active.de.toml
WORKDIR /app

EXPOSE 9658
ENTRYPOINT ["/app/main"]
