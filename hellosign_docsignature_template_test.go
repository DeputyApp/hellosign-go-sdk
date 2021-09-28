package hellosign

import (
	"github.com/DeputyApp/hellosign-go-sdk/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestClient_GetEmbeddedTemplateEditURL(t *testing.T) {
	vcr := fixture("fixtures/docsignature_template/get_embedded_template_edit_url")
	defer vcr.Stop() // Make sure recorder is stopped once done with it

	client := createVcrClient(vcr)

	res, err := client.GetEmbeddedTemplateEditURL("76a888f4ca1dc1f726cbfd3381d7b9a19066c047")

	assert.NotNil(t, res, "Should return response")
	assert.Nil(t, err, "Should not return error")

	assert.Contains(t, res.GetEditURL(), "https://embedded.hellosign.com/prep-and-send/embedded-template?cached_params_token=")
	assert.Equal(t, 1630908730, res.GetExpiresAt())
}

func TestClient_CreateEmbeddedTemplate(t *testing.T) {
	vcr := fixture("fixtures/docsignature_template/create_embedded_template")
	defer vcr.Stop()

	client := createVcrClient(vcr)
	customFields := []model.CustomField{
		{
			Name:     "Salary",
			Type:     "text",
		},
		{
			Name:     "zip",
			Type:     "text",
		},
	}
	
	req := model.CreateEmbeddedTemplateRequest{
		TestMode: true,
		ClientID: os.Getenv("HELLOSIGN_CLIENT_ID"),
		File:     []string{"fixtures/offer_letter.pdf"},
		Title:    "Offer Letter",
		SignerRoles: []model.SignerRole{model.SignerRole{
			Name:  "Employee",
			Order: 0,
		}},
		Metadata: map[string]string{
			"no":   "cats",
			"more": "dogs",
		},
		ShowPreview: true,
		CustomFields: customFields,
	}

	res, err := client.CreateEmbeddedTemplate(req)
	require.NotNil(t, res, "Should return response")
	require.Nil(t, err, "Should not return error")

	assert.NotEmpty(t, res.GetTemplateID())
	assert.NotEmpty(t, res.GetEditURL())
	assert.NotEmpty(t, res.GetExpiresAt())
}

func TestClient_ListTemplates(t *testing.T) {
	vcr := fixture("fixtures/docsignature_template/list_templates")
	defer vcr.Stop()

	client := createVcrClient(vcr)
	res, err := client.ListTemplates()
	assert.NotNil(t, res, "Should return response")
	assert.Nil(t, err, "Should not return error")
	assert.Greater(t, len(res.GetTemplates()), 0)

	assert.NotNil(t, res.GetListInfo())
	assert.Equal(t, res.GetListInfo().GetNumResults(), len(res.GetTemplates()))
}

func TestClient_DeleteTemplate(t *testing.T) {
	vcr := fixture("fixtures/docsignature_template/delete_template")
	defer vcr.Stop()

	client := createVcrClient(vcr)
	res, err := client.DeleteTemplate("9fe8fb79ecec08f61c9120912eba738c9011ebe1")
	assert.NotNil(t, res, "Should return response")
	assert.Nil(t, err, "Should not return error")
}