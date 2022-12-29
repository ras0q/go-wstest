APP_BIN   = ./bin/app
WSCAT_BIN = ./bin/wscat

.PHONY: build
build: $(APP_BIN) $(WSCAT_BIN)

$(APP_BIN): *.go
	go build -o $(APP_BIN) .

$(WSCAT_BIN): wscat/*.go
	go build -o $(WSCAT_BIN) ./wscat

.PHONY: clean
clean:
	rm -rf ./bin
