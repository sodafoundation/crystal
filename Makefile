# Copyright 2019 The OpenSDS Authors.
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

BASE_DIR := $(shell pwd)
BUILD_DIR := $(BASE_DIR)/build
DIST_DIR := $(BASE_DIR)/build/dist
VERSION ?= $(shell git describe --exact-match 2> /dev/null || \
	     git describe --match=$(git rev-parse --short=8 HEAD) \
             --always --dirty --abbrev=8)
BUILD_TGT := soda-multicloud-$(VERSION)-linux-amd64

.PHONY: all build prebuild api aksk backend s3 metadata docker clean

all: build

build: api aksk backend s3 metadata

prebuild:
	mkdir -p  $(BUILD_DIR)

api: prebuild
	CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s -extldflags "-static"' -o $(BUILD_DIR)/api github.com/opensds/crsytal/api/cmd

aksk: prebuild
	CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s -extldflags "-static"' -o $(BUILD_DIR)/aksk github.com/opensds/crsytal/aksk/cmd

backend: prebuild
	CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s -extldflags "-static"' -o $(BUILD_DIR)/backend github.com/opensds/crsytal/backend/cmd

s3: prebuild
	CGO_ENABLED=1 GOOS=linux go build -ldflags '-w -s -extldflags "-dynamic"' -o $(BUILD_DIR)/s3 github.com/opensds/crsytal/s3/cmd

metadata: prebuild
	CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s -extldflags "-static"' -o $(BUILD_DIR)/metadata github.com/opensds/crsytal/metadata/cmd

docker: build

	cp $(BUILD_DIR)/api api
	chmod 755 api/api
	docker build api -t sodafoundation/crsytal-api:latest

	cp $(BUILD_DIR)/aksk aksk
	chmod 755 aksk/aksk
	docker build aksk -t sodafoundation/crsytal-aksk:latest

	cp $(BUILD_DIR)/backend backend
	chmod 755 backend/backend
	docker build backend -t sodafoundation/crsytal-backend:latest

	cp $(BUILD_DIR)/metadata metadata
	chmod 755 metadata/metadata
	docker build metadata -t sodafoundation/crsytal-metadata:latest

	cp $(BUILD_DIR)/s3 s3
	chmod 755 s3/s3
	chmod 755 s3/initdb.sh
	docker build s3 -t sodafoundation/crsytal-s3:latest

.PHONY: goimports
ifeq ($(CHECK_ERROR),y)
goimports: UNFORMATTED_FILES = $(shell goimports -local $(shell go list -m) -l $(shell go list -f {{.Dir}} ./... |grep -v /vendor/ | grep -v /proto))
goimports: 
	@ if [ ! -z "$(UNFORMATTED_FILES)" ]; then \
		echo "Formatting error in $(UNFORMATTED_FILES)"; \
		exit 1; \
	fi
else
goimports:
	goimports -local $(shell go list -m) -w $(shell go list -f {{.Dir}} ./... |grep -v /vendor/ | grep -v /proto)
endif

clean:
	rm -rf $(BUILD_DIR) api/api aksk/aksk backend/backend s3/s3

version:
	@echo ${VERSION}

.PHONY: dist
dist: build
	rm -fr $(DIST_DIR) && mkdir -p $(DIST_DIR)/$(BUILD_TGT)/bin
	cd $(DIST_DIR) && \
	cp ../api $(BUILD_TGT)/bin/ && \
	cp ../aksk $(BUILD_TGT)/bin/ && \
	cp ../backend $(BUILD_TGT)/bin/ && \
	cp ../s3 $(BUILD_TGT)/bin/ && \
	cp ../metadata $(BUILD_TGT)/bin/ && \
	cp $(BASE_DIR)/LICENSE $(BUILD_TGT) && \
	zip -r $(DIST_DIR)/$(BUILD_TGT).zip $(BUILD_TGT) && \
	tar zcvf $(DIST_DIR)/$(BUILD_TGT).tar.gz $(BUILD_TGT)
	tree $(DIST_DIR)
