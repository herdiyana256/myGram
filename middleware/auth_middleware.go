// package middleware: Package yang berisi fungsi-fungsi middleware yang digunakan dalam routing aplikasi.

import (
	"github.com/gin-gonic/gin" // Mengimport library gin yang digunakan untuk membuat server HTTP dengan Go.
	"myGram/utils"              // Mengimport package utils yang berisi utilitas yang digunakan dalam aplikasi.
)

// AuthMiddleware adalah middleware untuk melakukan autentikasi pengguna.
// Middleware ini memeriksa keberadaan header Authorization pada setiap request yang masuk.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization") // Mendapatkan nilai header Authorization dari request.
		if header == "" {
			utils.Unauthorized(c, "Authorization header is missing") // Jika header Authorization tidak ditemukan, maka respon Unauthorized dikirimkan kembali ke client.
			c.Abort() // Menghentikan eksekusi request lebih lanjut.
			return
		}

		// tokenString := strings.Split(header, " ")[1]
		// Validate JWT token here
		// Validasi token JWT dapat dilakukan di sini.

		// Lanjutkan ke middleware berikutnya.
		c.Next()
	}
}

// AuthorizePhoto adalah middleware untuk melakukan otorisasi terhadap modifikasi foto.
// Middleware ini memeriksa apakah pengguna memiliki izin untuk memodifikasi foto yang dimaksud.
func AuthorizePhoto() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Periksa apakah pengguna memiliki izin untuk memodifikasi foto.
		// Logika otorisasi dapat diimplementasikan berdasarkan kebutuhan aplikasi.
		// Sebagai contoh, periksa apakah ID pengguna dalam token sesuai dengan ID pengguna yang terkait dengan foto.
		// Jika tidak diizinkan, kembalikan respon 403 Forbidden.
		// if !authorized {
		// 	utils.Forbidden(c, "Not authorized to modify this photo")
		// 	c.Abort()
		// 	return
		// }

		// Lanjutkan ke middleware berikutnya.
		c.Next()
	}
}

// Implement AuthorizeComment and AuthorizeSocialMedia middleware similarly
// Anda dapat mengimplementasikan middleware AuthorizeComment dan AuthorizeSocialMedia dengan cara yang serupa.
// Middleware ini digunakan untuk melakukan otorisasi terhadap komentar dan media sosial.
