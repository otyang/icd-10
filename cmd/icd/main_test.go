package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/otyang/icd-10/internal/icd"
	"github.com/otyang/icd-10/pkg/logger"
	"github.com/otyang/icd-10/pkg/middleware"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	reqMethod            string // method to use in sending requuest path
	reqRoute             string // route path to test
	reqBody              *string
	expectedResponseCode int    // expected HTTP status code
	description          string // description of the test case
}{
	{
		reqMethod:            "GET",
		reqRoute:             "/icd/100",
		reqBody:              nil,
		description:          "Get ICD endpoint",
		expectedResponseCode: http.StatusOK,
	},
}

func Test_main(t *testing.T) {
	logger.WithBaseInfo(zlog, cfg.AppName, cfg.AppAddress)

	defer cancel()
	defer func() {
		if err := db.Close(); err != nil {
			zlog.Fatal("Error closing database: " + err.Error())
		}
	}()

	{
		router.Use(
			middleware.Cors(),
			middleware.HttpLog(zlog),
		)
		icd.RegisterHttpHandlers(ctx, router, cfg, zlog, db)
		icd.RegisterEventsHandlers(ctx, cfg, zlog)
	}

	// Actual testcase
	for _, test := range tests {

		newBody, _ := json.Marshal(test.reqBody)
		newbody2bytes := bytes.NewReader(newBody)

		req := httptest.NewRequest(test.reqMethod, test.reqRoute, newbody2bytes)
		resp, err := router.Test(req, -1)

		assert.Nil(t, err, test.description)
		assert.Equalf(t, test.expectedResponseCode, resp.StatusCode, test.description)
	}
}
