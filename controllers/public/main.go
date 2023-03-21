package public

import (
	"github.com/gin-gonic/gin"
)

//	@title			Cinnabar Public REST API
//	@version		BETA
//	@description	# Introduction (OpenAPI Specification)
//	@description	Cinnabar Public REST API merupakan API publik yang tersedia untuk umum dan dapat digunakan secara gratis. Dokumentasi API ini merupakan dokumentasi Cinnabar REST API untuk publik. Apabila anda ingin melihat dokumentasi Cinnabar REST API yang privat, silakan menuju tautan [Cinnabar Private REST API -  Documentation](https://cinnabar.icaksh.my.id/private/docs).
//	@description	API ini didokumentasikan dalam bentuk **OpenAPI format** dan didasarkan pada [swagger.io](https://swagger.io). Halaman dokumentasi ini digenerate dengan menggunakan [swaggo](https://github.com/swaggo/swag) dan dipublikasikan menggunakan [redoc](https://github.com/Redocly/redoc).
//	@description	# Cross-Origin Resource Sharing
//	@description	API ini memiliki fitur Cross-Origin Resource Sharing (CORS) yang diterapkan sesuai dengan [spesifikasi W3C](https://www.w3.org/TR/cors/) dan memungkinkan komunikasi lintas domain. Semua response bersifat publik dan dapat diakses oleh semua orang dari aplikasi apapun ataupun website apapun.
//	@description	# Information
//	@description	API ini masih dalam tahap pengembangan. Apabila anda ingin memberikan saran terhadap pengembangan API ini, silakan hubungi alamat pos-el yang telah tersedia. Jangan sungkan untuk menghubungi alamat pos-el jika terjadi error.
//	@description	# Donation
//	@description	Pada dasarnya anda dapat menggunakan REST API ini secara gratis. Perlu diketahui bahwa Cinnabar Public REST API di-deploy di Google Cloud Run Free Tier dan hingga saat ini biaya yang dikeluarkan hanya terfokus pada pembelian bandwith ke Indonesia (sekitar Rp1500/GB) sehingga kami tidak membuka donasi dan kami harap anda tidak melakukan hal yang mengakibatkan tingginya bandwith keluar.
//	@contact.name	Cinnabar REST API Support
//	@contact.url	https://cinnabar.icaksh.my.id/
//	@contact.email	me@icaksh.my.id
//	@license.name	MIT
//	@license.url	https://opensource.org/licenses/MIT

//	@host						cinnabar.icaksh.my.id
//	@BasePath					/public
//	@query.collection.format	multi
//	@x-tagGroups				[{"name": "API", "tags": ["basic"]},{"name": "Daily", "tags": ["information"]}]
//	@x-logo						{"url" : "https://cinnabar.icaksh.my.id/view/logo.png", "backgroundColor" : "#ffffff", "altText":"Cinnabar REST API"}
//	@scheme						["https", "http"]
func Register(router *gin.Engine) {
	publicGroup := router.Group("/public")
	// Basic
	//@tag.name basic
	publicGroup.GET("/test", GetBasicResponse())

	// Information
	//@tag.name informasi
	//@tag.description API untuk mendapatkan suatu informasi seperti fakta unik, berita terkini, atau informasi darurat seperti gempa bumi
	publicGroup.GET("/latest/gempa", GetGempaTerkini())
	publicGroup.GET("/daily/tawiki", GetTahukahAnda())

	// Beta
	//@tag.name beta
	//@tag.description API yang masih dalam tahap pengembangan

}
