package routes_test

import (
	"Asiayo/constant"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"Asiayo/routes"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine

type test struct {
	name string
	code int
	path string
	args interface{}
	want constant.Response
}

type currencyResponseParam struct {
	Amount string `json:"amount"`
}

func TestCurrencyExchange(t *testing.T) {
	router = routes.InitRouter()
	path := "/api/v1/currency/exchange"
	tests := []test{
		{
			name: "Valid/SourceAndTarget",
			path: path,
			args: url.Values{
				"source": []string{"USD"},
				"target": []string{"JPY"},
				"amount": []string{"1525.0"},
			},
			code: http.StatusOK,
			want: constant.Response{
				Code: constant.SUCCESS,
				Data: currencyResponseParam{
					Amount: "170,496.53",
				},
			},
		},
		{
			name: "Invalid/Source",
			path: path,
			args: url.Values{
				"source": []string{"USDD"},
				"target": []string{"JPY"},
				"amount": []string{"1525.0"},
			},
			code: http.StatusOK,
			want: constant.Response{
				Code: constant.INVALID_PARAMS,
				Data: "Key: 'Source' Error:Field validation for 'Source' failed on the 'oneof' tag",
			},
		},
		{
			name: "Invalid/Target",
			path: path,
			args: url.Values{
				"source": []string{"USD"},
				"target": []string{"JPYY"},
				"amount": []string{"1525.0"},
			},
			code: http.StatusOK,
			want: constant.Response{
				Code: constant.INVALID_PARAMS,
				Data: "Key: 'Target' Error:Field validation for 'Target' failed on the 'oneof' tag",
			},
		},
		{
			name: "Invalid/amount",
			path: path,
			args: url.Values{
				"source": []string{"USD"},
				"target": []string{"JPY"},
				"amount": []string{"-1000"},
			},
			code: http.StatusOK,
			want: constant.Response{
				Code: constant.INVALID_PARAMS,
				Data: "Key: 'Amount' Error:Field validation for 'Amount' failed on the 'gt' tag",
			},
		},
		{
			name: "Invalid/amount/decimal",
			path: path,
			args: url.Values{
				"source": []string{"USD"},
				"target": []string{"JPY"},
				"amount": []string{"1000.8888"},
			},
			code: http.StatusOK,
			want: constant.Response{
				Code: constant.DECIMAL_INVALID,
				Data: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := url.URL{
				Path:     tt.path,
				RawQuery: tt.args.(url.Values).Encode(),
			}
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, u.String(), nil)
			router.ServeHTTP(w, r)
			assert.Equal(t, tt.code, w.Code)

			var resp constant.Response
			err := json.Unmarshal(w.Body.Bytes(), &resp)
			assert.NoError(t, err)

			assert.Equal(t, tt.want.Code, resp.Code)

			if tt.want.Data != nil {
				if tt.want.Code == constant.SUCCESS && resp.Code == constant.SUCCESS {
					var data currencyResponseParam
					jsonStr, err := json.Marshal(resp.Data)
					assert.NoError(t, err)
					err = json.Unmarshal(jsonStr, &data)
					assert.NoError(t, err)

					want := tt.want.Data.(currencyResponseParam)
					assert.Equal(t, want.Amount, data.Amount)

				} else {
					assert.Equal(t, tt.want.Data, resp.Data)
				}
			}
		})
	}
}
