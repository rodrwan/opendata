package tyaas

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	errInternalServerError = errors.New("Internal server error")
	baseURL                = url.URL{
		Scheme: "https",
		Host:   "api.adderou.cl",
		Path:   "tyaas/",
	}
)

// Client ...
type Client struct {
	client  *http.Client
	baseURL string
}

// NewClient ...
func NewClient() *Client {
	c := &http.Client{}

	return &Client{
		client:  c,
		baseURL: baseURL.String(),
	}
}

// Horoscope ...
type Horoscope struct {
	Date        string     `json:"titulo"`
	ZodiacSigns ZodiacSign `json:"horoscopo"`
}

// ZodiacSign ...
type ZodiacSign map[string]ZodiacSignData

// UnmarshalGQL implements the graphql.Marshaler interface
func (zs *ZodiacSign) UnmarshalGQL(v interface{}) error {
	value, ok := v.(ZodiacSign)
	if !ok {
		return fmt.Errorf("points must be strings")
	}

	*zs = value
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (zs ZodiacSign) MarshalGQL(w io.Writer) {
	resp, err := json.Marshal(zs)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	w.Write(resp)
}

// ZodiacSignData ...
type ZodiacSignData struct {
	Name     string `json:"nombre"`
	SignDate string `json:"fechaSigno"`
	Love     string `json:"amor"`
	Health   string `json:"salud"`
	Money    string `json:"dinero"`
	Colour   string `json:"color"`
	Number   string `json:"numero"`
}

// Do execute an http request.
func (c *Client) Do(req *http.Request) (*Horoscope, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusInternalServerError {
		return nil, errInternalServerError
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response Horoscope
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// Get http GET method.
func (c *Client) Get() (*Horoscope, error) {
	req, err := http.NewRequest("GET", c.setURL(""), nil)
	if err != nil {
		return nil, err
	}

	return c.Do(req)

}

func (c *Client) setURL(path string) string {
	return c.baseURL
}
