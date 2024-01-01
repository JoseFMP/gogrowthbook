package test_utils

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/JoseFMP/gogrowthbook"
	"github.com/JoseFMP/gogrowthbook/models"
	"github.com/stretchr/testify/assert"
)

func GrowthBookMockServer(
	t *testing.T,
	handler http.HandlerFunc,
) *httptest.Server {
	server := httptest.NewServer(handler)

	return server
}

func GetGrowthBookMockServerHandler(t *testing.T, secretKey string, projects []models.Project) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pathChunks := strings.Split(r.URL.Path, "/")

		assert.Equal(t, "api", pathChunks[1])
		assert.Equal(t, "v1", pathChunks[2])
		endpointPath := pathChunks[3]

		authHeader := r.Header.Get("Authorization")
		assert.Contains(t, authHeader, fmt.Sprintf("Bearer %s", secretKey))

		switch endpointPath {
		case "projects":
			isRequestAllProjects := len(pathChunks) == 4
			isRequestOneProject := len(pathChunks) == 5
			assert.True(t, isRequestAllProjects || isRequestOneProject)
			assert.Equal(t, r.Method, http.MethodGet)

			if isRequestOneProject {
				targetProjectID := pathChunks[4]
				assert.NotEmpty(t, targetProjectID)

				for _, p := range projects {
					if p.Id == targetProjectID {
						w.Header().Add("Content-Type", "application/json")
						w.WriteHeader(http.StatusOK)
						responsePayload := models.GetProject200Response{
							Project: p,
						}
						errEncoding := json.NewEncoder(w).Encode(responsePayload)
						assert.NoError(t, errEncoding)

						return
					}
				}
			}

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			limit := int32(10)
			count := int32(len(projects))
			offset := int32(0)
			hasMore := false
			projectsResponse := models.ListProjects200Response{
				Projects: projects,
			}

			projectsResponse.SetLimit(limit)
			projectsResponse.SetCount(count)
			projectsResponse.SetOffset(offset)
			projectsResponse.SetTotal(count)
			projectsResponse.SetHasMore(hasMore)
			projectsResponse.SetNextOffset(&offset)

			payload, errEncoding := json.Marshal(projectsResponse)
			assert.NoError(t, errEncoding)

			_, errWritingPayload := w.Write(payload)
			assert.NoError(t, errWritingPayload)
		}
	}
}

// GenerateClientContext generates a new context.Context and sets the GrowthBook API domain and secret key to access GrowthBook API
func GenerateClientContext(secretKey string, serverURL string) context.Context {
	clientContext := context.Background()
	clientContext = context.WithValue(clientContext, gogrowthbook.ContextAccessToken, secretKey)
	clientContext = context.WithValue(clientContext, gogrowthbook.ContextServerIndex, 1)
	clientContext = context.WithValue(clientContext, gogrowthbook.ContextServerVariables, map[string]string{"domain": serverURL})

	return clientContext
}

const growthBookSecretKeyEnvName = "GROWTHBOOK_SECRET_KEY"

func GenerateClientCtxRemoteServer(t *testing.T) context.Context {
	secretKey, isSet := os.LookupEnv(growthBookSecretKeyEnvName)
	if !isSet {
		t.Skipf(
			"To run this test you must specify a GrowthBook secret key with admin rights through the environment variable named %s",
			growthBookSecretKeyEnvName)
	}
	if secretKey == "" {
		t.Skipf(
			"The secret key specified in env %s is empty, therefore skipping test now.",
			growthBookSecretKeyEnvName)
	}

	clientContext := context.Background()
	return context.WithValue(clientContext, gogrowthbook.ContextAccessToken, secretKey)
}
