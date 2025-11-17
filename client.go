package annette

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	uri     *url.URL
	headers http.Header
}

func New(uri *url.URL) *Client {
	return &Client{
		uri:     uri,
		headers: http.Header{},
	}
}

func (c *Client) doWithBody(ctx context.Context, method, body string) (*Response, error) {
	sr := strings.NewReader(body)
	req, err := http.NewRequestWithContext(ctx, method, c.uri.String(), sr)
	if err != nil {
		return nil, err
	}
	req.Header = c.headers
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return &Response{res: res}, nil
}

func (c *Client) doWithoutBody(ctx context.Context, method string) (*Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, c.uri.String(), strings.NewReader(""))
	if err != nil {
		return nil, err
	}
	req.Header = c.headers
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return &Response{res: res}, nil
}

func (c *Client) Get() (*Response, error) {
	ctx := context.Background()
	return c.doWithoutBody(ctx, http.MethodGet)
}

func (c *Client) Head() (*Response, error) {
	ctx := context.Background()
	return c.doWithoutBody(ctx, http.MethodHead)
}

func (c *Client) Post(body string) (*Response, error) {
	ctx := context.Background()
	return c.doWithBody(ctx, http.MethodPost, body)
}

func (c *Client) Put(body string) (*Response, error) {
	ctx := context.Background()
	return c.doWithBody(ctx, http.MethodPut, body)
}

func (c *Client) Patch(body string) (*Response, error) {
	ctx := context.Background()
	return c.doWithBody(ctx, http.MethodPatch, body)
}

func (c *Client) Delete() (*Response, error) {
	ctx := context.Background()
	return c.doWithoutBody(ctx, http.MethodDelete)
}

func (c *Client) Options() (*Response, error) {
	ctx := context.Background()
	return c.doWithoutBody(ctx, http.MethodOptions)
}

func (c *Client) Trace() (*Response, error) {
	ctx := context.Background()
	return c.doWithoutBody(ctx, http.MethodTrace)
}

func (c *Client) GetHeader(key string) string {
	return c.headers.Get(key)
}

func (c *Client) AddHeader(key, value string) {
	c.headers.Add(key, value)
}

func (c *Client) DelHeader(key string) {
	c.headers.Del(key)
}

func (c *Client) SetHeader(key, value string) {
	c.headers.Set(key, value)
}
