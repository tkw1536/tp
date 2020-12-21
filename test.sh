#!/bin/bash
set -e

STATUS="PASS"

# check that the content is right
function assert_content() {
    GOT=`curl $1`;
    EXPECTED="$2";
    
    if [ "$GOT" == "$EXPECTED" ]; then
        echo -e "\033[0;32m[PASS]\033[0m $1: $GOT";
    else
        echo -e "\033[0;31m[FAIL]\033[0m $1: $GOT (expected $EXPECTED)"; 
        STATUS="FAIL";
    fi;
}

# prepare a testing network
docker network create smoke

# build the echo server
cd testserver
docker build -t tkw1536/echo .
docker run -d --name=echosmoke --network=smoke -p 8081:8080 tkw1536/echo
sleep 10

# selfcheck the echo server
assert_content "http://localhost:8081/" "/" 
assert_content "http://localhost:8081/path/" "/path/"

# stop the echoserver
docker stop echosmoke
cd ..

# Build IMAGE_NAME and start it
IMAGE_NAME="$1"
docker build -t "$IMAGE_NAME" .
docker run -d --name=smoke --network=smoke -p 8080:8080 -e TARGET=http://echosmoke:8080 "$IMAGE_NAME"
sleep 10

# the target isn't up => execept it to be white
assert_content "http://localhost:8080/" ""
assert_content "http://localhost:8080/path/" ""

# bring up the echoserver
docker start echosmoke
sleep 10

# echo should echo the path now
assert_content "http://localhost:8080/" "/" 
assert_content "http://localhost:8080/path/" "/path/"

# stop everything
docker stop smoke && docker rm smoke > /dev/null
docker stop echosmoke && docker rm echosmoke  > /dev/null
docker network rm smoke  > /dev/null

[ "$STATUS" == "PASS" ]