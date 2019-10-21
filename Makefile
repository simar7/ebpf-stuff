binary:
	GOOS=linux go build ./helloworld.go

container:
	docker build . -t helloworld:latest

all:
	make binary
	make container

run:
	docker run helloworld:latest

clean:
	rm -f helloworld