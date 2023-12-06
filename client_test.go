package gogrowthbook_test

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/JoseFMP/gogrowthbook"
	"github.com/JoseFMP/gogrowthbook/models"
	"github.com/stretchr/testify/assert"
)

const growthBookSecretKeyEnvName = "GROWTHBOOK_SECRET_KEY"

// TestCanCreateAndReadFeature
// creates a feature in GrowthBook, and reads it out afterward to double check it was really created
func TestCanCreateAndReadFeature(t *testing.T) {

	secretKey := os.Getenv(growthBookSecretKeyEnvName)
	if secretKey == "" {
		t.Logf(
			"To run this test you must specify a GrowthBook secret key with admin rights through the environment variable named %s",
			growthBookSecretKeyEnvName)
		t.FailNow()
	}

	clientContext := context.Background()
	clientContext = context.WithValue(clientContext, gogrowthbook.ContextAccessToken, secretKey)

	randomNumber := time.Now().UnixMilli()
	featureID := fmt.Sprintf("go-growthbook-integration-test-%d", randomNumber)
	featureValueType := "string"
	featureDescription :=
		fmt.Sprintf(
			"This is a feature created during integration tests of go-growthbook lib, remove safely - %d",
			randomNumber)

	cfg := gogrowthbook.NewConfiguration()
	client := gogrowthbook.NewAPIClient(cfg)

	t.Run("create feature",
		func(t *testing.T) {

			newFeature := models.PostFeatureRequest{
				Id:          featureID,
				ValueType:   featureValueType,
				Description: &featureDescription,
			}
			newFeatureReq := client.FeaturesAPI.PostFeature(clientContext).PostFeatureRequest(newFeature)

			responsePayload, httpResponse, errDoingReq := newFeatureReq.Execute()
			assert.NoErrorf(t, errDoingReq, "The request should return success")
			assert.NotNil(t, httpResponse)
			assert.NotNil(t, responsePayload)

			assert.Equal(t, http.StatusOK, httpResponse.StatusCode)

		})

	t.Run("readout feature", func(t *testing.T) {
		featureResponse, httpResponse, errReadingOutFeature :=
			client.FeaturesAPI.GetFeature(clientContext, featureID).Execute()

		assert.NoErrorf(t, errReadingOutFeature, "There should be no error fetching the feature just created")
		assert.NotNil(t, httpResponse)
		assert.Equal(t, http.StatusOK, httpResponse.StatusCode)

		assert.NotNil(t, featureResponse)
		assert.NotNil(t, featureResponse.Feature)
		assert.Equal(t, featureResponse.Feature.Id, featureID)
		assert.Equal(t, featureResponse.Feature.ValueType, featureValueType)
		assert.Equal(t, featureResponse.Feature.Description, featureDescription)

	},
	)
}
