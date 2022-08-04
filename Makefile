all: build

tidy:
	GOPROXY=https://mirrors.aliyun.com/goproxy,direct GO111MODULE=on go mod tidy

build:
	GOPROXY=https://mirrors.aliyun.com/goproxy,direct GO111MODULE=on go build -v

.PHONY: clean
clean:
	rm -f weixin