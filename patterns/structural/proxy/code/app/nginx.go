package app

import "github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/structural/proxy/code/interfaces"

type nginx struct {
	application       interfaces.Server
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func NewNginxServer() *nginx {
	return &nginx{
		application:       &application{},
		maxAllowedRequest: 2,
		rateLimiter:       make(map[string]int),
	}
}

func (n *nginx) HandleRequest(url string, method string) (int, string) {
	allowed := n.checkRateLimiting(url)
	if !allowed {
		return 403, "Not Allowed"
	}
	return n.application.HandleRequest(url, method)
}

func (n *nginx) checkRateLimiting(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	}
	if n.rateLimiter[url] > n.maxAllowedRequest {
		return false
	}
	n.rateLimiter[url] = n.rateLimiter[url] + 1
	return true
}
