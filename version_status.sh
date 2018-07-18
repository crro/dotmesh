#!/bin/bash

# script to tell bazel what version to set on the client
VERSION=$(cd cmd/dm/versioner && go run versioner.go)
echo VERSION ${VERSION}
echo DOCKERTAG ${CI_DOCKER_TAG}