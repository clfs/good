// Package uci implements the UCI protocol.
package uci

import (
	"bufio"
	"fmt"
	"io"
)

// Client is a client that communicates over UCI.
type Client struct {
	r io.Reader
	w io.Writer
}

// New returns a new client.
func New(r io.Reader, w io.Writer) *Client {
	return &Client{r, w}
}

// Run runs the client.
func (c *Client) Run() error {
	s := bufio.NewScanner(c.r)
	for s.Scan() {
		line := s.Text()
		if err := c.dispatch(line); err != nil {
			return err
		}
	}
	return s.Err()
}

func (c *Client) dispatch(line string) error {
	fmt.Fprintf(c.w, "echo: %s\n", line) // todo
	return nil
}
