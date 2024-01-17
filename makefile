build:
	go build -o ./bin/shorty.so -buildmode=plugin .


vendor:
	go mod vendor