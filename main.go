package main

import (
	"fmt"
	"gopkg.in/go-playground/webhooks.v3"
	"gopkg.in/go-playground/webhooks.v3/github"
	"log"
	"net/http"
)

func main() {
	hook := github.New(&github.Config{Secret: ""})
	hook.RegisterEvents(handlePush, github.PushEvent)

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	mux.Handle("/web-hooks", webhooks.Handler(hook))
	log.Fatal(http.ListenAndServe(":8123", mux))
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func handlePush(payload interface{}, header webhooks.Header) {
	// get latest head to ./builds/repo-branch/src folder
	// point logger at ./builds/repo-branch/history/{build-num}.log
	// execute script: ./builds/repo-branch/src./fork-over.sh
	// send and email to
	p := payload.(github.PushPayload)
	fmt.Println(p.Pusher)
	fmt.Println(p)
	fmt.Println(header)
}
