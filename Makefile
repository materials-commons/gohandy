DIRS = .

all: fmt test

test:
	-go test -v
	-for d in $(DIRS); do (cd $$d; go test -v); done

fmt:
	go fmt
	-for d in $(DIRS); do (cd $$d; go fmt); done
