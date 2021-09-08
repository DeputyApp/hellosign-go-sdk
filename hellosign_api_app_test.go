package hellosign

import (
	"github.com/DeputyApp/hellosign-go-sdk/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_CreateNewApiApp(t *testing.T) {
	vcr := fixture("fixtures/api_app/create_api_app")
	defer vcr.Stop()

	client := createVcrClient(vcr)
	expectedName := "Example API App for testing"
	expectedDomain := "example.com"
	expectedCallbackURL := "https://www.example.com/callback"
	expectedWhiteLabelingOptions := "{\"header_background_color\":\"#F7F8F9\",\"primary_button_color\":\"#C0A464\",\"text_color2\":\"#808080\"}"
	req := model.CreateApiAppRequest{
		Name:                 expectedName,
		Domain:               expectedDomain,
		CallbackURL:          expectedCallbackURL,
		WhiteLabelingOptions: expectedWhiteLabelingOptions,
	}
	res, err := client.CreateNewApiApp(req)

	assert.NotNil(t, res, "Should return response")
	assert.Nil(t, err, "Should not return error")

	assert.Equal(t, expectedName, res.GetName())
	assert.Equal(t, expectedDomain, res.GetDomain())
	assert.Equal(t, expectedCallbackURL, res.GetCallbackURL())
	assert.Equal(t, expectedWhiteLabelingOptions, res.GetWhiteLabelingOptions())
	assert.NotEmpty(t, res.GetCreatedAt())
}
