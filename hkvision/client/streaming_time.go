package client

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	// streaming "github.com/BenjaminLam1202/test-go-config-cameras/hkvision/types/streaming"

	time_streaming "github.com/BenjaminLam1202/test-go-config-cameras/hkvision/types/time"
)

/**
 * @author : Benjamin Lam
 * @created : 9/29/21, Friday
**/

/*
It is used to get the properties of streaming channels for the device.
*/
func (cli *Client) GetTimeStreamChannels() (time_streaming.Time, error) {
	var req http.Request
	req.URL = &url.URL{Scheme: cli.proto, Host: cli.host, Path: cli.getAPIPath("/ISAPI/System/time", nil)}
	req.Method = http.MethodGet
	var resp *http.Response
	var err error
	if resp, err = cli.client.RoundTrip(&req); err != nil {
		return time_streaming.Time{}, err
	}
	var response time_streaming.Time
	err = xml.NewDecoder(resp.Body).Decode(&response)
	if nil != err {
		fmt.Println("Error unmarshalling from XML", err)
		return time_streaming.Time{}, err
	}
	return response, err
}

/*
It is used to update the properties of streaming channels for the device.
*/
func (cli *Client) PutTimeStreamChannels(time_streaming time_streaming.Time) error {
	var req http.Request
	req.URL = &url.URL{Scheme: cli.proto, Host: cli.host, Path: cli.getAPIPath("/ISAPI/System/time", nil)}
	req.Method = http.MethodPut
	req.Header = map[string][]string{
		"Content-Type": {"text/xml"},
	}
	var err error
	xmlData, err := xml.Marshal(time_streaming)
	if err != nil {
		return err
	}
	req.Body = ioutil.NopCloser(bytes.NewReader(xmlData))
	if _, err = cli.client.RoundTrip(&req); err != nil {
		return err
	}
	return err
}
