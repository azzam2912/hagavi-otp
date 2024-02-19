package handler

import (
	"database/sql"
	"hagavi-otp/schema"
	"hagavi-otp/util"

	"github.com/gofiber/fiber/v2"
)

func VerifyOTP(c *fiber.Ctx, db *sql.DB) error {
	// request body data
	body := new(schema.VerifyOTP) 
	err := c.BodyParser(body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schema.ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: err.Error(),
		})
	}

	// find phone in database
	user, err := util.FindUserByPhoneNumber(body.Phone, db)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(schema.ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: err.Error(),
		})
	}

	if user == nil {
		return c.Status(fiber.StatusBadRequest).JSON(schema.ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: "Phone number does not exists",
		})
	}

	if user.Otp != body.Otp {
		return c.Status(fiber.StatusBadRequest).JSON(schema.ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: "Incorrect Otp",
		})
	}

	// remove old otp from db
	util.UpdateUserOTP(user.Phone, db, body.Otp, true)

	return c.Status(fiber.StatusOK).JSON(schema.ResponseHTTP{
		Success: true,
		Data: nil,
		Message: "OTP Verified",
	})
}

func ResendOTP(c *fiber.Ctx, db *sql.DB) error {
	// request body data
	body := new(schema.VerifyOTP)
	if err := c.Status(fiber.StatusBadRequest).BodyParser(body); err != nil {
		return c.JSON(schema.ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: err.Error(),
		})
	}

	user, err := util.FindUserByPhoneNumber(body.Phone, db)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(schema.ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: err.Error(),
		})
	}

	if user == nil {
		return c.Status(fiber.StatusBadRequest).JSON(schema.ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: "Phone number does not exists",
		})
	}

	otp, err := util.GenerateRandomOTP()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(schema.ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: err.Error(),
		})
	}

	util.UpdateUserOTP(body.Phone, db, body.Otp, false)

	err = util.SendOTP(c, user.Phone, otp)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(schema.ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(schema.ResponseHTTP{
		Success: true,
		Data:    nil,
		Message: ,
	})
}
