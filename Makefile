.PHONY: bin
bin:
	go build tweetbase.go

.PHONY: clean
clean:
	rm tweetbase

.PHONY: install
install: 
	go install tweetbase.go
