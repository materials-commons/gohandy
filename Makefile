GODIRS = ezhttp handyfile desktop
P = github.com/materials-commons/gohandy

all: fmt test

test:
	rm -rf test_data/t
	-go test -v $(P)/ezhttp $(P)/handyfile

fmt:
	-for d in $(GODIRS); do (cd $$d; go fmt); done
