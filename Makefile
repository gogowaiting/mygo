GO ?= go
HUGO ?= hugo
SITE_DIR := ./site

.PHONY: fmt vet test build ci
.PHONY: learn-list learn-tree learn-topic-advanced interview-list interview-random
.PHONY: site server site-build site-clean

# ===== Go =====

fmt:
	$(GO) fmt ./...

vet:
	$(GO) vet ./...

test:
	$(GO) test ./...

build:
	$(GO) build ./...

ci: fmt vet test build

# ===== Learn CLI =====

learn-list:
	$(GO) run ./cmd/learn list

learn-tree:
	$(GO) run ./cmd/learn tree

learn-topic-advanced:
	$(GO) run ./cmd/learn topic advanced

interview-list:
	$(GO) run ./cmd/interview list

interview-random:
	$(GO) run ./cmd/interview random

# ===== Hugo Site =====

server:
	cd $(SITE_DIR) && $(HUGO) server --bind 0.0.0.0 --port 1313 --disableFastRender

site-build:
	cd $(SITE_DIR) && $(HUGO) --gc --minify

site-clean:
	rm -rf $(SITE_DIR)/public

site: site-build
