#!/bin/bash

pushd app ; go test -v &>../test.txt; popd