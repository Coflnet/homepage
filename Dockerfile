# build stage
FROM golang:1.21-bookworm as builder

WORKDIR /app
# ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .

RUN go mod download

# FROM build_base AS server_builder
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build


# final stage
# add scratch when dev is done
FROM scratch
COPY --from=builder /app/homepage /app/

EXPOSE 9658
ENTRYPOINT ["/app/homepage"]
