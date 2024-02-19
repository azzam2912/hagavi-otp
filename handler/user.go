package handler

import (
	"database/sql"
	"hagavi-otp/schema"
	"hagavi-otp/util"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx, db *sql.DB) error {
	body := new(schema.RegisterBody)
	err := c.BodyParser(body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schema.ResponseHTTP{
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
	if user != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schema.ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: "Phone number has been already registered",
		})
	}

	err = util.AddUser(body, db)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(schema.ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(schema.ResponseHTTP{
		Success: true,
		Data: nil,
		Message: "Phone Number Registered",
	})
}

func Login(c *fiber.Ctx, db *sql.DB) error {
	body := new(schema.LoginSchema)
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

	if user.IsOTPVerified {
		if user.Password == body.Password {
			return c.Status(fiber.StatusOK).JSON(schema.ResponseHTTP{
				Success: true,
				Data:    nil,
				Message: "User Login Verified",
			})
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(schema.ResponseHTTP{
				Success: false,
				Data:    nil,
				Message: "Password incorrect",
			})
		}
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(schema.ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: "OTP Is Not Verified",
		})
	}
}



