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
	pageSize   = 20
)

// WebflowClient provides api calls as public methods
type WebflowClient struct {
	token    string
	baseUrl  string
	client   http.Client
	pageSize uint
}

// request makes a request to WebflowClient's API
func (m *WebflowClient) request(requestData request.Envelope, responseData interface{}) error {
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
func (m *WebflowClient) GetSites() ([]response.Site, error) {
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
func (m *WebflowClient) GetCollections(siteId string) ([]response.Collection, error) {
	var data []response.Collection
	err := m.request(request.Envelope{
		Method: request.MethodGet,
		Path:   fmt.Sprintf("/sites/%s/collections", siteId),
		Body:   nil,
	}, &data)

	return data, err
}

// GetItems returns list of items from specified collection
// https://developers.webflow.com/#get-all-items-for-a-collection
func (m *WebflowClient) GetItems(collectionId string, limit uint, offset uint, optionalType ...interface{}) (interface{}, error) {
	var data response.GenericItems
	err := m.request(request.Envelope{
		Method: request.MethodGet,
		Path:   fmt.Sprintf("/collections/%s/items?limit=%d&offset=%d", collectionId, limit, offset),
		Body:   nil,
	}, &data)

	if err != nil {
		return nil, err
	}

	if len(optionalType) > 0 {
		err = json.Unmarshal(data.Items, optionalType[0])
		return optionalType[0], err
	} else {
		var genericItems []response.Item
		err = json.Unmarshal(data.Items, &genericItems)
		return genericItems, err
	}

}

// PaginateItems wraps GetItems method for easier paginating
// first page starts with 0
func (m *WebflowClient) PaginateItems(collectionId string, page uint, optionalType ...interface{}) (interface{}, error) {
	return m.GetItems(collectionId, m.pageSize, page*m.pageSize, optionalType...)
}

// NewClient returns new instance for the client structure
func NewClient(secret string) (*WebflowClient, error) {
	if secret == "" {
		return nil, errors.New("missing webflow authentication token")
	}
	return &WebflowClient{
		token:   secret,
		baseUrl: baseUrl,
		client: http.Client{
			Timeout: time.Second * 10,
		},
		pageSize: pageSize,
	}, nil
}
