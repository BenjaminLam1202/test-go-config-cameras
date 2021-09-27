package client

import (
	"fmt"
	digestauth "github.com/fedge247/go-config-cameras/hkvision/digest-auth"
	"github.com/fedge247/go-config-cameras/hkvision/types"
	"net/url"
	"path"
)

/**
 * @author : Donald Trieu
 * @created : 9/24/21, Friday
**/
const  (
	DefaultProtocol = "http"
	DefaultHost = "0.0.0.0"
)
type Client struct {
	proto string
	host string
	path string
	username string
	password string
	client *digestauth.DigestTransport
}

func NewClient(opts types.ConstructClient) (*Client, error) {
	client := digestauth.NewTransport(opts.Username, opts.Password)
	c := &Client{
		proto: opts.Proto,
		host: opts.Host,
		path: opts.Path,
		username: opts.Username,
		password: opts.Password,
		client: &client,
	}
	if opts.Proto == "" {
		c.proto = DefaultProtocol
	}
	if opts.Host == ""{
		c.host = DefaultHost
	}
	return c, nil
}

func (cli *Client) getAPIPath(p string, query url.Values) string {
	var apiPath = path.Join(cli.path, p)
	return (&url.URL{Path: apiPath, RawQuery: query.Encode()}).String()
}

func (cli *Client) getURL() string {
	if cli.proto != "" {
		cli.proto = DefaultProtocol
	}
	if cli .host != ""{
		cli.host = DefaultHost
	}
	return fmt.Sprintf("%s://%s", cli.proto, cli.host)
}



