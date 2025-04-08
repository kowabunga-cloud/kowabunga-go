# Copyright (c) The Kowabunga Project
# Apache License, Version 2.0 (see LICENSE or https://www.apache.org/licenses/LICENSE-2.0.txt)
# SPDX-License-Identifier: Apache-2.0

.DEFAULT_GOAL := sdk

##########
# Common #
##########

V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

ifeq ($(V), 1)
  OUT = ""
else
  OUT = ">/dev/null"
endif

######################
# Build Dependencies #
######################

NODE_DIR = node_modules

YARN = $(NODE_DIR)/.bin/yarn
GENERATOR = $(NODE_DIR)/.bin/openapi-generator-cli

.PHONY: get-yarn
get-yarn: ; $(info $(M) [Npm] Installing yarn…) @
	$Q test -x $(YARN) || npm install --silent yarn

.PHONY: get-openapi-generator
get-openapi-generator: get-yarn ;$(info $(M) [Yarn] Installing openapi-generator-cli…) @
	$Q test -x $(GENERATOR) || $(YARN) add --silent @openapitools/openapi-generator-cli
	$Q chmod a+x $(GENERATOR)

#######
# SDK #
#######

BUILD_DIR = build

# use "heads/master" to build from latest
SDK_OPENAPI_VERSION = "tags/v0.52.5"
SDK_OPENAPI_SPEC = "https://raw.githubusercontent.com/kowabunga-cloud/openapi/refs/$(SDK_OPENAPI_VERSION)/openapi.yaml"
SDK_GENERATOR = go
SDK_PKG_NAME = kowabunga

.PHONY: sdk
sdk: get-openapi-generator ; $(info $(M) [OpenAPIv3] Generate Golang SDK client code…) @
	$Q $(GENERATOR) generate \
	  -g $(SDK_GENERATOR) \
	  --package-name $(SDK_PKG_NAME) \
	  --openapi-normalizer KEEP_ONLY_FIRST_TAG_IN_OPERATION=true \
	  -p withGoMod=false \
	  -i $(SDK_OPENAPI_SPEC) \
	  -o $(BUILD_DIR) \
	  $(OUT)
	 $Q rm -f $(BUILD_DIR)/.gitignore
	 $Q rm -rf $(BUILD_DIR)/.openapi-generator
	 $Q rm -f $(BUILD_DIR)/.openapi-generator-ignore
	 $Q rm -f $(BUILD_DIR)/.travis.yml
	 $Q rm -rf $(BUILD_DIR)/api
	 $Q rm -f $(BUILD_DIR)/git_push.sh
	 $Q rm -rf $(BUILD_DIR)/README.md

.PHONY: all
all: sdk ; @
	$Q echo "done"
