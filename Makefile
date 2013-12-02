GODIRS = ezhttp handyfile

all: fmt test

test:
	rm -rf test_data/t
	-for d in $(GODIRS); do (cd $$d; go test -v); done

fmt:
	-for d in $(GODIRS); do (cd $$d; go fmt); done
