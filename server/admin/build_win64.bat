go-bindata -pkg conf -o app/bindata/conf/conf_data.go conf/...
go-bindata -pkg static -o app/bindata/static/static_data.go resources/...

::go build -ldflags "-w -s" -o admin.exe

