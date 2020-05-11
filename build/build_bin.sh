projectDir=$(dirname "$PWD")
echo $projectDir
docker run --rm -v $projectDir:/projectDir -w /projectDir -e GOPROXY=https://mirrors.aliyun.com/goproxy,direct golang:1.14.2-buster go build -v -o /projectDir/build