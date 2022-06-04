PREFIX=/usr/local
BINARY_NAME=ranscanary

all: build

build:
	mkdir -p bin
	go build -o bin/$(BINARY_NAME) .

.PHONY: clean
clean:
	rm -rf ./bin

.PHONY: install
install: build
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp bin/$(BINARY_NAME) $(DESTDIR)$(PREFIX)/bin/$(BINARY_NAME)
	mkdir -p $(DESTDIR)$(PREFIX)/etc/$(BINARY_NAME)
	@echo "Installed to $(DESTDIR)$(PREFIX)/bin/$(BINARY_NAME)"

.PHONY: uninstall
uninstall:
	rm -f $(DESTDIR)$(PREFIX)/bin/$(BINARY_NAME)
	rm -rf $(DESTDIR)$(PREFIX)/etc/$(BINARY_NAME)
	@echo "Uninstalled from $(DESTDIR)$(PREFIX)"
