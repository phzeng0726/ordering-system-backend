package repository

import (
	"context"
	"errors"
	"fmt"
	"net/smtp"
	"ordering-system-backend/internal/domain"
	otp "ordering-system-backend/pkg/otp"

	"gorm.io/gorm"
)

const (
	senderEmail    = "mgc881017@gmail.com"
	senderPassword = "wsdt nnwk hpgh zaoj"
)

func sendVerificationMail(code string, recipient string) error {
	// Configuration
	smtpServer := "smtp.gmail.com"
	smtpPort := "587"
	expireTime := 30 // Mins

	// Message.
	subject := "Account Verification Code"
	codeHTML := fmt.Sprintf("<span style=\"font-size: 24px; font-weight: bold;\">%s</span>", code)
	message := fmt.Sprintf(`
	Dear User,<br><br>Your verification Code is:<br><br>%s<br><br>This code will expire in %d minutes, so please use it as soon as possible to activate your account.<br>
	If you did not request this code, please ignore this email.<br><br><br>Best regards,<br>PiPi Ordering System
	`, codeHTML, expireTime)
	msgStyle := "MIME-version: 1.0\r\n" + "Content-Type: text/html; charset=\"UTF-8\"\r\n\r\n"

	// Connect to SMTP server
	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpServer)
	to := []string{recipient}
	msg := []byte("Subject: " + subject + "\r\n" + msgStyle + message)

	// Sending email.
	if err := smtp.SendMail(smtpServer+":"+smtpPort, auth, senderEmail, to, msg); err != nil {
		return err
	}

	return nil
}

type OTPRepo struct {
	db *gorm.DB
}

func NewOTPRepo(db *gorm.DB) *OTPRepo {
	return &OTPRepo{
		db: db,
	}
}

func (r *OTPRepo) Create(ctx context.Context, token string, email string) error {
	code := otp.GenerateRandomCode(6)
	var otp domain.OTP
	otp.Token = token
	otp.Password = code
	otp.Email = email
	db := r.db.WithContext(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&otp).Error; err != nil {
			return err
		}

		if err := sendVerificationMail(code, email); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (r *OTPRepo) Verify(ctx context.Context, token string, password string) error {
	var otp domain.OTP
	db := r.db.WithContext(ctx)

	res := db.Where("token = ?", token).First(&otp)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return errors.New("token not found")
		}
		return res.Error
	}

	if otp.Password != password {
		return errors.New("invalid otp code")
	}

	return nil
}
