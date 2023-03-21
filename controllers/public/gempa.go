package public

import (
	model "cinnabar/models"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

type TResFromBMKGInfo struct {
	DateTime    string `json:"DateTime"`
	Coordinates string `json:"Coordinates"`
	Latitute    string `json:"Lintang"`
	Longitude   string `json:"Bujur"`
	Magnitude   string `json:"Magnitude"`
	Wilayah     string `json:"Wilayah"`
	Potensi     string `json:"Potensi"`
	Dirasakan   string `json:"Dirasakan"`
	Shakemap    string `json:"Shakemap"`
	Depth       string `json:"Kedalaman"`
}

type TResFromBMKG struct {
	InfoGempa TResFromBMKGEarthquake `json:"Infogempa"`
}
type TResFromBMKGEarthquake struct {
	Gempa TResFromBMKGInfo `json:"gempa"`
}
type TInfoEarthquake struct {
	// unix timestamp (utc)
	Date int `json:"time"`
	// garis lintang
	Latitude float32 `json:"latitude"`
	// garis bujur
	Longitude float32 `json:"longitude"`
	// koordinat
	Coordinates string `json:"coordinates"`
	// magnitudo (SR)
	Magnitude float32 `json:"magnitude"`
	// meter
	Depth       int    `json:"depth"`
	// affected
	Affected    string `json:"affected"`
	// information
	Information string `json:"information"`
}

type TResLatestEarthquake struct {
	// latest earthquake information
	Information TInfoEarthquake `json:"info"`
	// latest earthquake map
	ImageLink   string     `json:"peta_gempa"`
}

type TErrLatestEarthquake struct {
	// kode error untuk dikirimkan ke pengelola
	Code int `json:"code"`
	// pesan error untuk dikirimkan ke pengelola
}

//	@Summary		BMKG - Gempa Terbaru
//	@Description	Sumber: Badan Meteorologi, Klimatologi, dan Geofisika Republik Indonesia
//	@Tags			information
//	@Produce		json
//	@Success		200	{object}	TResLatestEarthquake
//	@Failure		500	{object}	TErrLatestEarthquake 
//	@Router			/latest/gempa [get]
func GetGempaTerkini() gin.HandlerFunc {
	return func(c *gin.Context) {
		uBMKG := "https://data.bmkg.go.id/DataMKG/TEWS/autogempa.json"
		client := &http.Client{}
		if res, err := client.Get(uBMKG); err != nil {
			c.JSON(
				http.StatusInternalServerError,
				model.TErrorResponse{
					Error:    map[string]interface{}{"data": err},
				},
			)
			return
		} else {
			defer res.Body.Close()
			if doc, err := goquery.NewDocumentFromReader(res.Body); err != nil {
				c.JSON(
					http.StatusInternalServerError,
					model.TErrorResponse{
						Error:    map[string]interface{}{"data": err},
					},
				)
				return
			} else {
				textBytes := []byte(doc.Text())
				res := TResFromBMKG{}
				var info TInfoEarthquake
				if jsonErr := json.Unmarshal(textBytes, &res); jsonErr != nil {
					c.JSON(
						http.StatusInternalServerError,
						model.TErrorResponse{
							Error:    map[string]interface{}{"data": jsonErr},
						},
					)
					return
				} else {
					timestamp, err := time.Parse(time.RFC3339, res.InfoGempa.Gempa.DateTime)
					if err != nil {
						c.JSON(
							http.StatusInternalServerError,
							model.TErrorResponse{
								Error:    map[string]interface{}{"data": err},
							},
						)
						return
					}
					magnitude, err := strconv.ParseFloat(res.InfoGempa.Gempa.Magnitude, 32)
					if err != nil {
						c.JSON(
							http.StatusInternalServerError,
							model.TErrorResponse{
								Error:    map[string]interface{}{"data": err},
							},
						)
						return
					}
					latitude, err := strconv.ParseFloat(res.InfoGempa.Gempa.Latitute[:len(res.InfoGempa.Gempa.Latitute)-3], 32)
					if err != nil {
						c.JSON(
							http.StatusInternalServerError,
							model.TErrorResponse{
								Error:    map[string]interface{}{"data": err},
							},
						)
						return
					}
					longitude, err := strconv.ParseFloat(res.InfoGempa.Gempa.Longitude[:len(res.InfoGempa.Gempa.Longitude)-3], 32)
					if err != nil {
						c.JSON(
							http.StatusInternalServerError,
							model.TErrorResponse{
								Error:    map[string]interface{}{"data": err},
							},
						)
						return
					}
					kedalaman, err := strconv.ParseFloat(res.InfoGempa.Gempa.Depth[:len(res.InfoGempa.Gempa.Depth)-3], 32)
					if err != nil {
						c.JSON(
							http.StatusInternalServerError,
							model.TErrorResponse{
								Error:    map[string]interface{}{"data": err},
							},
						)
						return
					}
					data := TInfoEarthquake{
						Date:        int(timestamp.Unix()),
						Magnitude:   float32(magnitude),
						Latitude:    float32(latitude),
						Longitude:   float32(longitude),
						Depth:       int(kedalaman) * 1000,
						Coordinates: res.InfoGempa.Gempa.Coordinates,
						Information: res.InfoGempa.Gempa.Wilayah + ", " + res.InfoGempa.Gempa.Potensi,
						Affected:    res.InfoGempa.Gempa.Dirasakan,
					}
					info = data
				}
				result := TResLatestEarthquake{
					Information: info,
					ImageLink:   "https://data.bmkg.go.id/DataMKG/TEWS/" + res.InfoGempa.Gempa.Shakemap,
				}
				c.JSON(
					http.StatusOK,
					model.TSuccessResponse{
						Data:    result,
					},
				)
			}
		}
	}
}
