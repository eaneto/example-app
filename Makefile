build: clean
	mkdir -p bin
	go build -o bin/app

clean:
	rm -rf bin
