package config

import (
	"fmt"
	"net/smtp"
	"os"
)

// SendVerificationEmail mengirim email berisi link verifikasi ke penyewa baru.
// Dipanggil secara goroutine (fire and forget) dari service.
func SendVerificationEmail(toEmail, toName, token string) {
	from    := os.Getenv("SMTP_FROM")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")
	appURL   := os.Getenv("APP_URL") // contoh: https://api.sewalapangan.com

	verifyURL := fmt.Sprintf("%s/api/auth/verify-email?token=%s", appURL, token)

	subject := "Verifikasi Email - Sewa Lapangan"
	body := fmt.Sprintf(`Halo %s,

Terima kasih sudah mendaftar di Sewa Lapangan!

Klik link berikut untuk memverifikasi email kamu:
%s

Link ini berlaku selama 24 jam.

Jika kamu tidak mendaftar, abaikan email ini.

Salam,
Tim Sewa Lapangan`, toName, verifyURL)

	msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s",
		from, toEmail, subject, body)

	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)
	addr := fmt.Sprintf("%s:%s", smtpHost, smtpPort)

	if err := smtp.SendMail(addr, auth, from, []string{toEmail}, []byte(msg)); err != nil {
		// Log error tanpa crash — user sudah terdaftar, bisa resend nanti
		fmt.Printf("[EMAIL ERROR] gagal kirim ke %s: %v\n", toEmail, err)
	}
}
