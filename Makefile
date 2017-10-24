
OUT=./build

.PHONY: build
build:
	echo "Build go"
	mkdir -p $(OUT)
	go build -v -o $(OUT)/jsonmarkdown main.go

.PHONY: clean
clean:
	rm -rf $(OUT)/*
