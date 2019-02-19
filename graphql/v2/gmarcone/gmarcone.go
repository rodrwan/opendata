package gmarcone

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client ...
type Client struct {
	URL        string
	clientHTTP *http.Client
}

type Error struct {
	Message string `json:"error"`
}

// NewClient ...
func NewClient(URL string) *Client {
	ch := &http.Client{}

	return &Client{
		URL:        URL,
		clientHTTP: ch,
	}
}

func (c *Client) do(req *http.Request) ([]byte, error) {
	resp, err := c.clientHTTP.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var e Error
	if err := json.Unmarshal(bb, &e); err != nil {
		return nil, err
	}

	if e.Message != "" {
		return nil, errors.New(e.Message)
	}

	return bb, nil
}

// GET ...
func (c *Client) GET(path string) ([]byte, error) {
	req, err := http.NewRequest("GET", c.setURL(path), nil)
	if err != nil {
		return nil, err
	}

	return c.do(req)
}

func (c *Client) setURL(path string) string {
	return fmt.Sprintf("%s/%s", c.URL, path)
}
