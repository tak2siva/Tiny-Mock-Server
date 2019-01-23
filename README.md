## Tiny mock server

Instantly run a container server with stub response code & content on given port.

## Usage

```
docker run -e PORT=9091 -e CONTENT='hello' -e CODE=200 -p 9091:9091 -it tak2siva/tiny-mock-server 
```

```
curl -i http://localhost:9091/

HTTP/1.1 200 OK
Date: Sun, 23 Sep 2018 13:06:14 GMT
Content-Length: 5
Content-Type: text/plain; charset=utf-8

hello
```

## Args
* PORT (default 9901)
* CODE (default 200)
* CONTENT
* PING_HOST         ## poll the url in given interval
* PING_INTERVAL     ## default 1 second
* CALLBACK_URL      ## each requests can have a callback url (also propagates headers)


