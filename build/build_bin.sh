go_src=$(dirname "$PWD")
echo $go_src
docker run --rm -v $go_src:/build -w /build -e GOPROXY=https://mirrors.aliyun.com/goproxy,direct golang:1.14.2-buster go build -v