run:
	go run ./src/app/main.go

test:
	echo "Using verbose style"
	go test ./... -v -count=1