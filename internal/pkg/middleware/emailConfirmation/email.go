package emailsender

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

type EmailSender struct {
	emailSmtp     string
	emailSsl      int
	emailAddress  string
	emailPassword string
}

func New() (*EmailSender, error) {
	_ = godotenv.Load()
	ssl, err := strconv.Atoi(os.Getenv("EMAIL_SSL"))
	if err != nil {
		return &EmailSender{}, err
	}

	return &EmailSender{
		emailSmtp:     os.Getenv("EMAIL_SMTP"),
		emailSsl:      ssl,
		emailAddress:  os.Getenv("EMAIL_ADDRESS"),
		emailPassword: os.Getenv("EMAIL_PASSWORD"),
	}, err
}

func GenerateUniqueCode() string {
	rand.Seed(time.Now().UnixNano())

	characters := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	code := make([]byte, 5)

	for i := range code {
		code[i] = characters[rand.Intn(len(characters))]
	}

	return string(code)
}

func (es *EmailSender) SendConfirmationEmail(code string, recipientEmail string) error {
	message := gomail.NewMessage()
	message.SetHeader("From", "mdhh.hackathon@gmail.com")
	message.SetHeader("To", recipientEmail)
	message.SetHeader("Subject", "Код подтвержения регистрации")
	message.SetBody("text/plain", "Ваш уникальный код для подтверждения регистрации: "+code)
	fmt.Println(message)
	d := gomail.NewDialer(es.emailSmtp, es.emailSsl, es.emailAddress, es.emailPassword)
	fmt.Println(d)
	if err := d.DialAndSend(message); err != nil {
		return errors.New("ошибка при отправке письма")
	}

	return nil
}
