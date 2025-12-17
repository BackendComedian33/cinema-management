package handler

import (
	"net/http"
	"technical-test/config"
	"technical-test/dto"
	"technical-test/dto/request"
	"technical-test/model"
	"technical-test/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type UserHandler struct {
	Env     *config.EnvironmentVariable
	Service service.UserService
}

func NewUserHandler(
	env *config.EnvironmentVariable,
	userService service.UserService,
) UserHandler {
	return UserHandler{
		Env:     env,
		Service: userService,
	}
}

func (h *UserHandler) Login(ctx *gin.Context) {
	var req request.LoginRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ApiResponse{
			StatusCode: http.StatusBadRequest,
			Success:    false,
		})
		return
	}
	log.Info().
		Interface("req", req).
		Msg("payload is valid")

	res := h.Service.Login(ctx, req)
	if !res.Success {
		ctx.JSON(res.StatusCode, res)
		return
	}

	ctx.JSON(res.StatusCode, res)

}

func (h *UserHandler) CreateShowtime(ctx *gin.Context) {
	var req request.CreateShowtimeRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Error().Err(err)
		ctx.JSON(http.StatusBadRequest, dto.ApiResponse{
			StatusCode: http.StatusBadRequest,
			Success:    false,
			Error:      err,
			ErrorCode:  40001,
		})
		return
	}

	// Convert to time.Time
	showDate, _ := time.Parse("2006-01-02", req.ShowDate)
	startTime, _ := time.Parse("15:04", req.StartTime)

	payload := model.Showtime{
		MovieID:         req.MovieID,
		StudioID:        req.StudioID,
		ShowDate:        showDate,
		StartTime:       startTime,
		DurationMinutes: 20,
		Status:          req.Status,
	}

	log.Info().
		Interface("req", payload).
		Msg("payload is valid")

	res := h.Service.Create(ctx, payload)
	if !res.Success {
		ctx.JSON(res.StatusCode, res)
		return
	}

	ctx.JSON(res.StatusCode, res)
}

func (h *UserHandler) GetShowtimeByID(ctx *gin.Context) {
	var req request.GetShowtimeRequest

	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ApiResponse{
			StatusCode: http.StatusBadRequest,
			Success:    false,
		})
		return
	}

	res := h.Service.GetByID(ctx, req)
	if !res.Success {
		ctx.JSON(res.StatusCode, res)
		return
	}

	ctx.JSON(res.StatusCode, res)
}

func (h *UserHandler) UpdateShowtime(ctx *gin.Context) {
	var req request.UpdateShowtimeRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ApiResponse{
			StatusCode: http.StatusBadRequest,
			Success:    false,
		})
		return
	}

	log.Info().
		Interface("req", req).
		Msg("payload is valid")

	showDate, _ := time.Parse("2006-01-02", req.ShowDate)
	startTime, _ := time.Parse("15:04", req.StartTime)

	payload := model.Showtime{
		ID:              int64(req.ID),
		MovieID:         req.MovieID,
		StudioID:        req.StudioID,
		ShowDate:        showDate,
		StartTime:       startTime,
		DurationMinutes: 20,
		Status:          req.Status,
	}

	res := h.Service.Update(ctx, payload)
	if !res.Success {
		ctx.JSON(res.StatusCode, res)
		return
	}

	ctx.JSON(res.StatusCode, res)
}

func (h *UserHandler) DeleteShowtime(ctx *gin.Context) {
	var req request.DeleteShowtimeRequest

	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ApiResponse{
			StatusCode: http.StatusBadRequest,
			Success:    false,
		})
		return
	}

	res := h.Service.Delete(ctx, req)
	if !res.Success {
		ctx.JSON(res.StatusCode, res)
		return
	}

	ctx.JSON(res.StatusCode, res)
}

func (h *UserHandler) GetAllShowtime(ctx *gin.Context) {
	res := h.Service.GetAllShowtime(ctx)
	ctx.JSON(res.StatusCode, res)
}
