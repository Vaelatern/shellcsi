.PHONY: build

build: cmd/shellcsi
	go build ./cmd/shellcsi

clean:
	rm ./shellcsi
