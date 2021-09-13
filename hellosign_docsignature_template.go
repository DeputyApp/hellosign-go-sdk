package hellosign

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/DeputyApp/hellosign-go-sdk/model"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"reflect"
	"strconv"
)

// CreateEmbeddedTemplate creates a new embedded Template
func (m *Client) CreateEmbeddedTemplate(req model.CreateEmbeddedTemplateRequest) (*model.EmbeddedTemplate, error) {
	params, writer, err := m.marshalMultipartCreateEmbeddedTemplateRequest(req)
	if err != nil {
		return nil, err
	}

	response, err := m.post("template/create_embedded_draft", params, *writer)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	resp := &model.CreateEmbeddedTemplateResponse{}
	err = json.NewDecoder(response.Body).Decode(resp)
	return resp.GetTemplate(), err
}

// ListTemplates retrieves a list that are accessible by your account
func (m *Client) ListTemplates() (*model.ListTemplatesResponse, error) {
	path := fmt.Sprintf("template/list")
	response, err := m.get(path)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	listResponse := &model.ListTemplatesResponse{}
	err = json.NewDecoder(response.Body).Decode(listResponse)
	if err != nil {
		return nil, err
	}

	return listResponse, err
}

// DeleteTemplate completely deletes the template specified from the account and is irreversible
func (m *Client) DeleteTemplate(templateID string) (*http.Response, error) {
	path := fmt.Sprintf("template/delete/%s", templateID)

	response, err := m.nakedPost(path)
	if err != nil {
		return nil, err
	}

	return response, err
}

// GetEmbeddedTemplateEditURL - Retrieves an embedded template object.
func (m *Client) GetEmbeddedTemplateEditURL(templateID string) (*model.EmbeddedTemplateEditURL, error) {
	if templateID == "" {
		return nil, fmt.Errorf("invalid argument: %s", templateID)
	}
	path := fmt.Sprintf("embedded/edit_url/%s", templateID)

	response, err := m.get(path)
	if err != nil {
		return nil, err
	}

	data := &model.EmbeddedTemplateResponse{}
	err = json.NewDecoder(response.Body).Decode(data)
	if err != nil {
		return nil, err
	}
	return data.GetEmbedded(), nil
}

func (m *Client) marshalMultipartCreateEmbeddedTemplateRequest(embRequest model.CreateEmbeddedTemplateRequest) (*bytes.Buffer, *multipart.Writer, error) {

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
			case "test_mode":
				tm, err := w.CreateFormField("test_mode")
				if err != nil {
					return nil, nil, err
				}
				tm.Write([]byte(m.boolToIntString(embRequest.GetTestMode())))
			case "client_id":
				c, err := w.CreateFormField("client_id")
				if err != nil {
					return nil, nil, err
				}
				if embRequest.GetClientID() != "" {
					c.Write([]byte(embRequest.GetClientID()))
				}
			case "signer_roles":
				for i, sr := range embRequest.GetSignerRoles() {
					name, err := w.CreateFormField(fmt.Sprintf("signer_roles[%v][name]", i))
					if err != nil {
						return nil, nil, err
					}
					name.Write([]byte(sr.GetName()))

					if sr.GetOrder() != 0 {
						order, err := w.CreateFormField(fmt.Sprintf("signer_roles[%v][order]", i))
						if err != nil {
							return nil, nil, err
						}
						order.Write([]byte(strconv.Itoa(sr.GetOrder())))
					}
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
			case "title":
				f, err := w.CreateFormField("title")
				if err != nil {
					return nil, nil, err
				}
				if embRequest.GetTitle() != "" {
					f.Write([]byte(embRequest.GetTitle()))
				}
			case "subject":
				f, err := w.CreateFormField("subject")
				if err != nil {
					return nil, nil, err
				}
				if embRequest.GetSubject() != "" {
					f.Write([]byte(embRequest.GetSubject()))
				}
			case "message":
				f, err := w.CreateFormField("message")
				if err != nil {
					return nil, nil, err
				}
				if embRequest.GetMessage() != "" {
					f.Write([]byte(embRequest.GetMessage()))
				}
			case "show_preview":
				tm, err := w.CreateFormField("show_preview")
				if err != nil {
					return nil, nil, err
				}
				tm.Write([]byte(m.boolToIntString(embRequest.IsShowingPreview())))
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
