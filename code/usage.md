## Compile docker

CGO_ENABLED=0 GOOS=linux go build -o agent  -a -ldflags '-extldflags "-static"'  -o agent ./cmd/agent/main.go

docker build -f ../docker/agent/Dockerfile -t registry.sensetime.com/diamond/zhangxiaohu/bootcamp/agent:v0.1.0

docker push registry.sensetime.com/diamond/zhangxiaohu/bootcamp/agent:v0.1.0
