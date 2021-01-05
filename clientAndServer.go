package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Request struct {
	q string
	c chan<- Response
}

type Response struct {
	q        string
	answer   string
	duration time.Duration
	serverId int
}

type ClientResponse struct {
	question string
	answer   string
}

var data = []string{
	"question1",
	"question2",
	"question3",
	"question4",
}

var answers = map[string]string{
	"question1": "ans1",
	"question2": "ans2",
	"question3": "ans3",
	"question4": "ans4",
}

const N = 5
const MAX_SERVER_TIME = 100

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	qCh := client(data)
	ansCh := hub(qCh)

	for a := range ansCh {
		fmt.Printf(`Your question: "%s" and server response: "%s"` + "\n", a.question, a.answer)
	}

}

func client(data []string) <-chan string {
	out := make(chan string)

	go func() {
		for _, q := range data {
			out <- q
		}
		close(out)
	}()

	return out
}

func hub(qCh <-chan string) <-chan ClientResponse {
	out := make(chan ClientResponse)
	var wg sync.WaitGroup

	go func() {
		for q := range qCh {
			ansCh := make(chan Response)
			req := Request{q: q, c: ansCh}

			for i := 0; i < N; i++ {
				go server(i, req)
			}
			wg.Add(1)
			go handler(out, q, ansCh, &wg)
		}
		wg.Wait()
		close(out)
	}()

	return out
}

func handler(out chan<- ClientResponse, question string, myAnsCh <-chan Response, wg *sync.WaitGroup) {
	timeout := time.After(MAX_SERVER_TIME / 4 * time.Nanosecond)
	select {
	case ans := <-myAnsCh:
		out <- ClientResponse{question: ans.q, answer: ans.answer}
	case <-timeout:
		out <- ClientResponse{question: question, answer: "No Server Response"}
	}
	wg.Done()
}

func server(id int, req Request) {
	d := time.Duration(rand.Intn(MAX_SERVER_TIME))
	time.Sleep(d * time.Nanosecond)
	answer := answers[req.q]
	resp := Response{q: req.q, answer: answer, duration: d, serverId: id}
	req.c <- resp
}
