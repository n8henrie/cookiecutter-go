LDFLAGS := -ldflags "-X main.version=$$($(GIT) describe --tags)"

# Remove the `-` to fail build comamnds if $VAR is unset
.PHONY: env
env:
	-@test $(VAR) || { echo 'err: $$VAR is unset'; exit 1; }

.PHONY: help
help:
	@awk '/^[^ ]*:/ { gsub(":.*", ""); print }' Makefile

.PHONY: build
build: env
	go build $(LDFLAGS) -v .

.PHONY: test
test:
	go test -race -v .

.PHONY: bench
bench:
	go test -race -bench=. -benchmem

.PHONY: rpi
rpi: env
	GOOS=linux GOARCH=arm GOARM=7 go build $(LDFLAGS) -v .

.PHONY: oldrpi
oldrpi: env
	GOOS=linux GOARCH=arm GOARM=6 go build $(LDFLAGS) -v .

.PHONY: macos
macos: env
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -v .

.PHONY: linux
linux: env
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -v .

.PHONY: clean
clean:
	go clean -x
