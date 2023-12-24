package user

import (
	"github.com/google/uuid"
	"twitter/internal/core/models"
	"twitter/internal/core/providers"
	"twitter/internal/dal/repositories"
	"twitter/internal/dal/schemas"
	"twitter/internal/endpoints/user/dtos"
	"twitter/internal/endpoints/user/validation"
)

type Service struct {
	userRepository   *repositories.User
	responseProvider *providers.Response
	updateValidator  *validation.Update
}

func NewService(userRepository *repositories.User,
	responseProvider *providers.Response,
	updateValidator *validation.Update) *Service {
	return &Service{
		userRepository:   userRepository,
		responseProvider: responseProvider,
		updateValidator:  updateValidator,
	}
}

func (service *Service) CreateUser(dto dtos.CreateUser) *models.Response {
	if service.userRepository.ExistsByUsername(dto.Username) {
		return service.responseProvider.UserResponse.UsernameTaken()
	}

	user := schemas.NewUser(uuid.New(), dto.Username, dto.Password)
	service.userRepository.Create(user)
	return service.responseProvider.UserResponse.Created()
}

func (service *Service) RequestUser(username string) *models.Response {
	user := service.userRepository.GetByUsername(username)

	if user == nil {
		return service.responseProvider.UserResponse.NotFound()
	}

	return service.responseProvider.UserResponse.UserInfo(user)
}

func (service *Service) UpdateUser(dto dtos.UpdateUser) *models.Response {
	user := service.userRepository.GetById(dto.Id)

	if user == nil {
		return service.responseProvider.UserResponse.NotFound()
	}

	user.Username = dto.Username
	user.Password = dto.Password
	user.Following = dto.Following
	user.Followers = dto.Followers

	validationResult := service.updateValidator.ValidateUserUpdate(user)

	switch validationResult {
	case validation.Valid:
		service.userRepository.Update(user)
		return service.responseProvider.UserResponse.Updated()

	case validation.FollowingInvalid:
		return service.responseProvider.UserResponse.InvalidFollowing(user.Following)

	case validation.FollowersInvalid:
		return service.responseProvider.UserResponse.InvalidFollowers(user.Followers)

	default:
		return service.responseProvider.GeneralResponse.InternalServerError()
	}
}
