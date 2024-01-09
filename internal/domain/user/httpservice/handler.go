package httpservice

import (
	userDatatype "github.com/Redchlorophyll/expense-tracker/internal/domain/user/datatype"
	"github.com/gofiber/fiber/v2"
)

func NewHandler(cfg UserServiceHandlerConfig) *UserServiceHandler {
	return &UserServiceHandler{
		userService: cfg.UserService,
	}
}

func (user *UserServiceHandler) CreateUserHandler(fiberCtx *fiber.Ctx) error {
	var request userDatatype.UserRequest

	err := fiberCtx.BodyParser(&request)

	if err != nil {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(userDatatype.GeneralResponse{
			StatusCode: 400,
			Message:    err.Error(),
		})
	}

	result, err := user.userService.CreateUser(fiberCtx.Context(), request)

	if err != nil {
		return fiberCtx.Status(fiber.StatusInternalServerError).JSON(userDatatype.GeneralResponse{
			StatusCode: 500,
			Message:    "there is something wrong in the system, please contact developer",
		})
	}

	return fiberCtx.Status(fiber.StatusOK).JSON(result)
}

func (user *UserServiceHandler) LoginHandler(fiberCtx *fiber.Ctx) error {
	var request userDatatype.UserRequest

	err := fiberCtx.BodyParser(&request)

	if err != nil {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(userDatatype.GeneralResponse{
			StatusCode: 400,
			Message:    err.Error(),
		})
	}

	result, err := user.userService.GetUser(fiberCtx.Context(), request)

	if err != nil {
		return fiberCtx.Status(fiber.StatusInternalServerError).JSON(userDatatype.GeneralResponse{
			StatusCode: 500,
			Message:    "there is something wrong in the system, please contact developer",
		})
	}

	return fiberCtx.Status(fiber.StatusOK).JSON(result)
}
