## GO源码编译

mac下编译为linux
> CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

mac下编译为windows
> CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build

linux下编译为mac
> CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build

linux下编译为windows
> CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build

windows下编译为mac
> CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go

windows下编译为linux
> CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
