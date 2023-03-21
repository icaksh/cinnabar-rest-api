package public

import (
	model "cinnabar/models"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

type TInfoTAWiki struct {
	// urutan
	Id int `json:"id"`
	// informasi
	DidYouKnow string `json:"tahukah_anda"`
}

type TResTAWiki struct {
	// data "tahukah anda"
	Information []TInfoTAWiki `json:"info"`
	// tautan gambar (opsional)
	ImageLink string `json:"image_link"`
}

type TErrTAWiki struct {
	// kode error untuk dikirimkan ke pengelola
	Code int `json:"code"`
	// pesan error untuk dikirimkan ke pengelola
	Message string `json:"message"`
}

//	@Summary		Wikipedia - Tahukah Anda
//	@Description	Sumber: Wikipedia Bahasa Indonesia
//	@Tags			information
//	@Produce		json
//	@Success		200	{object}	TResTAWiki
//	@Failure		500	{object}	TErrTAWiki
//	@Router			/daily/tawiki [get]
func GetTahukahAnda() gin.HandlerFunc {
	return func(c *gin.Context) {
		uWiki := "https://id.wikipedia.org"
		client := &http.Client{}
		if res, err := client.Get(uWiki); err != nil {
			c.JSON(http.StatusInternalServerError, model.TErrorResponse{Error: TErrTAWiki{
				Code: 1,
				Message: "Cannot open Wikipedia Indonesia",
			}})
			return
		} else {
			defer res.Body.Close()
			if doc, err := goquery.NewDocumentFromReader(res.Body); err != nil {
				c.JSON(http.StatusInternalServerError, model.TErrorResponse{Error: TErrTAWiki{
					Code: 2,
					Message: "Cannot Fetch Pages",
				}})
				return
			} else {
				p := doc.Find("#mf-tahukahanda")
				s := p.Find("ul")
				var infos []TInfoTAWiki
				s.Find("li").Each(func(i int, s *goquery.Selection) {
					info := TInfoTAWiki{
						Id:         i,
						DidYouKnow: s.Text(),
					}
					infos = append(infos, info)
				})
				if imgLink, ok := p.Find("img").Attr("src"); ok {
					c.JSON(http.StatusOK, model.TSuccessResponse{Data: TResTAWiki{infos, "https:" + imgLink}})
				} else {
					c.JSON(http.StatusOK, model.TSuccessResponse{Data: TResTAWiki{infos, "null"}})
				}
			}
		}
	}
}
