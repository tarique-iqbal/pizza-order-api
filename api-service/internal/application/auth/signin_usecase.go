package auth

import (
	"api-service/internal/domain/auth"
	"api-service/internal/domain/user"
	"api-service/internal/infrastructure/security"
	"errors"
)

type SignInUseCase struct {
	repo   user.UserRepository
	hasher auth.PasswordHasher
}

func NewSignInUseCase(repo user.UserRepository, hasher auth.PasswordHasher) *SignInUseCase {
	return &SignInUseCase{repo: repo, hasher: hasher}
}

func (uc *SignInUseCase) Execute(email string, password string) (string, error) {
	user, err := uc.repo.FindByEmail(email)
	if user == nil {
		return "", errors.New("no record found")
	}

	if err != nil {
		return "", errors.New("internal server error")
	}

	status := uc.hasher.Compare(user.Password, password)
	if !status {
		return "", errors.New("invalid credentials")
	}

	token, err := security.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
