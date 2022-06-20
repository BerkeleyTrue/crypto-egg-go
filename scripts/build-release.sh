IPATH="github.com/berkeleytrue/crypto-egg-go/config"

mkdir -pv build

go build -ldflags="-X $IPATH.Port=10000 \
  -X $IPATH.GinReleaseMode=true \
  -X '$IPATH.Hash=$(git rev-parse --short HEAD)'
  -X '$IPATH.Time=$(date)' \
  -X '$IPATH.User=$(id -u -n)'" -o ./build ./cmd/app ./cmd/cli
