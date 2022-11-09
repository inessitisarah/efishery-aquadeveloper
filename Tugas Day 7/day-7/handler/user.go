package handler

import (
	"mymodule/entity/response"
	"mymodule/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userUsecase *usecase.UserUseCase
}

func NewUserHandler(usercase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{userUsecase: usercase}
}

func (h UserHandler) CreateUser(ctx *fiber.Ctx) error {
	userRequest := response.CreateUserRequest{}
	if err := ctx.BodyParser(&userRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "invalid req body",
			Data:    nil,
		})
	}
	if err := h.userUsecase.CreateUser(userRequest); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Failed to create user",
			Data:    nil,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.BaseResponse{
		Code:    fiber.StatusCreated,
		Message: "user created successfully",
		Data:    nil,
	})

}

func (h UserHandler) GetList(ctx *fiber.Ctx) error {

	users, err := h.userUsecase.GetList()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "get list users failed",
			Data:    nil,
		})
	}
	if len(users) == 0 {
		return ctx.Status(fiber.StatusNoContent).JSON(response.BaseResponse{
			Code:    fiber.StatusOK,
			Message: "successfully get all users",
			Data:    users,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response.BaseResponse{
		Code:    fiber.StatusOK,
		Message: "successfully get all users",
		Data:    users,
	})

}

func (h UserHandler) DeleteUser(ctx *fiber.Ctx) error {
	userId, _ := strconv.Atoi(ctx.Params("id"))
	err := h.userUsecase.DeleteUser(userId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Delete user failed",
			Data:    nil,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.BaseResponse{
		Code:    fiber.StatusOK,
		Message: "Delete user successfully",
	})
}

func (handler UserHandler) UpdateUser(ctx *fiber.Ctx) error {
	userId, _ := strconv.Atoi(ctx.Params("id"))
	userRequest := entity.UpdateUserRequest{}
	if err := ctx.BodyParser(&userRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid request body",
			Data:    nil,
		})
	}
	user, err := handler.userUsecase.UpdateUser(userRequest, userId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Update user failed",
			Data:    nil,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.BaseResponse{
		Code:    fiber.StatusOK,
		Message: "Update user successfully",
		Data:    user,
	})
}
