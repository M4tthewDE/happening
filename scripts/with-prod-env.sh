#!/usr/bin/env bash
test -f .prod-env && export $(grep -v '^#' .prod-env | xargs -d '\n')
command "$@"
