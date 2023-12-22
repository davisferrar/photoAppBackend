.ONESHELL:

all: clean build run

clean:
	rm -f -r build

build: clean
	mkdir build
	go build -o ./build

run:
	./build/distributed-system-mission1