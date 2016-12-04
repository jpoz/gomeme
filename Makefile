BINARY=gomeme

VERSION=1.1.0

REPO=github.com/jpoz/gomeme

SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '**/*.go')

.DEFAULT_GOAL: $(BINARY)

$(BINARY): $(SOURCES)
	go build -o ${BINARY} cmd/gomeme/gomeme.go

.PHONY: install
install:
	go install ${LDFLAGS} ./...

.PHONY: clean
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: assets
assets:
	go-bindata -pkg gomeme inpact.ttf
