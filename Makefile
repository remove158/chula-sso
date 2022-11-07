run :
	go run ./cmd/main.go
dev :
	nodemon --exec go run ./cmd/main.go --signal SIGTERM
gen :
	go generate ./...
test:
	go test -v ./...