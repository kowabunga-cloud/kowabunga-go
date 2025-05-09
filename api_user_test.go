/*
Kowabunga API documentation

Testing UserAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package kowabunga

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	
)

func Test_kowabunga_UserAPIService(t *testing.T) {

	configuration := NewConfiguration()
	apiClient := NewAPIClient(configuration)

	t.Run("Test UserAPIService CreateUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserAPI.CreateUser(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService DeleteUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId string

		httpRes, err := apiClient.UserAPI.DeleteUser(context.Background(), userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService ListUsers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserAPI.ListUsers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService Login", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserAPI.Login(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService Logout", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.UserAPI.Logout(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService ReadUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId string

		resp, httpRes, err := apiClient.UserAPI.ReadUser(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService ResetPassword", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.UserAPI.ResetPassword(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService ResetUserPassword", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId string

		httpRes, err := apiClient.UserAPI.ResetUserPassword(context.Background(), userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService SetUserApiToken", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId string

		httpRes, err := apiClient.UserAPI.SetUserApiToken(context.Background(), userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService SetUserPassword", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId string

		httpRes, err := apiClient.UserAPI.SetUserPassword(context.Background(), userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService UpdateUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId string

		resp, httpRes, err := apiClient.UserAPI.UpdateUser(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
