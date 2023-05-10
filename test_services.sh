#!/bin/bash

NUM_REQUESTS=10


GOLANG_API_URL="http://localhost:8080"
PY_API_URL="http://localhost:8081"

HEADERS="Content-Type: application/json"


DATA="{\"data\": \"test\"}"


echo "Counter,Golang Response Time (ms),Python Response Time (ms)"


for i in $(seq 1 $NUM_REQUESTS); do
    GOLANG_START=$(date +%s)
    GOLANG_NANO=$(date +%N)
    GOLANG_START_NANO=$(echo "$GOLANG_START * 1000000000 + $GOLANG_NANO" | bc)
    
    ...
    
    PY_START=$(date +%s)
    PY_NANO=$(date +%N)
    PY_START_NANO=$(echo "$PY_START * 1000000000 + $PY_NANO" | bc)
    

  echo "$i,$GOLANG_TIME,$PY_TIME"

  if [ "$GOLANG_RESPONSE" -ne "200" ] || [ "$PY_RESPONSE" -ne "200" ]; then
    echo "Error: One or more requests returned an error status code."
    exit 1
  fi


  SLEEP_TIME=$(( ( RANDOM % 3 )  + 1 ))
  sleep $SLEEP_TIME
done
