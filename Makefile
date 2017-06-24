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

GO ?= CGO_ENABLED=0 go
LDF :=

# The `make default` command cleans
# the go build and test files and
# then runs a build and install.

.PHONY: default
default:
	@echo "Choose a Makefile target:"
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print "  - " $$1}}' | sort

# The `make kill` command ensures that
# any hanging cirrius processes are force
# killed. Useful in development.

.PHONY: kill
kill:
	pkill -9 -f cirrius

# The `make clean` command cleans
# all object, build, and test files
# and removes the executable file.

.PHONY: clean
clean:
	rm -rf vendor
	$(GO) clean -i `glide novendor`
	find . -name '*.test' -type f -exec rm -f {} \;
	find . -name '*.cover' -type f -exec rm -f {} \;

# The `make glide` command ensures that
# all imported dependencies are synced
# and located within the vendor folder.

.PHONY: glide
glide:
	glide install

# The `make setup` command runs the
# go generate command in all of the
# project subdirectories.

.PHONY: setup
setup:
	$(GO) generate -v `glide novendor`

# The `make patch` command applies
# any git patches to the necessary
# source 3rd party source files.

.PHONY: patch
patch:
	git apply build/patch.txt

# The `make build` command compiles
# the build flags, gets the project
# dependencies, and runs a build.

.PHONY: build
build: LDF += $(shell GOPATH=${GOPATH} build/flags.sh)
build:
	$(GO) build -v -tags 'pdflib8' -ldflags '$(LDF)'

# The `make install` command compiles
# the build flags, gets the project
# dependencies, and runs an install.

.PHONY: install
install: LDF += $(shell GOPATH=${GOPATH} build/flags.sh)
install:
	$(GO) install -v -tags 'pdflib8' -ldflags '$(LDF)'

# The `make tests` command runs all
# tests, found within all sub-folders
# in the project folder.

.PHONY: tests
tests:
	$(GO) test -v `glide novendor`

# The `make cover` command runs all
# tests, and produces and uploads a
# coverage profile to coveralls.

.PHONY: cover
cover:
	echo 'mode: atomic' > main.cover
	glide novendor | cut -d '/' -f-2 | xargs -I % sh -c 'touch temp.cover; go test -covermode=count -coverprofile=temp.cover %; tail -n +2 temp.cover >> main.cover; rm temp.cover;'
	goveralls -coverprofile=./main.cover -service=circle-ci -repotoken=${COVERALLS}
