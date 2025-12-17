package service

import (
	"errors"
	"strconv"
	"technical-test/config"
	"technical-test/database"
	"technical-test/dto"
	"technical-test/dto/request"
	"technical-test/helper"
	"technical-test/model"
	"technical-test/repository"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	Login(ctx *gin.Context, req request.LoginRequest) (res dto.ApiResponse)
	Create(ctx *gin.Context, req model.Showtime) (res dto.ApiResponse)
	GetByID(ctx *gin.Context, req request.GetShowtimeRequest) (res dto.ApiResponse)
	Update(ctx *gin.Context, req model.Showtime) (res dto.ApiResponse)
	Delete(ctx *gin.Context, req request.DeleteShowtimeRequest) (res dto.ApiResponse)
	GetAllShowtime(ctx *gin.Context) (res dto.ApiResponse)
}
type UserServiceImpl struct {
	Env                *config.EnvironmentVariable
	UsersRepository    repository.UserRepository
	ShowtimeRepository repository.ShowtimeRepository
	WrapDB             *database.WrapDB
}

func NewUserService(
	UsersRepo repository.UserRepository,
	showtimeRepo repository.ShowtimeRepository,
	wrapDB *database.WrapDB,
	env *config.EnvironmentVariable,
) UserService {
	return &UserServiceImpl{
		Env:                env,
		UsersRepository:    UsersRepo,
		ShowtimeRepository: showtimeRepo,
		WrapDB:             wrapDB,
	}
}

func (s *UserServiceImpl) Login(ctx *gin.Context, req request.LoginRequest) (res dto.ApiResponse) {

	userID, password, err := s.UsersRepository.Login(ctx, nil, req.Email)
	if err != nil {
		res = helper.BuildResponse(422, false, nil, errors.New("invalid user and password"), 42201)
		return
	}
	if !helper.CheckPasswordHash(req.Password, password) {
		res = helper.BuildResponse(422, false, nil, errors.New("invalid user and password"), 42201)
		return
	}
	idString := strconv.Itoa(userID)
	token, err := helper.GenerateToken(s.Env, idString)
	if err != nil {
		res = helper.BuildResponse(500, false, nil, err, 50001)
		return
	}
	data := map[string]string{
		"token": token,
	}
	res = helper.BuildResponse(200, true, data, nil, 0)
	return
}

func (s *UserServiceImpl) Create(ctx *gin.Context, req model.Showtime) (res dto.ApiResponse) {

	_, err := s.ShowtimeRepository.Create(ctx, nil, req)
	if err != nil {
		res = helper.BuildResponse(500, false, nil, err, 50001)
		return
	}

	res = helper.BuildResponse(201, true, nil, nil, 0)
	return
}

func (s *UserServiceImpl) GetByID(ctx *gin.Context, req request.GetShowtimeRequest) (res dto.ApiResponse) {

	data, err := s.ShowtimeRepository.GetByID(ctx, nil, int64(req.ID))
	if err != nil {
		res = helper.BuildResponse(404, false, nil, errors.New("showtime not found"), 40401)
		return
	}

	res = helper.BuildResponse(200, true, data, nil, 0)
	return
}

func (s *UserServiceImpl) Update(ctx *gin.Context, req model.Showtime) (res dto.ApiResponse) {
	_, err := s.ShowtimeRepository.GetByID(ctx, nil, req.ID)
	if err != nil {
		res = helper.BuildResponse(404, false, nil, errors.New("showtime not found"), 40401)
		return
	}

	err = s.ShowtimeRepository.Update(ctx, nil, req)
	if err != nil {
		res = helper.BuildResponse(500, false, nil, err, 50001)
		return
	}

	res = helper.BuildResponse(200, true, nil, nil, 0)
	return
}

func (s *UserServiceImpl) Delete(ctx *gin.Context, req request.DeleteShowtimeRequest) (res dto.ApiResponse) {
	_, err := s.ShowtimeRepository.GetByID(ctx, nil, int64(req.ID))
	if err != nil {
		res = helper.BuildResponse(404, false, nil, errors.New("showtime not found"), 40401)
		return
	}

	err = s.ShowtimeRepository.Delete(ctx, nil, int64(req.ID))
	if err != nil {
		res = helper.BuildResponse(500, false, nil, err, 50001)
		return
	}

	res = helper.BuildResponse(200, true, nil, nil, 0)
	return
}

func (s *UserServiceImpl) GetAllShowtime(ctx *gin.Context) (res dto.ApiResponse) {
	data, err := s.ShowtimeRepository.GetAllAvailable(ctx, nil)
	if err != nil {
		res = helper.BuildResponse(500, false, []model.Showtime{}, err, 50001)
		return
	}
	res = helper.BuildResponse(200, true, data, nil, 0)
	return

}
