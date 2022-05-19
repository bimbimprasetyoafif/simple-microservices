package comment

import (
	"encoding/json"
	"errors"
	pbCom "github.com/bimbimprasetyoafif/api-gateway/pkg/proto/comment"
	"github.com/bimbimprasetyoafif/api-gateway/pkg/provider"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestCreateComment(t *testing.T) {
	testTable := []struct {
		name               string
		f                  func(orgName string, value string) (*pbCom.CreateResp, error)
		expectedResult     map[string]interface{}
		expectedStatusCode int
	}{
		{
			name: "success",
			f: func(orgName string, value string) (*pbCom.CreateResp, error) {
				return &pbCom.CreateResp{
					Status:  200,
					Message: "success",
				}, nil
			},
			expectedResult: map[string]interface{}{
				"status":  float64(200),
				"message": "success",
			},
			expectedStatusCode: 200,
		},
		{
			name: "internal error",
			f: func(orgName string, value string) (*pbCom.CreateResp, error) {
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

			orgClientMock := provider.ClientCommentMock{
				CreateFunc: v.f,
			}
			c := Controller{
				&orgClientMock,
			}
			eCtx := e.NewContext(r, w)
			err := c.CreateComment(eCtx)
			resBody, _ := ioutil.ReadAll(w.Result().Body)
			defer w.Result().Body.Close()

			res := make(map[string]interface{})
			json.Unmarshal(resBody, &res)

			assert.NoError(t, err)
			assert.Equal(t, v.expectedResult, res)
		})
	}
}

func TestGetComment(t *testing.T) {
	testTable := []struct {
		name               string
		f                  func(orgName string) (*pbCom.GetAllResp, error)
		expectedResult     map[string]interface{}
		expectedStatusCode int
	}{
		{
			name: "success",
			f: func(orgName string) (*pbCom.GetAllResp, error) {
				return &pbCom.GetAllResp{
					Status:  200,
					Message: "success",
				}, nil
			},
			expectedResult: map[string]interface{}{
				"status":  float64(200),
				"message": "success",
				"data":    nil,
			},
			expectedStatusCode: 200,
		},
		{
			name: "success with data",
			f: func(orgName string) (*pbCom.GetAllResp, error) {
				return &pbCom.GetAllResp{
					Status:  200,
					Message: "success",
					Data: []string{
						"abc",
						"def",
					},
				}, nil
			},
			expectedResult: map[string]interface{}{
				"status":  float64(200),
				"message": "success",
				"data": []interface{}{
					"abc",
					"def",
				},
			},
			expectedStatusCode: 200,
		},
		{
			name: "internal error",
			f: func(orgName string) (*pbCom.GetAllResp, error) {
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
			r := httptest.NewRequest("GET", "/", nil)
			w := httptest.NewRecorder()
			e := echo.New()

			orgClientMock := provider.ClientCommentMock{
				GetFunc: v.f,
			}
			c := Controller{
				&orgClientMock,
			}
			eCtx := e.NewContext(r, w)
			err := c.GetAllComment(eCtx)
			resBody, _ := ioutil.ReadAll(w.Result().Body)
			defer w.Result().Body.Close()

			res := make(map[string]interface{})
			json.Unmarshal(resBody, &res)

			assert.NoError(t, err)
			assert.Equal(t, v.expectedResult, res)
		})
	}
}

func TestDeleteComment(t *testing.T) {
	testTable := []struct {
		name               string
		f                  func(orgName string) (*pbCom.DeleteResp, error)
		expectedResult     map[string]interface{}
		expectedStatusCode int
	}{
		{
			name: "success",
			f: func(orgName string) (*pbCom.DeleteResp, error) {
				return &pbCom.DeleteResp{
					Status:  200,
					Message: "success",
				}, nil
			},
			expectedResult: map[string]interface{}{
				"status":  float64(200),
				"message": "success",
			},
			expectedStatusCode: 200,
		},
		{
			name: "internal error",
			f: func(orgName string) (*pbCom.DeleteResp, error) {
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
			r := httptest.NewRequest("GET", "/", nil)
			w := httptest.NewRecorder()
			e := echo.New()

			orgClientMock := provider.ClientCommentMock{
				DelFunc: v.f,
			}
			c := Controller{
				&orgClientMock,
			}
			eCtx := e.NewContext(r, w)
			err := c.DeleteComment(eCtx)
			resBody, _ := ioutil.ReadAll(w.Result().Body)
			defer w.Result().Body.Close()

			res := make(map[string]interface{})
			json.Unmarshal(resBody, &res)

			assert.NoError(t, err)
			assert.Equal(t, v.expectedResult, res)
		})
	}
}
