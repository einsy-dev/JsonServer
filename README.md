go generate ./cmd/

go build -o server.exe -ldflags="-H windowsgui" ./cmd/