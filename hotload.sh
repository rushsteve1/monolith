#!/usr/bin/env bash

cd $(dirname "$1")
fd | entr -r go run rushsteve1.us/monolith/overseer config.test.json
