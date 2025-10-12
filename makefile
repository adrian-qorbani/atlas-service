# debug
run:
	go run apis/services/sales/main.go | go run apis/tooling/logfmt/main.go

# module support
tidy:
	go mod tidy
	go mod vendor