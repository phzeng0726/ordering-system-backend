package repository

import (
	"context"
	"errors"
	"fmt"
	"ordering-system-backend/internal/domain"
	otp "ordering-system-backend/pkg/otp"
	"time"

	"github.com/mailgun/mailgun-go/v4"
	"gorm.io/gorm"
)

const (
	mailDomain    = "sandbox6a4dcc235c40420bb28f1e2b7c7c72df.mailgun.org"
	privateAPIKey = "677baac0b688af79dc3eac11105625a0-4b98b89f-598eb6b9"
	sender        = "pipi_ordering@noreply.pipi.ordering.com"
	expireTime    = 30 // 30分鐘

)

func sendVerificationMail(code string, recipient string) error {
	// Create an instance of the Mailgun Client
	mg := mailgun.NewMailgun(mailDomain, privateAPIKey)

	subject := "Account Verification Code"
	codeHTML := fmt.Sprintf("<span style=\"font-size: 24px; font-weight: bold;\">%s</span>", code)
	body := fmt.Sprintf(`
	Dear User,<br><br>Your verification Code is:<br><br>%s<br><br>This code will expire in %d minutes, so please use it as soon as possible to activate your account.<br>
	If you did not request this code, please ignore this email.<br><br><br>Best regards,<br>PiPi Ordering System
	`, codeHTML, expireTime)

	// The message object allows you to add attachments and Bcc recipients
	message := mg.NewMessage(sender, subject, body, recipient)
	message.SetHtml(body)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message with a 10 second timeout
	resp, id, err := mg.Send(ctx, message)

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
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

	if err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&otp).Error; err != nil {
			return err
		}

		if err := sendVerificationMail(code, email); err != nil {
			return errors.New("invalid email for mailgun")
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (r *OTPRepo) Verify(ctx context.Context, token string, password string) error {
	var otp domain.OTP
	res := r.db.WithContext(ctx).Where("token = ?", token).First(&otp)
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
