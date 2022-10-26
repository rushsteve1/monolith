#!/usr/bin/env bash

cd $(dirname "$1")/overseer
fd . '../' | entr -r go run . ../config.test.json