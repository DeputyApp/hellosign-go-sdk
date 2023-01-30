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

const (
	ClientIDKey    string = "client_id"
	TestModeKey    string = "test_mode"
	TitleKey       string = "title"
	SubjectKey     string = "subject"
	MessageKey     string = "message"
	ShowPreviewKey string = "show_preview"
	PreviewOnlyKey string = "preview_only"
	MetadataKey    string = "metadata"
	SignerRolesKey string = "signer_roles"
	FileURLKey     string = "file_url"
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

// GetEmbeddedTemplateEditURL - Retrieves an embedded template object to edit.
func (m *Client) GetEmbeddedTemplateEditURL(templateID string, customFields string, enableEdit bool) (*model.EmbeddedTemplateEditURL, error) {
	if templateID == "" {
		return nil, fmt.Errorf("invalid argument: %s", templateID)
	}

	req := model.EditEmbeddedTemplateRequest{}
	req.ShowPreview = enableEdit
	req.TestMode = true
	req.CustomFields = customFields

	params, writer, err := m.marshalMultipartEditEmbeddedTemplateRequest(req)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("embedded/edit_url/%s", templateID)

	response, err := m.post(path, params, *writer)
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

// GetEmbeddedTemplateEditURLForPreview - Retrieves an embedded template object for preview.
// This method uses model.CreateEmbeddedTemplateRequest under the hood which probably needs to be
// renamed so that it can be reused across different methods.
func (m *Client) GetEmbeddedTemplateEditURLForPreview(templateID string) (*model.EmbeddedTemplateEditURL, error) {
	if templateID == "" {
		return nil, fmt.Errorf("invalid argument: %s", templateID)
	}

	req := model.EditEmbeddedTemplateRequest{}
	req.PreviewOnly = true
	req.TestMode = true

	params, writer, err := m.marshalMultipartEditEmbeddedTemplateRequest(req)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("embedded/edit_url/%s", templateID)

	response, err := m.post(path, params, *writer)
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

func (m *Client) GetTemplate(templateID string) (*model.Template, error) {
	if templateID == "" {
		return nil, fmt.Errorf("invalid argument: %s", templateID)
	}
	path := fmt.Sprintf("template/%s", templateID)

	response, err := m.get(path)
	if err != nil {
		return nil, err
	}

	if response.StatusCode == 404 {
		return nil, fmt.Errorf("template not found")
	} else if response.StatusCode != 200 {
		return nil, fmt.Errorf("error occurred when retrieving template details")
	}

	data := &model.GetTemplateResponse{}
	err = json.NewDecoder(response.Body).Decode(data)
	if err != nil {
		return nil, err
	}
	return data.GetTemplate(), nil
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
		fieldTag := field.Tag.Get(FormFieldKey)

		switch val.Kind() {
		case reflect.Map:
			if fieldTag == MetadataKey {
				for k, v := range embRequest.GetMetadata() {
					formField, err := w.CreateFormField(fmt.Sprintf("%s[%v]", MetadataKey, k))
					if err != nil {
						return nil, nil, err
					}
					formField.Write([]byte(v))
				}
			}
		case reflect.Slice:
			switch fieldTag {
			case TestModeKey:
				tm, err := w.CreateFormField(TestModeKey)
				if err != nil {
					return nil, nil, err
				}
				tm.Write([]byte(m.boolToIntString(embRequest.GetTestMode())))
			case ClientIDKey:
				c, err := w.CreateFormField(ClientIDKey)
				if err != nil {
					return nil, nil, err
				}
				if embRequest.GetClientID() != "" {
					c.Write([]byte(embRequest.GetClientID()))
				}
			case SignerRolesKey:
				for i, sr := range embRequest.GetSignerRoles() {
					name, err := w.CreateFormField(fmt.Sprintf("%s[%v][name]", SignerRolesKey, i))
					if err != nil {
						return nil, nil, err
					}
					name.Write([]byte(sr.GetName()))

					if sr.GetOrder() != 0 {
						order, err := w.CreateFormField(fmt.Sprintf("%s[%v][order]", SignerRolesKey, i))
						if err != nil {
							return nil, nil, err
						}
						order.Write([]byte(strconv.Itoa(sr.GetOrder())))
					}
				}
			case FileKey:
				for i, path := range embRequest.GetFile() {
					file, _ := os.Open(path)

					formField, err := w.CreateFormFile(fmt.Sprintf("%s[%v]", FileKey, i), file.Name())
					if err != nil {
						return nil, nil, err
					}
					_, err = io.Copy(formField, file)
				}
			case FileURLKey:
				for i, fileURL := range embRequest.GetFileURL() {
					formField, err := w.CreateFormField(fmt.Sprintf("%s[%v]", FileURLKey, i))
					if err != nil {
						return nil, nil, err
					}
					formField.Write([]byte(fileURL))
				}
			case TitleKey:
				f, err := w.CreateFormField(TitleKey)
				if err != nil {
					return nil, nil, err
				}
				if embRequest.GetTitle() != "" {
					f.Write([]byte(embRequest.GetTitle()))
				}
			case SubjectKey:
				f, err := w.CreateFormField(SubjectKey)
				if err != nil {
					return nil, nil, err
				}
				if embRequest.GetSubject() != "" {
					f.Write([]byte(embRequest.GetSubject()))
				}
			case MessageKey:
				f, err := w.CreateFormField(MessageKey)
				if err != nil {
					return nil, nil, err
				}
				if embRequest.GetMessage() != "" {
					f.Write([]byte(embRequest.GetMessage()))
				}
			case ShowPreviewKey:
				tm, err := w.CreateFormField(ShowPreviewKey)
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

// TODO Abhishek: merge this with the create and use common parsing code
func (m *Client) marshalMultipartEditEmbeddedTemplateRequest(embRequest model.EditEmbeddedTemplateRequest) (*bytes.Buffer, *multipart.Writer, error) {

	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	structType := reflect.TypeOf(embRequest)
	val := reflect.ValueOf(embRequest)

	for i := 0; i < val.NumField(); i++ {

		valueField := val.Field(i)
		f := valueField.Interface()
		val := reflect.ValueOf(f)
		field := structType.Field(i)
		fieldTag := field.Tag.Get(FormFieldKey)

		switch val.Kind() {
		case reflect.Map:
			if fieldTag == MetadataKey {
				for k, v := range embRequest.GetMetadata() {
					formField, err := w.CreateFormField(fmt.Sprintf("%s[%v]", MetadataKey, k))
					if err != nil {
						return nil, nil, err
					}
					formField.Write([]byte(v))
				}
			}
		case reflect.Slice:
			switch fieldTag {
			case TestModeKey:
				tm, err := w.CreateFormField(TestModeKey)
				if err != nil {
					return nil, nil, err
				}
				tm.Write([]byte(m.boolToIntString(embRequest.GetTestMode())))
			case ClientIDKey:
				c, err := w.CreateFormField(ClientIDKey)
				if err != nil {
					return nil, nil, err
				}
				if embRequest.GetClientID() != "" {
					c.Write([]byte(embRequest.GetClientID()))
				}
			case SignerRolesKey:
				for i, sr := range embRequest.GetSignerRoles() {
					name, err := w.CreateFormField(fmt.Sprintf("%s[%v][name]", SignerRolesKey, i))
					if err != nil {
						return nil, nil, err
					}
					name.Write([]byte(sr.GetName()))

					if sr.GetOrder() != 0 {
						order, err := w.CreateFormField(fmt.Sprintf("%s[%v][order]", SignerRolesKey, i))
						if err != nil {
							return nil, nil, err
						}
						order.Write([]byte(strconv.Itoa(sr.GetOrder())))
					}
				}
			case FileKey:
				for i, path := range embRequest.GetFile() {
					file, _ := os.Open(path)

					formField, err := w.CreateFormFile(fmt.Sprintf("%s[%v]", FileKey, i), file.Name())
					if err != nil {
						return nil, nil, err
					}
					_, err = io.Copy(formField, file)
				}
			case FileURLKey:
				for i, fileURL := range embRequest.GetFileURL() {
					formField, err := w.CreateFormField(fmt.Sprintf("%s[%v]", FileURLKey, i))
					if err != nil {
						return nil, nil, err
					}
					formField.Write([]byte(fileURL))
				}
			case TitleKey:
				f, err := w.CreateFormField(TitleKey)
				if err != nil {
					return nil, nil, err
				}
				if embRequest.GetTitle() != "" {
					f.Write([]byte(embRequest.GetTitle()))
				}
			case SubjectKey:
				f, err := w.CreateFormField(SubjectKey)
				if err != nil {
					return nil, nil, err
				}
				if embRequest.GetSubject() != "" {
					f.Write([]byte(embRequest.GetSubject()))
				}
			case MessageKey:
				f, err := w.CreateFormField(MessageKey)
				if err != nil {
					return nil, nil, err
				}
				if embRequest.GetMessage() != "" {
					f.Write([]byte(embRequest.GetMessage()))
				}
			case ShowPreviewKey:
				tm, err := w.CreateFormField(ShowPreviewKey)
				if err != nil {
					return nil, nil, err
				}
				tm.Write([]byte(m.boolToIntString(embRequest.IsShowingPreview())))
			case PreviewOnlyKey:
				tm, err := w.CreateFormField(PreviewOnlyKey)
				if err != nil {
					return nil, nil, err
				}
				tm.Write([]byte(m.boolToIntString(embRequest.IsPreviewOnly())))
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
