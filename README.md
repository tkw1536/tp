# tp -- Tiny Reverse Proxy

![CI Status](https://github.com/tkw1536/tp/workflows/CI/badge.svg)

This repository contains a tiny http reverse proxy written in go. 
All it can do is reverse proxy unconditionally to a url given as an environment variable. 

The code is licensed under the Unlicense, hence in the public domain. 

This is intended to be used inside of Docker, and can be found as [a GitHub Package](https://github.com/users/tkw1536/packages/container/package/tp). 
To start it up run:

```
docker run -e TARGET=http://example.com -p 8080:8080 tkw01536/tp
```

For legacy reasons this image is also available on DockerHub as the automated build [tkw01536/tr](https://hub.docker.com/r/tkw01536/tp/). 

The code is licensed under the Unlicense, hence in the public domain. 
