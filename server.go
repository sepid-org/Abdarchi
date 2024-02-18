package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RequestData struct {
	Data string `json:"data"`
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/post/", handlePost)
	e.GET("/get/", handleGet)


	e.Logger.Fatal(e.Start(":8000"))
}

func handleGet(c echo.Context) error {
	return c.String(http.StatusOK, "سلام فرمانده!")
}

func handlePost(c echo.Context) error {
	var requestData RequestData
	if err := c.Bind(&requestData); err != nil {
		return c.String(http.StatusBadRequest, "Error parsing JSON request body")
	}

	data := requestData.Data

	responseString := fmt.Sprintf("Received POST request with data: %s", data)

	otherURL := "https://example.com/other-endpoint"
	requestDataBytes, err := json.Marshal(requestData)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error marshaling JSON data")
	}

	resp, err := http.Post(otherURL, "application/json", bytes.NewBuffer(requestDataBytes))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error making POST request to other URL")
	}
	defer resp.Body.Close()
	// responseBody, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	//     return c.String(http.StatusInternalServerError, "Error reading response body from other URL")
	// }
	// responseString = string(responseBody)
	return c.String(http.StatusOK, responseString)
}