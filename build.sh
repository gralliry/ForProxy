#!/bash/bin

env GOOS=linux GOARCH=mipsle GOMIPS=softfloat go build -ldflags='-s -w' -o /usr/bin/proxyserver
