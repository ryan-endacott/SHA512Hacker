#!/bin/sh

for i in `seq 1 8`; do
  echo "Launching instance $i"
  nohup go run main.go &
  sleep 1
done

echo "Done launching instances."
