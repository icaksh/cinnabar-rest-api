package private

import (
	"github.com/gin-gonic/gin"
)

//	@title			Cinnabar Private REST API
//	@version		BETA
//	@description	# Introduction (OpenAPI Specification)
//	@description	Cinnabar Private REST API merupakan API private yang tersedia untuk umum dan dapat digunakan secara secara berbayar. Dokumentasi API ini merupakan dokumentasi Cinnabar REST API untuk privat. Apabila anda ingin melihat dokumentasi Cinnabar REST API yang privat, silakan menuju tautan [Cinnabar Public REST API -  Documentation](https://cinnabar.icaksh.my.id/public/docs).
//	@description	API ini didokumentasikan dalam bentuk **OpenAPI format** dan didasarkan pada [swagger.io](https://swagger.io). Halaman dokumentasi ini digenerate dengan menggunakan [swaggo](https://github.com/swaggo/swag) dan dipublikasikan menggunakan [redoc](https://github.com/Redocly/redoc).
//	@description	# Cross-Origin Resource Sharing
//	@description	API ini memiliki fitur Cross-Origin Resource Sharing (CORS) yang diterapkan sesuai dengan [spesifikasi W3C](https://www.w3.org/TR/cors/) dan memungkinkan komunikasi lintas domain. Semua response bersifat publik dan dapat diakses oleh semua orang dari aplikasi apapun ataupun website apapun.
//	@description	# Information
//	@description	API ini masih dalam tahap pengembangan. Apabila anda ingin memberikan saran terhadap pengembangan API ini, silakan hubungi alamat pos-el yang telah tersedia. Jangan sungkan untuk menghubungi alamat pos-el jika terjadi error

//	@securityDefinitions.apikey	apikey
//	@description				For this sample, you can use the api key `special-key` to test the authorization filters.
//	@name						apikey
//	@in							header

//	@contact.name	Cinnabar REST API Support
//	@contact.url	https://cinnabar.icaksh.my.id/
//	@contact.email	me@icaksh.my.id

//	@license.name	MIT
//	@license.url	https://opensource.org/licenses/MIT

//	@host						cinnabar.icaksh.my.id
//	@BasePath					/public
//	@query.collection.format	multi
//	@x-tagGroups				[{"name": "API", "tags": ["basic"]}]
//	@x-logo						{"url" : "https://cinnabar.icaksh.my.id/view/logo.png", "backgroundColor" : "#ffffff", "altText":"Cinnabar REST API"}
//	@scheme						["https", "http"]

func Register(router *gin.Engine) {
	privateGroup := router.Group("/private")
	// Basic
	//@tag.name basic
	privateGroup.POST("/api/test", PostBasicResponse())
	privateGroup.POST("/api/info", PostCheckApiInfo())

}
