package main

import (
	"fmt"
	"github.com/headzoo/surf/agent"
	"gopkg.in/headzoo/surf.v1"
	"testing"
)

func TestFetch(t *testing.T) {

	bow := surf.NewBrowser()
	bow.SetUserAgent(agent.Chrome())
	err := bow.Open("http://weiwei0923.pixnet.net/blog/post/299526739-%E5%A6%82%E5%A6%82")
	if err != nil {
		panic(err)
	}

	// Outputs: "The Go Programming Language"
	fmt.Println(bow.Title())

}
