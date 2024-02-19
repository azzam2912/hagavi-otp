package util

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"hagavi-otp/config"
	"hagavi-otp/schema"
	"github.com/gofiber/fiber/v2"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func UpdateUserOTP(phoneNumber string, db *sql.DB, otp string, isOTPVerified bool) error {
	_, err := db.Exec(`UPDATE $1 SET otp = $2, is_otp_verified = $3 WHERE phone = $4`, config.Config("SQL_TABLE_NAME"), otp, isOTPVerified, phoneNumber)
	if err != nil {
		fmt.Println(err)
	} else {
		err = nil
	}
	return err
}

func GenerateRandomOTP() (string, error) {
	const length = 6
	const otpChars = "0123456789"
	temp := make([]byte, length)
	_, err := rand.Read(temp)
	if err != nil {
		return "", err
	}
	otpCharsLength := len(otpChars)
	for i := 0; i < length; i++ {
		temp[i] = otpChars[int(temp[i])%otpCharsLength]
	}
	return string(temp), nil
}

func SendOTP(c *fiber.Ctx, toPhone string, otp string) error {
	accountSid := config.Config("TWILIO_ACCOUNT_SID")
	authToken := config.Config("TWILIO_AUTH_TOKEN")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})
	params := &twilioApi.CreateMessageParams{}
	params.SetTo(toPhone)
	params.SetFrom(config.Config("TWILIO_PHONE_NUMBER"))
	params.SetBody("OTP sent from Hagavi!")

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("Error sending SMS message: " + err.Error())
		return err
	} 
	return c.Status(fiber.StatusAccepted).JSON(schema.ResponseHTTP{
		Success: true,
		Data:    nil,
		Message: "Otp sent to registered mobile number",
	})
}
