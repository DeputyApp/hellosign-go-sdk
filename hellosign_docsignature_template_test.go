package hellosign

import (
	"encoding/json"
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

	customFields := []map[string]string{
		{
			"name": "Salary",
			"type": "text",
		},
		{
			"name": "zip",
			"type": "text",
		},
	}
	cf, _ := json.Marshal(customFields)

	res, err := client.GetEmbeddedTemplateEditURL("87553598c48774de21a32ec198624868ecb1667d", string(cf), true, true)

	assert.NotNil(t, res, "Should return response")
	assert.Nil(t, err, "Should not return error")

	assert.Contains(t, res.GetEditURL(), "https://embedded.hellosign.com/prep-and-send/embedded-template?cached_params_token=")
	assert.Equal(t, 1676894925, res.GetExpiresAt())
}

func TestClient_GetEmbeddedTemplateEditURLForPreview(t *testing.T) {
	vcr := fixture("fixtures/docsignature_template/get_embedded_template_edit_url_for_preview")
	defer vcr.Stop() // Make sure recorder is stopped once done with it

	client := createVcrClient(vcr)

	res, err := client.GetEmbeddedTemplateEditURLForPreview("87553598c48774de21a32ec198624868ecb1667d", true)

	assert.NotNil(t, res, "Should return response")
	assert.Nil(t, err, "Should not return error")

	assert.Contains(t, res.GetEditURL(), "https://embedded.hellosign.com/prep-and-send/embedded-template?cached_params_token=")
	assert.Equal(t, 1676894926, res.GetExpiresAt())
}

func TestClient_CreateEmbeddedTemplate(t *testing.T) {
	vcr := fixture("fixtures/docsignature_template/create_embedded_template")
	defer vcr.Stop()

	client := createVcrClient(vcr)
	customFields := []map[string]string{
		{
			"name": "Salary",
			"type": "text",
		},
		{
			"name": "zip",
			"type": "text",
		},
	}
	cf, _ := json.Marshal(customFields)

	req := model.CreateEmbeddedTemplateRequest{
		TestMode: true,
		ClientID: os.Getenv("HELLOSIGN_CLIENT_ID"),
		File:     []string{"fixtures/offer_letter.pdf"},
		Title:    "Offer Letter",
		SignerRoles: []model.SignerRole{
			{
				Name:  "Employee",
				Order: 0,
			},
		},
		Metadata: map[string]string{
			"no":   "cats",
			"more": "dogs",
		},
		ShowPreview:  true,
		CustomFields: string(cf),
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

func TestClient_GetTemplate_Invalid(t *testing.T) {
	vcr := fixture("fixtures/docsignature_template/get_template_invalid")
	defer vcr.Stop()

	templateID := "randomstringhere"

	client := createVcrClient(vcr)
	res, err := client.GetTemplate(templateID)
	require.NotNil(t, err, "Should return an error")
	assert.Nil(t, res, "Should return response")
}

func TestClient_GetTemplate_Valid(t *testing.T) {
	vcr := fixture("fixtures/docsignature_template/get_template_valid")
	defer vcr.Stop()

	// This template is created with two custom fields via the embedded draft template api.
	// However, only one is used in the embedded template editing flow.
	// We want to make sure that there is only one custom field in the document, instead of two.
	templateID := "c5754a7b3dae669d8a645a6cba79b340cdfdf3a2"

	client := createVcrClient(vcr)
	res, err := client.GetTemplate(templateID)
	require.Nil(t, err, "Should return an error")
	assert.NotNil(t, res, "Should return response")
	assert.Equal(t, res.GetTemplateID(), templateID)
	assert.True(t, len(res.GetDocuments()) > 0)
	assert.True(t, len(res.GetDocuments()[0].GetCustomFields()) == 1)
}
