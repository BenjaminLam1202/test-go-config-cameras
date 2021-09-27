package client

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"

	"github.com/BenjaminLam1202/test-go-config-cameras/hkvision/types/streaming"
)

/**
 * @author : Donald Trieu
 * @created : 9/24/21, Friday
**/

/*
	It is used to get a device streaming status.
*/
func (cli *Client) StreamingStatus() (streaming.StreamStatus, error) {
	var req http.Request
	req.URL = &url.URL{Scheme: cli.proto, Host: cli.host, Path: cli.getAPIPath("/ISAPI/Streaming/status", nil)}
	req.Method = http.MethodGet
	var resp *http.Response
	var err error
	if resp, err = cli.client.RoundTrip(&req); err != nil {
		return streaming.StreamStatus{}, err
	}
	var response streaming.StreamStatus
	err = xml.NewDecoder(resp.Body).Decode(&response)
	if nil != err {
		fmt.Println("Error unmarshalling from XML", err)
		return streaming.StreamStatus{}, err
	}
	return response, err
}
