
bin/customcamera: *.go cmd/module/*.go
	go build -o bin/customcamera cmd/module/cmd.go

lint:
	gofmt -w -s .

updaterdk:
	go get go.viam.com/rdk@latest
	go mod tidy

module: bin/customcamera
	tar czf module.tar.gz bin/customcamera