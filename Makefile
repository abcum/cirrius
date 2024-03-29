# Copyright © 2016 Abcum Ltd
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

GO ?= CGO_ENABLED=1 go
LDF := -s -w

.PHONY: default
default:
	@echo "Choose a Makefile target:"
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print "  - " $$1}}' | sort

.PHONY: clean
clean:
	$(GO) clean -i -n

.PHONY: tests
tests:
	$(GO) test ./...

.PHONY: build
build: LDF += $(shell GOPATH=${GOPATH} build/flags.sh)
build:
	$(GO) build -v -tags 'pdflib9' -ldflags '$(LDF)'

.PHONY: install
install: LDF += $(shell GOPATH=${GOPATH} build/flags.sh)
install:
	$(GO) install -v -tags 'pdflib9' -ldflags '$(LDF)'

.PHONY: compile
compile: LDF += $(shell GOPATH=${GOPATH} build/flags.sh)
compile:
	$(GO) install -v -tags 'pdflib9' -ldflags '$(LDF)'
	$(GO) build -o main-osx -v -tags 'pdflib9' -ldflags '$(LDF)'
	aws s3 cp --cache-control "no-store" main-osx s3://pkg.cirrius.io/osx
	rm -rf main-osx
	#
	go get github.com/karalabe/xgo
	xgo -x -v -out main-aws -tags 'aws pdflib8' -ldflags '$(LDF)' --targets=linux/amd64 github.com/abcum/cirrius
	aws s3 cp --cache-control "no-store" main-aws-linux-amd64 s3://pkg.cirrius.io/aws
	rm -rf main-aws-linux-amd64
