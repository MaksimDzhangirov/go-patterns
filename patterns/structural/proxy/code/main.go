package main

import (
	"fmt"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/structural/proxy/code/app"
)

func main() {
	nginxServer := app.NewNginxServer()
	appStatusURL := "/app/status"
	createUserURL := "/create/user"
	httpCode, body := nginxServer.HandleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)
	httpCode, body = nginxServer.HandleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)
	httpCode, body = nginxServer.HandleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)
	httpCode, body = nginxServer.HandleRequest(createUserURL, "POST")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", createUserURL, httpCode, body)
	httpCode, body = nginxServer.HandleRequest(createUserURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", createUserURL, httpCode, body)
}
