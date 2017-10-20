.PHONY: bin
bin:
	go build tweetbase.go

.PHONY: clean
clean:
	rm tweetbase

.PHONY: install
install: 
	go install tweetbase.go

.PHONY: install-gotools
install-gotools:
	go get github.com/Masterminds/glide
	go get github.com/sgotti/glide-vc

.PHONY: update-vendor
update-vendor: install-gotools
	glide update --strip-vendor
	glide-vc --only-code --no-tests
