package hellosign

import (
	"bytes"
	"encoding/json"
	"github.com/DeputyApp/hellosign-go-sdk/model"
	"mime/multipart"
	"reflect"
)

// CreateNewApiApp – Creates a new API App.
func (m *Client) CreateNewApiApp(req model.CreateApiAppRequest) (*model.APIApp, error) {
	var params bytes.Buffer
	writer := multipart.NewWriter(&params)

	structType := reflect.TypeOf(req)
	val := reflect.ValueOf(req)

	for i:= 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		f := valueField.Interface()
		val := reflect.ValueOf(f)
		field := structType.Field(i)
		fieldTag := field.Tag.Get("form_field")

		switch val.Kind() {
		default:
			if val.String() != "" {
				formField, err := writer.CreateFormField(fieldTag)
				if err != nil {
					return nil, err
				}
				formField.Write([]byte(val.String()))
			}
		}
	}

	response, err := m.post("api_app", &params, *writer)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	resp := &model.CreateAPIAppResponse{}
	err = json.NewDecoder(response.Body).Decode(resp)
	return resp.GetAPIApp(), err
}
