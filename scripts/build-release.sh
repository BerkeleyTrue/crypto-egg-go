IPATH="github.com/berkeleytrue/crypto-egg-go/config"
HASH=$(git rev-parse --short HEAD)
USER=$(id -u -n)

mkdir -pv build

go build -ldflags="-X $IPATH.Port=10000 \
  -X $IPATH.GinReleaseMode=true \
  -X '$IPATH.Hash=$HASH' \
  -X '$IPATH.Time=$(date)' \
  -X '$IPATH.User=$USER'" -o ./build ./cmd/app ./cmd/cli

echo "Build Complete: Hash: $HASH, User: $USER"
