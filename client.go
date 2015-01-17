package zendesk

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/savaki/httpctx"
	"golang.org/x/net/context"
)

var (
	DomainNotSetErr = fmt.Errorf("ZENDESK_DOMAIN not set")
	EmailNotSetErr  = fmt.Errorf("ZENDESK_EMAIL not set")
	TokenNotSetErr  = fmt.Errorf("ZENDESK_TOKEN not set")
)

type Client struct {
	domain string
	http   httpctx.HttpClient
}

func (c *Client) Users() *UserApi {
	return &UserApi{
		client:  c,
		context: context.Background(),
	}
}

func (c *Client) toFullUrl(path string) string {
	return fmt.Sprintf("https://%v.zendesk.com%s", c.domain, path)
}

func (c *Client) get(ctx context.Context, path string, params *url.Values, v interface{}) error {
	return c.http.Get(ctx, c.toFullUrl(path), params, v)
}

func (c *Client) post(ctx context.Context, path string, payload, v interface{}) error {
	return c.http.Post(ctx, c.toFullUrl(path), payload, v)
}

func (c *Client) put(ctx context.Context, path string, payload, v interface{}) error {
	return c.http.Put(ctx, c.toFullUrl(path), payload, v)
}

func (c *Client) delete(ctx context.Context, path string, v interface{}) error {
	return c.http.Do(ctx, "DELETE", c.toFullUrl(path), nil, nil, v)
}

func NewFromEnv() (*Client, error) {
	domain := os.Getenv("ZENDESK_DOMAIN")
	email := os.Getenv("ZENDESK_EMAIL")
	token := os.Getenv("ZENDESK_TOKEN")

	if domain == "" {
		return nil, DomainNotSetErr
	}
	if email == "" {
		return nil, EmailNotSetErr
	}
	if token == "" {
		return nil, TokenNotSetErr
	}

	return New(domain, email, token), nil
}

func New(domain, email, token string) *Client {
	username := fmt.Sprintf("%s/token", email)
	password := token
	authFunc := func(req *http.Request) *http.Request {
		req.SetBasicAuth(username, password)
		return req
	}

	return &Client{
		domain: domain,
		http:   httpctx.WithAuthFunc(authFunc),
	}
}
