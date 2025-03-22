#!/bin/sh

echo "Waiting for balancer service to be ready..."
sleep 5

TEST_URL=${TEST_URL:-"http://s1.origin-cluster/video/123/xcg2djHckad.m3u8"}

echo "Starting load tests..."
ghz --insecure \
  --proto /protos/balancer.proto \
  --call balancer.VideoBalancer.GetVideoURL \
  -d '{"video": "http://s1.origin-cluster/video/123/xcg2djHckad.m3u8"}' \
  -c 100 \
  -n 30000 \
  balancer:50051

echo "Tests completed."