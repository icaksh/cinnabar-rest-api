package daily

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status	int	`json:"status"`
	Message	string `json:"message"`
	Data	map[string]interface{} `json:"data"`
}

type Builder struct {
	Info	string	`json:"info"`
}

func TahukahAnda() gin.HandlerFunc{
	return func(c *gin.Context) {
		wikipediaUrl := "https://id.wikipedia.org"
		client := &http.Client{}
		response, err := client.Get(wikipediaUrl)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				Response{
					Status: http.StatusInternalServerError,
					Message: "error",
					Data: map[string]interface{}{"data": err},
				},
			)
			return
		}
		defer response.Body.Close()
		document, err := goquery.NewDocumentFromReader(response.Body)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				Response{
					Status: http.StatusInternalServerError,
					Message: "error",
					Data: map[string]interface{}{"data": err},
				},
			)
			return
		}
		s := document.Find("#mf-tahukahanda ul")
		var data []Builder
		s.Find("li").Each(func(i int, s *goquery.Selection) {
			datax := Builder{
				Info: s.Text(),
			}
			data = append(data,datax)
		})
		c.JSON(
			http.StatusOK,
			Response{
				Status: http.StatusOK,
				Message: "success",
				Data: map[string]interface{}{"data": data},
			},
		)
	}
}