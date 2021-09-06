package hellosign

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/DeputyApp/hellosign-go-sdk/model"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"reflect"
	"strconv"
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

	return m.parseSignatureRequestResponse(response)
}

// GetSignatureRequest - Gets a SignatureRequest that includes the current status for each signer.
func (m *Client) GetSignatureRequest(signatureRequestID string) (*model.SignatureRequest, error) {
	path := fmt.Sprintf("signature_request/%s", signatureRequestID)
	response, err := m.get(path)
	if err != nil {
		return nil, err
	}
	return m.parseSignatureRequestResponse(response)
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

	return data.GetEmbedded(), nil
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

	return m.parseSignatureRequestResponse(response)
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

// DeleteSignatureRequest - Remove access to a completed SignatureRequest. This action is not reversible.
func (m *Client) DeleteSignatureRequest(signatureRequestID string) (*http.Response, error) {
	return m.nakedPost(fmt.Sprintf("signature_request/remove/%s", signatureRequestID))
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
			for k, v := range embRequest.GetMetadata() {
				formField, err := w.CreateFormField(fmt.Sprintf("metadata[%v]", k))
				if err != nil {
					return nil, nil, err
				}
				formField.Write([]byte(v))
			}
		case reflect.Slice:
			switch fieldTag {
			case "signers":
				for i, signer := range embRequest.GetSigners() {
					email, err := w.CreateFormField(fmt.Sprintf("signers[%v][email_address]", i))
					if err != nil {
						return nil, nil, err
					}
					email.Write([]byte(signer.GetEmail()))

					name, err := w.CreateFormField(fmt.Sprintf("signers[%v][name]", i))
					if err != nil {
						return nil, nil, err
					}
					name.Write([]byte(signer.GetName()))

					if signer.Order != 0 {
						order, err := w.CreateFormField(fmt.Sprintf("signers[%v][order]", i))
						if err != nil {
							return nil, nil, err
						}
						order.Write([]byte(strconv.Itoa(signer.GetOrder())))
					}

					if signer.Pin != "" {
						pin, err := w.CreateFormField(fmt.Sprintf("signers[%v][pin]", i))
						if err != nil {
							return nil, nil, err
						}
						pin.Write([]byte(signer.GetPin()))
					}
				}
			case "cc_email_addresses":
				for k, v := range embRequest.GetCCEmailAddresses() {
					formField, err := w.CreateFormField(fmt.Sprintf("cc_email_addresses[%v]", k))
					if err != nil {
						return nil, nil, err
					}
					formField.Write([]byte(v))
				}
			case "form_fields_per_document":
				if len(embRequest.GetFormFieldsPerDocument()) > 0 {
					formField, err := w.CreateFormField(fieldTag)
					if err != nil {
						return nil, nil, err
					}
					ffpdJSON, err := json.Marshal(embRequest.GetFormFieldsPerDocument())
					if err != nil {
						return nil, nil, err
					}
					formField.Write([]byte(ffpdJSON))
				}
			case "file":
				for i, path := range embRequest.GetFile() {
					file, _ := os.Open(path)

					formField, err := w.CreateFormFile(fmt.Sprintf("file[%v]", i), file.Name())
					if err != nil {
						return nil, nil, err
					}
					_, err = io.Copy(formField, file)
				}
			case "file_url":
				for i, fileURL := range embRequest.GetFileURL() {
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

// parseSignatureRequestResponse – Parses the signature request response and converts it into the signature request model
func (m *Client) parseSignatureRequestResponse(response *http.Response) (*model.SignatureRequest, error) {
	defer response.Body.Close()

	sigRequestResponse := &model.SignatureRequestResponse{}

	err := json.NewDecoder(response.Body).Decode(sigRequestResponse)

	sigRequest := sigRequestResponse.GetSignatureRequest()

	return sigRequest, err
}

func (m *Client) boolToIntString(value bool) string {
	if value == true {
		return "1"
	}
	return "0"
}
