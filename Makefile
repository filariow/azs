.PHONY: build

GO      := go
GOFLAGS := -a -ldflags "-s -w" -trimpath
OUT     := bin/azs
SRC     := cmd/azs/azs.go

build:
	$(GO) build $(GOFLAGS) -o $(OUT) $(SRC)

