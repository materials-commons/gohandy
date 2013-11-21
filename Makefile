GODIRS = ezhttp .

all: fmt test

test:
	-for d in $(GODIRS); do (cd $$d; go test -v); done

fmt:
	-for d in $(GODIRS); do (cd $$d; go fmt); done
