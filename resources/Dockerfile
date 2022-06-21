FROM golang:1.18-alpine AS builder
ENV IPATH="github.com/berkeleytrue/crypto-egg-go/config"
ARG HASH
ENV HASH=$HASH
ARG USER
ENV USER=$USER
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -ldflags="\
  -X $IPATH.Port=10000 \
  -X $IPATH.GinReleaseMode=true \
  -X '$IPATH.Hash=$HASH' \
  -X '$IPATH.Time=$(date)'" \
  -o /usr/local/bin/server \
  ./cmd/app

FROM golang:1.18-alpine AS runner
WORKDIR /app
COPY config/config.yml ./config/config.yml
COPY --from=builder /usr/local/bin/server ./
EXPOSE 10000
CMD ["/app/server"]
