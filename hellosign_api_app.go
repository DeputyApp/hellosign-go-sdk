package hellosign

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/DeputyApp/hellosign-go-sdk/model"
	"io"
	"mime/multipart"
	"os"
	"reflect"
)

const (
	HellosignCustomLogoFileKey = "custom_logo_file"
	DomainsKey    string = "domains"
)

// CreateNewApiApp – Creates a new API App.
// Note: we don't support a single domain at the moment as it is  out of our current use cases
func (m *Client) CreateNewApiApp(req model.CreateApiAppRequest) (*model.APIApp, error) {
	var params bytes.Buffer
	writer := multipart.NewWriter(&params)

	structType := reflect.TypeOf(req)
	val := reflect.ValueOf(req)

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		f := valueField.Interface()
		val := reflect.ValueOf(f)
		field := structType.Field(i)
		fieldTag := field.Tag.Get(FormFieldKey)

		switch val.Kind() {
		case reflect.Slice:
			switch fieldTag {
			case DomainsKey:
				for i, domain := range req.GetDomains() {
					name, err := writer.CreateFormField(fmt.Sprintf("%s[%v]", DomainsKey, i))
					if err != nil {
						return nil, err
					}
					name.Write([]byte(domain))
				}
			}
		default:
			if val.String() != "" {
				if fieldTag == HellosignCustomLogoFileKey {
					path := val.String()
					file, err := os.Open(path)
					if err != nil {
						return nil, err
					}
					formField, err := writer.CreateFormFile(fieldTag, file.Name())
					if err != nil {
						return nil, err
					}
					_, err = io.Copy(formField, file)
				} else {
					formField, err := writer.CreateFormField(fieldTag)
					if err != nil {
						return nil, err
					}
					formField.Write([]byte(val.String()))
				}
			}
		}
	}
	writer.Close()

	response, err := m.post("api_app", &params, *writer)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	resp := &model.CreateAPIAppResponse{}
	err = json.NewDecoder(response.Body).Decode(resp)
	return resp.GetAPIApp(), err
}