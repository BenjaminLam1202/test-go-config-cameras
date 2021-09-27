package client

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/BenjaminLam1202/test-go-config-cameras/hkvision/types/streaming"
)

/**
 * @author : Donald Trieu
 * @created : 9/24/21, Friday
**/

/*
It is used to get the properties of streaming channels for the device.
*/
func (cli *Client) GetStreamChannels() ([]streaming.StreamChannel, error) {
	var req http.Request
	req.URL = &url.URL{Scheme: cli.proto, Host: cli.host, Path: cli.getAPIPath("/ISAPI/Streaming/channels", nil)}
	req.Method = http.MethodGet
	var resp *http.Response
	var err error
	if resp, err = cli.client.RoundTrip(&req); err != nil {
		return nil, err
	}
	var response streaming.StreamChannelList
	err = xml.NewDecoder(resp.Body).Decode(&response)
	if nil != err {
		fmt.Println("Error unmarshalling from XML", err)
		return nil, err
	}
	return response.StreamChannel, err
}

/*
It is used to update the properties of streaming channels for the device.
*/
func (cli *Client) PutStreamChannels(streamChannelList []streaming.StreamChannel) error {
	var req http.Request
	req.URL = &url.URL{Scheme: cli.proto, Host: cli.host, Path: cli.getAPIPath("/ISAPI/Streaming/channels", nil)}
	req.Method = http.MethodPut
	req.Header = map[string][]string{
		"Content-Type": {"text/xml"},
	}
	var err error
	xmlData, err := xml.Marshal(streamChannelList)
	if err != nil {
		return err
	}
	req.Body = ioutil.NopCloser(bytes.NewReader(xmlData))
	if _, err = cli.client.RoundTrip(&req); err != nil {
		return err
	}
	return err
}

/*
It is used to add a streaming channel for the device.
*/
func (cli *Client) AddStreamChannels(streamChannelList streaming.StreamChannel) error {
	var req http.Request
	req.URL = &url.URL{Scheme: cli.proto, Host: cli.host, Path: cli.getAPIPath("/ISAPI/Streaming/channels", nil)}
	req.Method = http.MethodPost
	req.Header = map[string][]string{
		"Content-Type": {"text/xml"},
	}
	var err error
	xmlData, err := xml.Marshal(streamChannelList)
	if err != nil {
		return err
	}
	req.Body = ioutil.NopCloser(bytes.NewReader(xmlData))
	if _, err = cli.client.RoundTrip(&req); err != nil {
		return err
	}
	return err
}

/*
It is used to delete the list of streaming channels for the device.
*/
func (cli *Client) DeleteStreamChannels() error {
	var req http.Request
	req.URL = &url.URL{Scheme: cli.proto, Host: cli.host, Path: cli.getAPIPath("/ISAPI/Streaming/channels", nil)}
	req.Method = http.MethodDelete
	req.Header = map[string][]string{
		"Content-Type": {"text/xml"},
	}
	var err error
	if _, err = cli.client.RoundTrip(&req); err != nil {
		return err
	}
	return err
}

/*
It is used to get the properties of a particular streaming channel for the device.
*/
func (cli *Client) GetSingleStreamChannel(id string) (streaming.StreamChannel, error) {
	var req http.Request
	req.URL = &url.URL{Scheme: cli.proto, Host: cli.host, Path: cli.getAPIPath(fmt.Sprintf("/ISAPI/Streaming/channels/%s", id), nil)}
	req.Method = http.MethodGet
	var resp *http.Response
	var err error
	if resp, err = cli.client.RoundTrip(&req); err != nil {
		return streaming.StreamChannel{}, err
	}
	var response streaming.StreamChannel
	err = xml.NewDecoder(resp.Body).Decode(&response)
	if nil != err {
		fmt.Println("Error unmarshalling from XML", err)
		return streaming.StreamChannel{}, err
	}
	return response, err
}

/*
It is used to update the properties of a particular streaming channel for the device.
*/
func (cli *Client) UpdateSingleStreamChannel(id string, streamChannel streaming.StreamChannel) error {
	var req http.Request
	req.URL = &url.URL{Scheme: cli.proto, Host: cli.host, Path: cli.getAPIPath(fmt.Sprintf("/ISAPI/Streaming/channels/%s", id), nil)}
	req.Method = http.MethodPut
	req.Header = map[string][]string{
		"Content-Type": {"text/xml"},
	}
	var resp *http.Response
	var err error
	xmlData, err := xml.Marshal(streamChannel)
	if err != nil {
		return err
	}
	req.Body = ioutil.NopCloser(bytes.NewReader(xmlData))
	if resp, err = cli.client.RoundTrip(&req); err != nil {
		return err
	}
	var response streaming.StreamChannel
	err = xml.NewDecoder(resp.Body).Decode(&response)
	if nil != err {
		fmt.Println("Error unmarshalling from XML", err)
		return err
	}
	return err
}

/*
It is used to delete a particular streaming channel for the device.
*/
func (cli *Client) DeleteSingleStreamChannel(id string) error {
	var req http.Request
	req.URL = &url.URL{Scheme: cli.proto, Host: cli.host, Path: cli.getAPIPath(fmt.Sprintf("/ISAPI/Streaming/channels/%s", id), nil)}
	req.Method = http.MethodDelete
	req.Header = map[string][]string{
		"Content-Type": {"text/xml"},
	}
	var err error
	if _, err = cli.client.RoundTrip(&req); err != nil {
		return err
	}
	return err
}
