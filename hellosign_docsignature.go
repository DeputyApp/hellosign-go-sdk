package hellosign

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/DeputyApp/hellosign-go-sdk/model"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
)

const (
	baseURL string = "https://api.hellosign.com/v3/"
)

// Client contains APIKey and optional http.client
type Client struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
}

// CreateEmbeddedSignatureRequest creates a new embedded signature
func (m *Client) CreateEmbeddedSignatureRequest(embeddedRequest model.EmbeddedSignatureRequest) (*model.SignatureRequest, error) {

	params, writer, err := m.marshalMultipartRequest(embeddedRequest)
	if err != nil {
		return nil, err
	}

	response, err := m.post("signature_request/create_embedded", params, *writer)
	if err != nil {
		return nil, err
	}

	return m.sendSignatureRequest(response)
}

// GetSignatureRequest - Gets a SignatureRequest that includes the current status for each signer.
func (m *Client) GetSignatureRequest(signatureRequestID string) (*model.SignatureRequest, error) {
	path := fmt.Sprintf("signature_request/%s", signatureRequestID)
	response, err := m.get(path)
	if err != nil {
		return nil, err
	}
	return m.sendSignatureRequest(response)
}

// GetEmbeddedSignURL - Retrieves an embedded signing object.
func (m *Client) GetEmbeddedSignURL(signatureRequestID string) (*model.SignURLResponse, error) {
	path := fmt.Sprintf("embedded/sign_url/%s", signatureRequestID)
	response, err := m.get(path)
	if err != nil {
		return nil, err
	}

	data := &model.EmbeddedSignatureResponse{}
	err = json.NewDecoder(response.Body).Decode(data)
	if err != nil {
		return nil, err
	}

	return data.Embedded, nil
}

func (m *Client) SaveFile(signatureRequestID, fileType, destFilePath string) (os.FileInfo, error) {
	bytes, err := m.GetFiles(signatureRequestID, fileType)

	out, err := os.Create(destFilePath)
	if err != nil {
		return nil, err
	}
	out.Write(bytes)
	out.Close()

	info, err := os.Stat(destFilePath)
	if err != nil {
		return nil, err
	}
	return info, nil
}

// GetPDF - Obtain a copy of the current pdf specified by the signature_request_id parameter.
func (m *Client) GetPDF(signatureRequestID string) ([]byte, error) {
	return m.GetFiles(signatureRequestID, "pdf")
}

// GetFiles - Obtain a copy of the current documents specified by the signature_request_id parameter.
// signatureRequestID - The id of the SignatureRequest to retrieve.
// fileType - Set to "pdf" for a single merged document or "zip" for a collection of individual documents.
func (m *Client) GetFiles(signatureRequestID, fileType string) ([]byte, error) {
	path := fmt.Sprintf("signature_request/files/%s", signatureRequestID)

	var params bytes.Buffer
	writer := multipart.NewWriter(&params)

	signatureIDField, err := writer.CreateFormField("file_type")
	if err != nil {
		return nil, err
	}
	signatureIDField.Write([]byte(fileType))

	emailField, err := writer.CreateFormField("get_url")
	if err != nil {
		return nil, err
	}
	emailField.Write([]byte("false"))

	response, err := m.request("GET", path, &params, *writer)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// ListSignatureRequests - Lists the SignatureRequests (both inbound and outbound) that you have access to.
func (m *Client) ListSignatureRequests() (*model.ListResponse, error) {
	path := fmt.Sprintf("signature_request/list")
	response, err := m.get(path)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	listResponse := &model.ListResponse{}
	err = json.NewDecoder(response.Body).Decode(listResponse)
	if err != nil {
		return nil, err
	}

	return listResponse, err
}

// UpdateSignatureRequest - Update an email address on a signature request.
func (m *Client) UpdateSignatureRequest(signatureRequestID string, signatureID string, email string) (*model.SignatureRequest, error) {
	path := fmt.Sprintf("signature_request/update/%s", signatureRequestID)

	var params bytes.Buffer
	writer := multipart.NewWriter(&params)

	signatureIDField, err := writer.CreateFormField("signature_id")
	if err != nil {
		return nil, err
	}
	signatureIDField.Write([]byte(signatureID))

	emailField, err := writer.CreateFormField("email_address")
	if err != nil {
		return nil, err
	}
	emailField.Write([]byte(email))

	response, err := m.post(path, &params, *writer)
	if err != nil {
		return nil, err
	}

	return m.sendSignatureRequest(response)
}

// CancelSignatureRequest - Cancels an incomplete signature request. This action is not reversible.
func (m *Client) CancelSignatureRequest(signatureRequestID string) (*http.Response, error) {
	path := fmt.Sprintf("signature_request/cancel/%s", signatureRequestID)

	response, err := m.nakedPost(path)
	if err != nil {
		return nil, err
	}

	return response, err
}

// Private Methods

func (m *Client) marshalMultipartRequest(embRequest model.EmbeddedSignatureRequest) (*bytes.Buffer, *multipart.Writer, error) {

	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	structType := reflect.TypeOf(embRequest)
	val := reflect.ValueOf(embRequest)

	for i := 0; i < val.NumField(); i++ {

		valueField := val.Field(i)
		f := valueField.Interface()
		val := reflect.ValueOf(f)
		field := structType.Field(i)
		fieldTag := field.Tag.Get("form_field")

		switch val.Kind() {
		case reflect.Map:
			for k, v := range embRequest.Metadata {
				formField, err := w.CreateFormField(fmt.Sprintf("metadata[%v]", k))
				if err != nil {
					return nil, nil, err
				}
				formField.Write([]byte(v))
			}
		case reflect.Slice:
			switch fieldTag {
			case "signers":
				for i, signer := range embRequest.Signers {
					email, err := w.CreateFormField(fmt.Sprintf("signers[%v][email_address]", i))
					if err != nil {
						return nil, nil, err
					}
					email.Write([]byte(signer.Email))

					name, err := w.CreateFormField(fmt.Sprintf("signers[%v][name]", i))
					if err != nil {
						return nil, nil, err
					}
					name.Write([]byte(signer.Name))

					if signer.Order != 0 {
						order, err := w.CreateFormField(fmt.Sprintf("signers[%v][order]", i))
						if err != nil {
							return nil, nil, err
						}
						order.Write([]byte(strconv.Itoa(signer.Order)))
					}

					if signer.Pin != "" {
						pin, err := w.CreateFormField(fmt.Sprintf("signers[%v][pin]", i))
						if err != nil {
							return nil, nil, err
						}
						pin.Write([]byte(signer.Pin))
					}
				}
			case "cc_email_addresses":
				for k, v := range embRequest.CCEmailAddresses {
					formField, err := w.CreateFormField(fmt.Sprintf("cc_email_addresses[%v]", k))
					if err != nil {
						return nil, nil, err
					}
					formField.Write([]byte(v))
				}
			case "form_fields_per_document":
				if len(embRequest.FormFieldsPerDocument) > 0 {
					formField, err := w.CreateFormField(fieldTag)
					if err != nil {
						return nil, nil, err
					}
					ffpdJSON, err := json.Marshal(embRequest.FormFieldsPerDocument)
					if err != nil {
						return nil, nil, err
					}
					formField.Write([]byte(ffpdJSON))
				}
			case "file":
				for i, path := range embRequest.File {
					file, _ := os.Open(path)

					formField, err := w.CreateFormFile(fmt.Sprintf("file[%v]", i), file.Name())
					if err != nil {
						return nil, nil, err
					}
					_, err = io.Copy(formField, file)
				}
			case "file_url":
				for i, fileURL := range embRequest.FileURL {
					formField, err := w.CreateFormField(fmt.Sprintf("file_url[%v]", i))
					if err != nil {
						return nil, nil, err
					}
					formField.Write([]byte(fileURL))
				}
			}
		case reflect.Bool:
			formField, err := w.CreateFormField(fieldTag)
			if err != nil {
				return nil, nil, err
			}
			formField.Write([]byte(m.boolToIntString(val.Bool())))
		default:
			if val.String() != "" {
				formField, err := w.CreateFormField(fieldTag)
				if err != nil {
					return nil, nil, err
				}
				formField.Write([]byte(val.String()))
			}
		}
	}

	w.Close()
	return &b, w, nil
}

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

func (m *Client) sendSignatureRequest(response *http.Response) (*model.SignatureRequest, error) {
	defer response.Body.Close()

	sigRequestResponse := &model.SignatureRequestResponse{}

	err := json.NewDecoder(response.Body).Decode(sigRequestResponse)

	sigRequest := sigRequestResponse.SignatureRequest

	return sigRequest, err
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
		httpClient = &http.Client{}
	}
	return httpClient
}

func (m *Client) boolToIntString(value bool) string {
	if value == true {
		return "1"
	}
	return "0"
}
