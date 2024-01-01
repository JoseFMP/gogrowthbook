package gogrowthbook_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/JoseFMP/gogrowthbook"
	"github.com/JoseFMP/gogrowthbook/models"
	"github.com/JoseFMP/gogrowthbook/test_utils"
	"github.com/stretchr/testify/assert"
)

// TestCanCreateAndReadFeature
// creates a feature in GrowthBook, and reads it out afterward to double check it was really created
func TestCanCreateAndReadFeature(t *testing.T) {
	clientContext := test_utils.GenerateClientCtxRemoteServer(t)

	randomNumber := time.Now().UnixMilli()
	featureID := fmt.Sprintf("go-growthbook-integration-test-%d", randomNumber)
	featureValueType := "string"
	featureDescription := fmt.Sprintf(
		"This is a feature created during integration tests of go-growthbook lib, remove safely - %d",
		randomNumber)
	forceRuleDescription := "custom-force-rule-description"
	forcedValue := "abc"
	forcedEnable := false
	forcedType := models.FeatureRuleTypeForce

	cfg := gogrowthbook.NewConfiguration()
	client := gogrowthbook.NewAPIClient(cfg)

	growthBookEnvironmentName := "production"
	t.Run("create feature",
		func(t *testing.T) {
			newFeature := models.PostFeatureRequest{
				Id:          featureID,
				ValueType:   featureValueType,
				Description: &featureDescription,
				Environments: map[string]models.PostFeatureRequestEnvironmentsValue{
					growthBookEnvironmentName: {
						Enabled: false,
						Rules: []*models.FeatureRule{
							{
								FeatureForceRule: &models.FeatureForceRule{
									Type:        &forcedType,
									Enabled:     &forcedEnable,
									Value:       &forcedValue,
									Description: &forceRuleDescription,
								},
							},
						},
					},
				},
			}

			createFeature(clientContext, t, newFeature)
		})

	t.Run("readout feature", func(t *testing.T) {
		featureResponse, httpResponse, errReadingOutFeature := client.FeaturesAPI.GetFeature(clientContext, featureID).Execute()

		assert.NoErrorf(t, errReadingOutFeature, "There should be no error fetching the feature just created")
		assert.NotNil(t, httpResponse)
		assert.Equal(t, http.StatusOK, httpResponse.StatusCode)

		assert.NotNil(t, featureResponse)
		assert.NotNil(t, featureResponse.Feature)
		assert.Equal(t, featureResponse.Feature.Id, featureID)
		assert.Equal(t, featureResponse.Feature.ValueType, featureValueType)
		assert.Equal(t, featureResponse.Feature.Description, featureDescription)

		_, hasTargetEnvironment := featureResponse.Feature.Environments[growthBookEnvironmentName]
		assert.True(t, hasTargetEnvironment)
		assert.Len(t, featureResponse.Feature.Environments[growthBookEnvironmentName].Rules, 1)

		assert.NotNil(t, featureResponse.Feature.Environments[growthBookEnvironmentName].Rules[0].FeatureForceRule, 1)
		assert.Nil(t, featureResponse.Feature.Environments[growthBookEnvironmentName].Rules[0].FeatureRolloutRule, 1)
		assert.Nil(t, featureResponse.Feature.Environments[growthBookEnvironmentName].Rules[0].FeatureExperimentRule, 1)
		assert.Nil(t, featureResponse.Feature.Environments[growthBookEnvironmentName].Rules[0].FeatureExperimentRefRule, 1)

		assert.Equal(t, *featureResponse.Feature.Environments[growthBookEnvironmentName].Rules[0].FeatureForceRule.Type, models.FeatureRuleType("force"))
		assert.Equal(t, *featureResponse.Feature.Environments[growthBookEnvironmentName].Rules[0].FeatureForceRule.Value, forcedValue)
	},
	)
}

func TestCanReadAllFeatures(t *testing.T) {
	clientContext := test_utils.GenerateClientCtxRemoteServer(t)

	t.Run("Read with pagination", func(t *testing.T) {
		// arrange
		randomFeatures := make([]models.PostFeatureRequest, 10)
		for i := range randomFeatures {
			randomFeatures[i] = *generateRandomFeature()
			createFeature(clientContext, t, randomFeatures[i])
			time.Sleep(time.Second)
		}
		clt := gogrowthbook.NewAPIClient(gogrowthbook.NewConfiguration())

		// act
		resp, httpResp, errDoingReq := clt.FeaturesAPI.ListFeatures(clientContext).Limit(1).Execute()

		// assert
		assert.NoError(t, errDoingReq)
		assert.NotNil(t, resp)
		assert.NotNil(t, httpResp)
		assert.Equal(t, http.StatusOK, httpResp.StatusCode)

		assert.NotNil(t, resp.HasMore)
		assert.True(t, *resp.HasMore)

		assert.NotNil(t, resp.Count)
		assert.Equal(t, 1, *resp.Count)

		assert.NotNil(t, resp.Total)
		assert.GreaterOrEqual(t, *resp.Total, len(randomFeatures))
	})
}

func createFeature(ctx context.Context, t *testing.T, feature models.PostFeatureRequest) {
	cfg := gogrowthbook.NewConfiguration()
	client := gogrowthbook.NewAPIClient(cfg)
	newFeatureReq := client.FeaturesAPI.PostFeature(ctx).PostFeatureRequest(feature)

	responsePayload, httpResponse, errDoingReq := newFeatureReq.Execute()
	assert.NoErrorf(t, errDoingReq, "The request should return success", errDoingReq)
	assert.NotNil(t, httpResponse)
	assert.NotNil(t, responsePayload)

	assert.Equalf(t, http.StatusOK, httpResponse.StatusCode, httpResponse.Status)
}

func generateRandomFeature() *models.PostFeatureRequest {
	randomizer := fmt.Sprintf("%d", time.Now().Unix())
	id := fmt.Sprintf("gogrowthbook-tests-%s", randomizer)
	description := "Feature created by tests of package github.com/JoseFMP/gogrowthbook"
	return &models.PostFeatureRequest{
		Id:           id,
		ValueType:    "string",
		Description:  &description,
		Environments: map[string]models.PostFeatureRequestEnvironmentsValue{},
	}
}
