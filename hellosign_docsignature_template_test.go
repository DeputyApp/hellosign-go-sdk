package hellosign

import (
	"github.com/stretchr/testify/assert"
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
