binary:
	GOOS=linux go build -o charlie.bin ./charlie.go

container:
	docker build -f Dockerfile.charlie -t charlie:latest .

all:
	make binary
	make container

run:
	docker run charlie:latest

clean:
	rm -f charlie.bin