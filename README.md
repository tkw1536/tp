# tp -- Tiny Reverse Proxy

[![Build Status](https://travis-ci.com/tkw1536/tp.svg?branch=master)](https://travis-ci.com/tkw1536/tp)
[![Docker Hub](https://img.shields.io/docker/automated/tkw01536/tp)](https://hub.docker.com/r/tkw01536/tp/)

This repository contains a tiny http reverse proxy written in go. 
All it can do is reverse proxy unconditionally to a url given as an environment variable. 

This is intended to be used inside of Docker, and can be found as on DockerHub as [tkw01536/tp](https://hub.docker.com/r/tkw01536/tp/) as an automated build. 
To start it up run:

```
docker run -e TARGET=http://example.com -p 8080:8080 tkw01536/tp
```

The code is licensed under the Unlicense, hence in the public domain. 