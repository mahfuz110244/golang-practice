package main

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUsersSalesOrder(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/order?user_id=1")
	res := getUsersSalesOrder(c)

	// Assertions
	if assert.NoError(t, res) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestProvideUsersSalesOrderRating(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/order/rating?user_id=1&sales_order_id=1&rating_score=5&rating_comment=Good")
	res := provideUsersSalesOrderRating(c)

	// Assertions
	if assert.NoError(t, res) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
