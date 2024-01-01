package gogrowthbook_test

import (
	"net/http"
	"testing"

	"github.com/JoseFMP/gogrowthbook"
	"github.com/JoseFMP/gogrowthbook/models"
	"github.com/JoseFMP/gogrowthbook/test_utils"
	"github.com/stretchr/testify/assert"
)

func TestCanReadProjects(t *testing.T) {
	mockProject1 := models.Project{
		Id:   "project1",
		Name: "Mock project 1",
	}

	mockProject2 := models.Project{
		Id:   "project2",
		Name: "Mock project 2",
	}

	mockProject3 := models.Project{
		Id:   "project2",
		Name: "Mock project 2",
	}

	mockProjects := []models.Project{
		mockProject1,
		mockProject2,
		mockProject3,
	}

	mockSercretKey := "09lknaslfkjalsdfaksdfj"
	mockServer := test_utils.GrowthBookMockServer(t, test_utils.GetGrowthBookMockServerHandler(t, mockSercretKey, mockProjects))

	t.Run("can read all projects in org", func(t *testing.T) {
		// arrange
		clientContext := test_utils.GenerateClientContext(mockSercretKey, mockServer.URL)
		gbClient := gogrowthbook.NewAPIClient(gogrowthbook.NewConfiguration())

		// act
		projectsResp, httpResp, errDoingReq := gbClient.ProjectsAPI.ListProjects(clientContext).Execute()

		// verify
		assert.NoError(t, errDoingReq)
		assert.NotNil(t, projectsResp)
		assert.NotNil(t, httpResp)

		assert.Equal(t, httpResp.StatusCode, http.StatusOK)
		assert.Len(t, projectsResp.Projects, 3)
	})

	t.Run("can read specific project", func(t *testing.T) {
		// arrange
		clientContext := test_utils.GenerateClientContext(mockSercretKey, mockServer.URL)
		gbClient := gogrowthbook.NewAPIClient(gogrowthbook.NewConfiguration())

		// act
		projectResp, httpResp, errDoingReq := gbClient.ProjectsAPI.GetProject(clientContext, mockProject1.Id).Execute()

		// verify
		assert.NoError(t, errDoingReq)
		assert.NotNil(t, projectResp)
		assert.NotNil(t, httpResp)

		assert.Equal(t, httpResp.StatusCode, http.StatusOK)
		assert.NotNil(t, projectResp)
		assert.NotNil(t, projectResp.Project)
		assert.Equal(t, projectResp.Project.Id, mockProject1.Id)
	})
}

func TestCanReadProjectsFromRemote(t *testing.T) {
	clientCtx := test_utils.GenerateClientCtxRemoteServer(t)
	projectsAPI := gogrowthbook.NewAPIClient(gogrowthbook.NewConfiguration()).ProjectsAPI
	targetProjectID := ""
	t.Run("read all projects from remote", func(t *testing.T) {
		// act
		projectsResp, httpResp, errDoingReq := projectsAPI.ListProjects(clientCtx).Execute()

		// verify
		assert.NoError(t, errDoingReq)
		assert.NotNil(t, projectsResp)
		assert.NotNil(t, httpResp)

		assert.Equal(t, http.StatusOK, httpResp.StatusCode)

		if len(projectsResp.Projects) < 1 {
			t.Skip("This test must run on a GrowthBook account that contains at least one project projects")
		}

		targetProjectID = projectsResp.Projects[0].Id
	})

	t.Run("Fetch one particular project", func(t *testing.T) {
		// arrange
		if targetProjectID == "" {
			t.Skip("Did not specify project id")
		}

		// act
		projectResp, httpResp, errDoingReq := projectsAPI.GetProject(clientCtx, targetProjectID).Execute()

		// verify
		assert.NoError(t, errDoingReq)
		assert.NotNil(t, httpResp)
		assert.NotNil(t, projectResp)

		assert.NotNil(t, projectResp.Project)
		assert.Equal(t, targetProjectID, projectResp.Project.Id)
	})
}
