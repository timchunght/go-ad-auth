build_cross:
	mkdir release
	GOOS=linux   GOARCH=amd64 CGO_ENABLED=0 go build -o release/ad-auth-Linux-x86_64 .
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o release/ad-auth-Windows-x86_64.exe .
	GOOS=darwin  GOARCH=amd64 CGO_ENABLED=0 go build -o release/ad-auth-Darwin-x86_64 .