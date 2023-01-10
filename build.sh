# Linux
GOOS=linux GOARCH=amd64 go build -o build/epb-linux-64 epb.go
GOOS=linux GOARCH=386 go build -o build/epb-linux-32 epb.go

# Windows
GOOS=windows GOARCH=amd64 go build -o build/epb-win-64.exe epb.go
GOOS=windows GOARCH=386 go build -o build/epb-win-32.exe epb.go

# MacOS
GOOS=darwin GOARCH=amd64 go build -o build/epb-darwin-64 epb.go