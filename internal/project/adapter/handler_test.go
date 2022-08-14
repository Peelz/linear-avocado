// An Example test for handler, in this handle I preferred to use Integration test instant unittest.

package adapter

import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/monopeelz/linear-avocado/internal/project/mocks"
	"github.com/monopeelz/linear-avocado/internal/project/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_Create(t *testing.T) {
	t.Run("on error should not returned status 200", func(t *testing.T) {
		uc := mocks.ProjectUseCase{}
		// Stub
		uc.On("Create", mock.Anything, mock.Anything).Return(models.Project{}, errors.New("error"))
		h := NewHandler(&uc)

		w := httptest.NewRecorder()
		gin.SetMode(gin.TestMode)
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest(http.MethodPost, "api/v1/projects", bytes.NewReader([]byte(`{
			"name": "foo",
			"url": "bar"
		}`)))
		h.Create(c)
		assert.NotEqual(t, w.Code, 200)
	})
	t.Run("status should 200", func(t *testing.T) {
		uc := mocks.ProjectUseCase{}
		// Stub
		uc.On("Create", mock.Anything, mock.Anything).Return(models.Project{
			Name: "foo",
			URL:  "bar",
		}, nil)
		h := NewHandler(&uc)

		w := httptest.NewRecorder()
		gin.SetMode(gin.TestMode)
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest(http.MethodPost, "api/v1/projects", bytes.NewReader([]byte(`{
			"name": "foo",
			"url": "bar"
		}`)))
		h.Create(c)
		assert.Equalf(t, 200, w.Code, w.Body.String())
	})

}
