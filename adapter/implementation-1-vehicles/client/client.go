package client

import "fmt"

type Transport interface {
	NavigateToDestination()
}

type Client struct{}

func (t *Client) StartNavigation(trans Transport) {
	fmt.Println("client starting the navigation process.")
	trans.NavigateToDestination()
}
