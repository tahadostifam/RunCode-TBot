#!/usr/bin/env bash

docker build -t exec_code_bot_server .

docker run -dp 22:22 exec_code_bot_server

