test:
	go test

build:
	env GOOS="linux" go build -o bin/uni main.go
	go build -o bin/uni main.go GOOS="windows"
	env GOOS="darwin" go build -o bin/uni main.go



install-windows:
	choco install nasm
	choco install mingw

install-linux:
	curl -O https://www.nasm.us/pub/nasm/releasebuilds/2.15.05/nasm-2.15.05.tar.gz
	tar -cvzf nasm-2.15.05.tar.gz nasm
