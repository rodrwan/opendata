package transapi

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	invalidStopNumber = "Paradero no válido"
)

var (
	errInvalidStopNumber   = errors.New("Invalid stop number")
	errInternalServerError = errors.New("Internal server error")
	baseURL                = url.URL{
		Scheme: "https",
		Host:   "api.adderou.cl",
		Path:   "ts/",
	}
)

// Response represent api message response.
type Response struct {
	ID       string     `json:"id"`
	Time     string     `json:"horaConsulta"`
	Message  string     `json:"descripcion"`
	Services []*Service `json:"servicios"`
}

// Service represent a bus service.
type Service struct {
	Valid     int8   `json:"valido"`
	Service   string `json:"servicio"`
	BusPatent string `json:"patente"`
	Time      string `json:"tiempo"`
	Distance  string `json:"distancia"`
}

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

// Do execute an http request.
func (c *Client) Do(req *http.Request) (*Response, error) {
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

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	if response.Message == invalidStopNumber {
		return nil, errInvalidStopNumber
	}

	return &response, nil
}

// Get http GET method.
func (c *Client) Get(stop string) (*Response, error) {
	req, err := http.NewRequest("GET", c.setURL(""), nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("paradero", stop)
	req.URL.RawQuery = q.Encode()
	return c.Do(req)

}

func (c *Client) setURL(path string) string {
	return c.baseURL
}
