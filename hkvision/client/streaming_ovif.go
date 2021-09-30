package client

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	// streaming "github.com/BenjaminLam1202/test-go-config-cameras/hkvision/types/streaming"

	"github.com/BenjaminLam1202/test-go-config-cameras/hkvision/types/onvif"
)

/**
 * @author : Benjamin Lam
 * @created : 9/29/21, Friday
**/
func (cli *Client) GetIntegrate() (onvif.Integrate, error) {
	var req http.Request
	req.URL = &url.URL{Scheme: cli.proto, Host: cli.host, Path: cli.getAPIPath("/ISAPI/System/Network/Integrate", nil)}
	req.Method = http.MethodGet
	var resp *http.Response
	var err error
	if resp, err = cli.client.RoundTrip(&req); err != nil {
		return onvif.Integrate{}, err
	}
	var response onvif.Integrate
	err = xml.NewDecoder(resp.Body).Decode(&response)
	if nil != err {
		fmt.Println("Error unmarshalling from XML", err)
		return onvif.Integrate{}, err
	}
	return response, err
}

/*
It is used to update the properties of streaming time for the device.
*/
func (cli *Client) PutIntegrate(ovif onvif.Integrate) error {
	var req http.Request
	req.URL = &url.URL{Scheme: cli.proto, Host: cli.host, Path: cli.getAPIPath("/ISAPI/System/Network/Integrate", nil)}
	req.Method = http.MethodPut
	req.Header = map[string][]string{
		"Content-Type": {"text/xml"},
	}
	var err error
	xmlData, err := xml.Marshal(ovif)
	if err != nil {
		return err
	}
	req.Body = ioutil.NopCloser(bytes.NewReader(xmlData))
	if _, err = cli.client.RoundTrip(&req); err != nil {
		return err
	}
	return err
}
