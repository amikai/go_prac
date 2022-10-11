.PHONY: test fuzz
test:
	go test -v .

fuzz:
	go test -pattern=Fuzz -fuzztime 10s
