provision:
	go run main.go provision -s lambhack

update:
	go run main.go provision -s lambhack -c

test:
	@go test ./...
