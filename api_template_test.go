/*
Kowabunga API documentation

Testing TemplateAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package kowabunga

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	
)

func Test_kowabunga_TemplateAPIService(t *testing.T) {

	configuration := NewConfiguration()
	apiClient := NewAPIClient(configuration)

	t.Run("Test TemplateAPIService DeleteTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var templateId string

		httpRes, err := apiClient.TemplateAPI.DeleteTemplate(context.Background(), templateId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplateAPIService ListTemplates", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TemplateAPI.ListTemplates(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplateAPIService ReadTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var templateId string

		resp, httpRes, err := apiClient.TemplateAPI.ReadTemplate(context.Background(), templateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplateAPIService UpdateTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var templateId string

		resp, httpRes, err := apiClient.TemplateAPI.UpdateTemplate(context.Background(), templateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
