# single-app

This example uses Envoy to proxy requests to a Go application.

Steps to run:
1. Run `docker-compose build && docker-compose up -d`
2. Envoy will be accessed through `localhost:10000` - this will then proxy the request to the Go application running on `localhost:8080`