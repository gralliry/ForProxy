set GOOS=linux
set GOARCH=mipsle
set GOMIPS=softfloat

go build -ldflags="-s -w" -o proxyserver