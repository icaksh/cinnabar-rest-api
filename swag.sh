#!/usr/bin/bash

set PATH (go env GOPATH)/bin:$PATH
swag init --parseDependency --parseInternal -d controllers/public -instanceName public
swag init --parseDependency --parseInternal -d controllers/private -instanceName private