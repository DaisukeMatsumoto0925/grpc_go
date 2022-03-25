#!/bin/bash

protoc blogpb/blog.proto --go_out=. --go-grpc_out=.

