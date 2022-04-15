package http_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	handler "github.com/Mustafa0831/TSarka/internal/server/delivery/http"

	"github.com/Mustafa0831/TSarka/internal/counter"
	"github.com/Mustafa0831/TSarka/internal/email"
	"github.com/stretchr/testify/require"
)

func TestNewHandler(t *testing.T) {
	h := handler.NewHandler(&counter.Service{}, &email.Service{})
	require.IsType(t, &handler.Handler{}, h)
}

func TestNewHandler_Init(t *testing.T) {
	h := handler.NewHandler(&counter.Service{}, &email.Service{})

	router := h.Init()

	ts := httptest.NewServer(router)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/ping")
	if err != nil {
		t.Error(err)
	}
	require.Equal(t, http.StatusOK, res.StatusCode)
}
