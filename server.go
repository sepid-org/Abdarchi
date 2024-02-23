package main

import (
	// "bytes"
	// "encoding/json"
	// "fmt"
	"strings"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"regexp"
	// "io/ioutil"
)

// https://backend.kamva.academy/api/fsm/articles/?page=1
type RequestData struct {
	Data string `json:"data"`
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	
	e.Any("/api/:path*", SendReqGet)

	e.Logger.Fatal(e.Start(":8000"))
}

func SendReqGet(c echo.Context) error {
	// Retrieve the wildcard parameter
	otherURL := ""
	wildcardParam := c.Request().URL
	wildcardParam1 := wildcardParam.String()
	re := regexp.MustCompile(`/api/([^/]+)`)
	matches := re.FindStringSubmatch(wildcardParam1)

	if len(matches) > 1 {
		firstWordAfterAPI := matches[1] 
		if firstWordAfterAPI == "fsm"{
			otherURL = "https://backend.kamva.academy/api"
		}
		if firstWordAfterAPI == "party-manager"{
			otherURL = "https://mps.sepid.org/api"
		}
		// find service
	}
	
	resultString := strings.Replace(wildcardParam1, "api/", "", 1)
	// fmt.Println("parameter:", resultString)
	// fsm/articles/?page=1
	if otherURL == ""{
		return c.String(http.StatusOK, "سلام فرمانده!")
	}
	return c.Redirect(http.StatusMovedPermanently, otherURL + resultString)
	// resp, err := http.Get(otherURL)
	// if err != nil {
	// 	return c.String(http.StatusInternalServerError, "Error making POST request to other URL")
	// }
	// defer resp.Body.Close()
	// responseBody, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	//     return c.String(http.StatusInternalServerError, "Error reading response body from other URL")
	// }
	// responseString := string(responseBody)
	// return c.String(http.StatusOK, responseString)
	// }
	// return c.String(http.StatusOK, "Wildcard parameter: "+"lk")
}
