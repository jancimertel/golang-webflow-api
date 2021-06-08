package gowebflow

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jancimertel/gowebflow/request"
	"github.com/jancimertel/gowebflow/response"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	baseUrl    = "https://api.webflow.com"
	apiVersion = "1.0.0"
)

// webflowClient provides api calls as public methods
type webflowClient struct {
	token   string
	baseUrl string
	client  http.Client
}

// request makes a request to webflowClient's API
func (m *webflowClient) request(requestData request.Envelope, responseData interface{}) error {
	bytesData, err := json.Marshal(requestData.Body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(string(requestData.Method), m.baseUrl+requestData.Path, bytes.NewReader(bytesData))
	if err != nil {
		return fmt.Errorf("could not create request: %s", err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", m.token))
	req.Header.Add("Accept-Version", apiVersion)

	res, err := m.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	rawResponse, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// in case of successful request - unmarshal to expected container
	if res.StatusCode >= http.StatusOK && res.StatusCode < http.StatusMultipleChoices {
		if err = json.Unmarshal(rawResponse, responseData); err != nil {
			return err
		}

		return nil
	}

	// in case of unsuccessful request - unmarshal to common error container
	var errData response.Error
	if err = json.Unmarshal(rawResponse, &errData); err != nil {
		return err
	}

	return fmt.Errorf("api returned an error (%d): %v", errData.Code, errData.Name)
}

// GetSites returns list of sites associated with the curernt account
// https://developers.webflow.com/#list-sites
func (m *webflowClient) GetSites() ([]response.Site, error) {
	var data []response.Site
	err := m.request(request.Envelope{
		Method: request.MethodGet,
		Path:   "/sites",
		Body:   nil,
	}, &data)

	return data, err
}

// GetCollections returns list of collections for specific site
// https://developers.webflow.com/#collections
func (m *webflowClient) GetCollections(siteId string) ([]response.Collection, error) {
	var data []response.Collection
	err := m.request(request.Envelope{
		Method: request.MethodGet,
		Path:   fmt.Sprintf("/sites/%s/collections", siteId),
		Body:   nil,
	}, &data)

	return data, err
}

// NewClient returns new instance for the client structure
func NewClient(secret string) (*webflowClient, error) {
	if secret == "" {
		return nil, errors.New("missing webflow authentication token")
	}
	return &webflowClient{
		token:   secret,
		baseUrl: baseUrl,
		client: http.Client{
			Timeout: time.Second * 10,
		},
	}, nil
}
