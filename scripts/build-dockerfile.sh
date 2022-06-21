HASH=$(git rev-parse --short HEAD)
docker build \
  -f resources/Dockerfile \
  --build-arg HASH=$(git rev-parse --short HEAD) \
  --build-arg USER=$(id -u -n) \
  -t crypto-egg-go/server:$HASH .

docker tag \
  crypto-egg-go/server:$HASH \
  crypto-egg-go/server:latest
