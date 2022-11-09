package simplehttp

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	type MyHttpResponse struct {
		Foo string `json:"foo"`
	}
	//t.Parallel()
	tcs := []struct {
		name         string
		path         string
		clientParams HttpRequestParams
		serverBody   string
		expectedRes  *MyHttpResponse
	}{
		{
			name:         "ok with custom path",
			path:         "/my-path",
			clientParams: HttpRequestParams{},
			serverBody:   "{\"foo\": \"bar\"}",
			expectedRes:  &MyHttpResponse{Foo: "bar"},
		},
		{
			name:         "ok with query param",
			path:         "/my-query-return-id",
			clientParams: HttpRequestParams{QueryParams: map[string]string{"id": "my-id"}},
			serverBody:   "{\"foo\": \"my-id\"}",
			expectedRes:  &MyHttpResponse{Foo: "my-id"},
		},
	}

	// Table testing
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				switch strings.TrimSpace(r.URL.Path) {
				case "/my-path":
					sc := http.StatusOK
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(sc)
					w.Write([]byte(tc.serverBody))
				case "/my-query-return-id":
					sc := http.StatusOK
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(sc)
					id := r.URL.Query().Get("id")
					log.Debug().Msgf("My id is %v", id)
					w.Write([]byte("{\"foo\": \"" + id + "\"}"))

				default:
					http.NotFoundHandler().ServeHTTP(w, r)
				}
			}))
			res, err := GetJson[MyHttpResponse](http.DefaultClient, server.URL+tc.path, tc.clientParams)
			assert.ErrorIs(t, nil, err)
			assert.Equal(t, tc.expectedRes.Foo, res.Foo)
		})
	}
}
