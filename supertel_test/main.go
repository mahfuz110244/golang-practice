package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4/middleware"
)

type (
	SalesOrder struct {
		ID            int        `json:"id"`
		UserId        int        `json:"user_id"`
		StoreId       int        `json:"store_id"`
		CreatedAt     *time.Time `json:"created_at"`
		UpdatedAt     *time.Time `json:"updated_at"`
		RatingScore   int        `json:"rating_score"`
		RatingComment string     `json:"rating_comment"`
		Status        string     `json:"status"`
	}
)

var createAt = time.Now()
var salesOrder = []SalesOrder{
	{ID: 1, UserId: 1, StoreId: 1, CreatedAt: &createAt, UpdatedAt: nil, RatingScore: 5, RatingComment: "Good", Status: "completed"},
	{ID: 2, UserId: 1, StoreId: 2, CreatedAt: &createAt, UpdatedAt: nil, RatingScore: 4, RatingComment: "Good", Status: "pending"},
	{ID: 3, UserId: 1, StoreId: 2, CreatedAt: &createAt, UpdatedAt: nil, RatingScore: 4, RatingComment: "Good", Status: "completed"},
	{ID: 4, UserId: 3, StoreId: 2, CreatedAt: &createAt, UpdatedAt: nil, RatingScore: 4, RatingComment: "Good", Status: "completed"},
}

func getUsersSalesOrder(c echo.Context) error {
	userId := c.QueryParam("user_id")
	userIdInt, _ := strconv.Atoi(userId)
	userSalesOrder := []SalesOrder{}
	for _, so := range salesOrder {
		if so.UserId == userIdInt {
			userSalesOrder = append(userSalesOrder, so)
		}
	}
	return c.JSON(http.StatusOK, userSalesOrder)
}

func provideUsersSalesOrderRating(c echo.Context) error {
	userId := c.QueryParam("user_id")
	salesOrderId := c.QueryParam("sales_order_id")
	ratingScore := c.QueryParam("rating_score")
	ratingComment := c.QueryParam("rating_comment")
	userIdInt, _ := strconv.Atoi(userId)
	salesOrderIdInt, _ := strconv.Atoi(salesOrderId)
	ratingScoreInt, _ := strconv.Atoi(ratingScore)
	for index, so := range salesOrder {
		if so.Status == "completed" && so.UserId == userIdInt && so.ID == salesOrderIdInt {
			dateTimeTwoDaysAgo := time.Now().AddDate(0, 0, -2)
			if so.CreatedAt != nil && so.CreatedAt.Before(dateTimeTwoDaysAgo) {
				return c.JSON(http.StatusBadRequest, "Rating can only be given within 2 days after the order is completed")
			}
			salesOrder[index] = SalesOrder{
				ID:            so.ID,
				UserId:        so.UserId,
				StoreId:       so.StoreId,
				CreatedAt:     so.CreatedAt,
				UpdatedAt:     so.UpdatedAt,
				RatingScore:   ratingScoreInt,
				RatingComment: ratingComment,
				Status:        so.Status,
			}
			fmt.Println(so, userId, salesOrderId, ratingScore, ratingComment)
		}
	}
	return c.JSON(http.StatusOK, nil)
}

func getSpecialWords(c echo.Context) error {
	text := c.QueryParam("text")
	fmt.Println(text)
	textArr := strings.Split(text, " ")
	textArraLower := []string{}
	for _, word := range textArr {
		textArraLower = append(textArraLower, strings.ToLower(word))
	}
	//make a map
	// take two word as array
	specialWordsArray := [][]string{}
	for _, word1 := range textArraLower {
		for _, word2 := range textArraLower {
			if word1 != word2 && sameCharter(word1, word2) {
				fmt.Println(word1, word2)
				specialWords := []string{}
				specialWords = append(specialWords, word1, word2)
				for _, specialWordsArr := range specialWordsArray {
					if specialWordsArr[0] == word2 && specialWordsArr[1] == word1 {
						specialWords = []string{}
						break
					}
				}
				if len(specialWords) > 0 {
					specialWordsArray = append(specialWordsArray, specialWords)
				}
			}
		}
	}
	return c.JSON(http.StatusOK, specialWordsArray)
}

func sameCharter(word1, word2 string) bool {
	if len(word1) == len(word2) {
		word := word1 + word2
		wordMap := map[string]int{}
		for _, char := range word {
			wordMap[string(char)]++
		}
		for _, count := range wordMap {
			if count%2 != 0 {
				return false
			}
		}
	} else {
		return false
	}
	return true
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/users/order", getUsersSalesOrder)
	e.POST("/users/order/rating", provideUsersSalesOrderRating)
	e.GET("/special/words", getSpecialWords)

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}
