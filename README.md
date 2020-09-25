# Docker `show-headers`

```
docker pull patrickdappollonio/show-headers
```

A simple tool built in Go to show the HTTP headers. Here's an example output:

```
REQUEST INFORMATION:
--------------------
URL:                /
Method:             GET
Protocol:           HTTP/1.1
Host:               localhost:8080
Client Address:     172.17.0.1:39960

REQUEST HEADERS:
----------------
Proxy-Connection:    Keep-Alive
Connection:          keep-alive
User-Agent:          curl/7.68.0
Accept:              */*
```

The application listen by default on port `8080` but you can configure the port
number using the `$PORT` environment variable.`
