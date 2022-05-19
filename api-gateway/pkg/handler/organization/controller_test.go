package organization

import (
	"encoding/json"
	"errors"
	pbOrg "github.com/bimbimprasetyoafif/api-gateway/pkg/proto/organization"
	"github.com/bimbimprasetyoafif/api-gateway/pkg/provider"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestRegisterOrganization(t *testing.T) {
	testTable := []struct {
		name               string
		f                  func(orgName string) (*pbOrg.MessageResp, error)
		expectedResult     map[string]interface{}
		expectedStatusCode int
	}{
		{
			name: "success",
			f: func(orgName string) (*pbOrg.MessageResp, error) {
				return &pbOrg.MessageResp{
					Status:  200,
					Message: "success",
					Data: &pbOrg.Organization{
						Name: "abc",
						Slug: "abc",
					},
				}, nil
			},
			expectedResult: map[string]interface{}{
				"status":  float64(200),
				"message": "success",
				"data": map[string]interface{}{
					"name": "abc",
					"slug": "abc",
				},
			},
			expectedStatusCode: 200,
		},
		{
			name: "internal error",
			f: func(orgName string) (*pbOrg.MessageResp, error) {
				return nil, errors.New("error")
			},
			expectedResult: map[string]interface{}{
				"status":  float64(500),
				"message": "internal error",
			},
			expectedStatusCode: 500,
		},
	}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			r := httptest.NewRequest("POST", "/", nil)
			w := httptest.NewRecorder()
			e := echo.New()

			orgClientMock := provider.ClientOrgMock{
				RegisFunc: v.f,
			}
			c := Controller{
				&orgClientMock,
			}
			eCtx := e.NewContext(r, w)
			err := c.RegisterOrganization(eCtx)
			resBody, _ := ioutil.ReadAll(w.Result().Body)
			defer w.Result().Body.Close()

			res := make(map[string]interface{})
			json.Unmarshal(resBody, &res)

			assert.NoError(t, err)
			assert.Equal(t, v.expectedResult, res)
		})
	}
}
