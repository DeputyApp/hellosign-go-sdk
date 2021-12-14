package hellosign

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/DeputyApp/hellosign-go-sdk/model"
	"mime/multipart"
	"net/http"
	"strings"
	"time"
)

func (m *Client) get(path string) (*http.Response, error) {
	endpoint := fmt.Sprintf("%s%s", m.getEndpoint(), path)

	var b bytes.Buffer
	request, _ := http.NewRequest("GET", endpoint, &b)
	request.SetBasicAuth(m.APIKey, "")

	response, err := m.getHTTPClient().Do(request)
	if err != nil {
		return nil, err
	}

	return response, err
}

func (m *Client) post(path string, params *bytes.Buffer, w multipart.Writer) (*http.Response, error) {
	return m.request("POST", path, params, w)
}

func (m *Client) request(method string, path string, params *bytes.Buffer, w multipart.Writer) (*http.Response, error) {
	endpoint := fmt.Sprintf("%s%s", m.getEndpoint(), path)
	request, _ := http.NewRequest(method, endpoint, params)
	request.Header.Add("Content-Type", w.FormDataContentType())
	request.SetBasicAuth(m.APIKey, "")

	response, err := m.getHTTPClient().Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode >= 400 {
		msg := fmt.Sprintf("hellosign request failed with status %d", response.StatusCode)
		e := &model.ErrorResponse{}
		json.NewDecoder(response.Body).Decode(e)
		if e.Error != nil {
			msg = fmt.Sprintf("%s: %s", e.Error.Name, e.Error.Message)
		} else {
			messages := []string{}
			for _, w := range e.Warnings {
				messages = append(messages, fmt.Sprintf("%s: %s", w.Name, w.Message))
			}
			msg = strings.Join(messages, ", ")
		}

		return response, errors.New(msg)
	}

	return response, err
}

func (m *Client) nakedPost(path string) (*http.Response, error) {
	endpoint := fmt.Sprintf("%s%s", m.getEndpoint(), path)
	var b bytes.Buffer
	request, _ := http.NewRequest("POST", endpoint, &b)
	request.SetBasicAuth(m.APIKey, "")

	response, err := m.getHTTPClient().Do(request)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (m *Client) getEndpoint() string {
	var url string
	if m.BaseURL != "" {
		url = m.BaseURL
	} else {
		url = baseURL
	}
	return url
}

func (m *Client) getHTTPClient() *http.Client {
	var httpClient *http.Client
	if m.HTTPClient != nil {
		httpClient = m.HTTPClient
	} else {
		httpClient = &http.Client{
			Timeout: time.Second * 7,
		}
	}
	return httpClient
}
