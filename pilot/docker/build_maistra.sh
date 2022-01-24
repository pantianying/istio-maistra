CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o pilot-discovery ../cmd/pilot-discovery/*.go

docker buildx build -t quay.xiaodiankeji.net/yangchun/pilot-rhel8:latest -f Dockerfile.pilot.maistra . --platform=linux/amd64 --push