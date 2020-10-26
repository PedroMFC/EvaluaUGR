#!/bin/sh
set -ex
git clone https://github.com/go-task/task
cd task
go install -v ./cmd/task