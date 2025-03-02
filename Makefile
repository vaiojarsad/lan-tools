ifeq ($(OS), Windows_NT)
    BIN := lan-tools.exe
else
    BIN := lan-tools
endif

# Build
build:
	go build -o $(BIN) .

# Clean
clean:
	rm -f lan-tools lan-tools.exe
