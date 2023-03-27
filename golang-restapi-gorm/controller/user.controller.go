package controller

import (
	"log"

	"github.com/Artzy1401/golang-restapi-gorm/database"
	"github.com/Artzy1401/golang-restapi-gorm/model/entity"
	"github.com/Artzy1401/golang-restapi-gorm/model/request"
	"github.com/Artzy1401/golang-restapi-gorm/utils"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func UserControllerGetAll(ctx *fiber.Ctx) error {
	// userInfo := ctx.Locals("userInfo")
	// log.Println("user info data :: ", userInfo)
	var User []entity.User
	result := database.DB.Find(&User)
	if result.Error != nil {
		log.Println(result.Error)
	}

	// err := database.DB.Find(&User).Error
	// if err != nil {
	// 	log.Println(err)
	// }

	return ctx.JSON(User)
}

func UserControllerCreate(ctx *fiber.Ctx) error {
	User := new(request.UserCreateRequest)
	if err := ctx.BodyParser(User); err != nil {
		return err
	}

	// VALIDASI REQUEST
 	validate := validator.New()
	errValidate := validate.Struct(User)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message" : "failed to validate",
			"error" : errValidate.Error(),
		})
	}

	newUser := entity.User{
		Name:		User.Name,
		Email:		User.Email,
	}

	hashedPassword, err := utils.HashingPassword(User.Password)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : "internal server error",
		})
	}

	newUser.Password = hashedPassword

	errCreateUser := database.DB.Create(&newUser).Error
	if errCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to create user",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "successfully",
		"data": newUser,
	})
}


func UserControllerGetById(ctx *fiber.Ctx) error{
	UserId := ctx.Params("id")

	var User entity.User
	err := database.DB.First(&User, "id = ?", UserId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "successfully",
		"data": User,
	})
}

func UserControllerUpdate(ctx *fiber.Ctx) error {
	UserRequest := new(request.UserUpdateRequest)
	if err := ctx.BodyParser(UserRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	var User entity.User

	UserId := ctx.Params("id")
	// CHECK AVALAIBLE User
	err := database.DB.First(&User, "id = ?", UserId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// UPDATE USER DATA
	if UserRequest.Name != "" {
		User.Name = UserRequest.Name
	}
	User.Email = UserRequest.Email
	User.Password = UserRequest.Password

	errUpdate := database.DB.Save(&User).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "successfully",
		"data": User,
	})
}

/*func UserControllerUpdateEmail(ctx *fiber.Ctx) error {
	UserRequest := new(request.UserEmailRequest)
	if err := ctx.BodyParser(UserRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	var User entity.User
	var EmailUser entity.User

	UserId := ctx.Params("id")
	// CHECK AVALAIBLE User
	err := database.DB.First(&User, "id = ?", UserId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// CHECK AVAILABLE EMAIL
	errCheckEmail := database.DB.First(&EmailUser, "email = ?", *UserRequest.Email).Error
	if errCheckEmail != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Email not found",
		})
	}
	
	User.Email = UserRequest.Email

	errUpdate := database.DB.Save(&User).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "successfully",
		"data": User,
	})
} */

func UserControllerDelete(ctx *fiber.Ctx) error {
	UserId := ctx.Params("id")
	var User entity.User

	// CHECK AVAILABLE User
	err := database.DB.Debug().First(&User, "id=?" ,&UserId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "User Not Found",
		})
	}

	errDelete := database.DB.Debug().Delete(&User).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "User deleted",
	})
}