# ForProxy-A proxy server

## Install

```shell
git clone https://github.com/gralliry/ForProxy.git
cd ForProxy
```

## Build

### Linux
```shell
env GOOS=linux GOARCH=amd64 GIN_MODE=debug go build -o authserver
```

### windows
```shell
set GOOS=windows
set GOARCH=amd64
set GIN_MODE=debug
go build -o authserver.exe
```

### darwin(MacOS)
```shell
env GOOS=darwin GOARCH=amd64 GIN_MODE=debug go build -o authserver
```

### openwrt
```shell
env GOOS=linux GOARCH=mipsle GOMIPS=softfloat go build -ldflags='-s -w' -o /usr/bin/proxyserver
```

## Start

```shell
nohup /usr/bin/proxyserver -h 0.0.0.0 -p 8080 -t http://example.com > /tmp/proxyserver.log 2>&1 &
```

or

```shell
nohup go run main.go -h=0.0.0.0 -p=8080 -t=http://example.com > /tmp/proxyserver.log 2>&1 &
```