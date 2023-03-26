#!/bin/bash

cd node && npm run client && cd ../
cd go && make start_client
