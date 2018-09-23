## Tiny mock server

Instantly run a container server with stub response code & content on given port.

## Usage

```
docker run -e PORT=9091 -e CONTENT='hello' -e CODE=200 -p 9091:9091 -it tak2siva/tiny-mock-server 
```

## Args
* PORT (default 9901)
* CODE (default 200)
* CONTENT