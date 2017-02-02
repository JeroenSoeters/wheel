#! /usr/bin/make

NAME=wheel

DEPEND=golang.org/x/tools/cmd/cover github.com/onsi/ginkgo/ginkgo \
       github.com/onsi/gomega github.com/rlmcpherson/s3gof3r/gof3r \
       github.com/golang/lint/golint

BRANCH?=dev
DATE=$(shell date '+%F %T')
COMMIT?=$(shell git symbolic-ref HEAD | cut -d"/" -f 3)

VERSION=$(NAME) $(BRANCH) - $(DATE) - $(COMMIT)
VFLAG=-X 'main.VERSION=$(VERSION)'

.PHONY: depend clean default

default: $(NAME)
$(NAME): $(shell find . -name \*.go)
	go build -ldflags "$(VFLAG)" -o $(NAME) .

build: $(NAME) build/$(NAME)-linux-amd64.tgz build/$(NAME)-darwin-amd64.tgz
# build/$(NAME)-linux-arm.tgz build/$(NAME)-windows-amd64.zip

build/$(NAME)-%.tgz: *.go
	rm -rf build/$(NAME)
	mkdir -p build/$(NAME)
	tgt=$*; GOOS=$${tgt%-*} GOARCH=$${tgt#*-} go build -ldflags "$(VFLAG)" -o build/$(NAME)/$(NAME) .
	chmod +x build/$(NAME)/$(NAME)
	cp README.md build/$(NAME)/
	tar -zcf $@ -C build ./$(NAME)
	rm -r build/$(NAME)

build/$(NAME)-%.zip: *.go
	touch $@


# Installing build dependencies. You will need to run this once manually when you clone the repo
depend:
	go get -v $(DEPEND)

clean:
	rm -rf build .vendor/pkg

lint:
	@if gofmt -l . | egrep -v ^vendor/ | grep .go; then \
	  echo "^- Repo contains improperly formatted go files; run gofmt -w *.go" && exit 1; \
	  else echo "All .go files formatted correctly"; fi
	#go tool vet -v -composites=false *.go
	#go tool vet -v -composites=false **/*.go
	for pkg in $$(go list ./... |grep -v /vendor/); do golint $$pkg; done

ci-test: cover

test: lint
	ginkgo -r -skipPackage vendor --randomizeAllSpecs --randomizeSuites --failOnPending

race: lint
	ginkgo -r -skipPackage vendor --randomizeAllSpecs --randomizeSuites --failOnPending --race


cover: lint
	ginkgo -r -skipPackage vendor --randomizeAllSpecs --randomizeSuites --failOnPending -cover
	@for d in `echo */*suite_test.go`; do \
	  dir=`dirname $$d`; \
	  (cd $$dir; go test -ginkgo.randomizeAllSpecs -ginkgo.failOnPending -cover -coverprofile $$dir_x.coverprofile -coverpkg $$(go list ./...|egrep -v vendor)); \
	done
	@rm -f _total
	@for f in `find . -name \*.coverprofile`; do tail -n +2 $$f >>_total; done
	@echo 'mode: atomic' >total.coverprofile
	@awk -f merge-profiles.awk <_total >>total.coverprofile
	COVERAGE=$$(go tool cover -func=total.coverprofile | grep "^total:" | grep -o "[0-9\.]*");\
	  echo "*** Code Coverage is $$COVERAGE% ***"