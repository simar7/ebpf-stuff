binary:
	GOOS=linux go build -o snoopy.bin ./execsnoop.go ./output.go

container:
	docker build -f Dockerfile.snoopy -t snoopy:latest .

all:
	make binary
	make container

run:
	docker run --priviledged charlie:latest

clean:
	rm -f snoopy.bin