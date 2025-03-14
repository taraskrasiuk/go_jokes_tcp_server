build:
	go build -o ./bin/jokes ./cmd
run: build
	./bin/jokes $(ARGS)
clear:
	rm -vrf ./bin
test:
	go test ./...
