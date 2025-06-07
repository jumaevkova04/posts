package service

import (
	"fmt"
	"github.com/jumaevkova04/posts/internal/dto"
	"github.com/jumaevkova04/posts/internal/models"
	"github.com/jumaevkova04/posts/pkg/crypto"
	"github.com/jumaevkova04/posts/pkg/password"
	"github.com/jumaevkova04/posts/pkg/response"
	"time"
)

var (
	otpToRegisterKeyPrefix = "register_otp_email_"
)

func (s *Service) OTPToRegister(input *dto.OTPToRegisterRequest) error {
	if err := input.Validate(); err != nil {
		return response.NewFailedValidationError(err)
	}

	if err := s.checkUniqueUserEmail(input.Email); err != nil {
		return err
	}

	otpToRegisterKey := otpToRegisterKeyPrefix + input.Email

	otp, err := crypto.GenerateOTP(crypto.DefaultOtpLength)
	if err != nil {
		return response.NewServerError(err)
	}

	if err := s.RedisDB.Set(otpToRegisterKey, otp, time.Minute*10); err != nil {
		return response.NewServerError(err)
	}

	subject := "OTP to register"
	payload := fmt.Sprintf("Your OTP to register is %s", otp)

	err = s.Mailer.Send(input.Email, subject, payload)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) CheckRegisterOTP(input *dto.CheckRegisterOTPRequest) (*dto.CheckRegisterOTPResponse, error) {
	if err := input.Validate(); err != nil {
		return nil, response.NewFailedValidationError(err)
	}

	otpToRegisterKey := otpToRegisterKeyPrefix + input.Email

	otp, err := s.RedisDB.Get(otpToRegisterKey)
	if err != nil {
		return nil, response.NewServerError(err)
	}

	if otp == "" {
		return nil, response.NewBadRequestError(response.ErrInvalidEmail)
	}

	if otp != input.Otp {
		return nil, response.NewBadRequestError(response.ErrInvalidOTP)
	}

	nonce, err := crypto.GenerateTokenBase64(input.Email)
	if err != nil {
		return nil, response.NewServerError(err)
	}

	if err := s.RedisDB.Set(nonce, input.Email, time.Hour); err != nil {
		return nil, response.NewServerError(err)
	}

	_ = s.RedisDB.Delete(otpToRegisterKey)

	data := &dto.CheckRegisterOTPResponse{Nonce: nonce}

	return data, nil
}

func (s *Service) Register(input *dto.RegisterRequest) error {
	if err := input.Validate(); err != nil {
		return response.NewFailedValidationError(err)
	}

	email, err := s.RedisDB.Get(input.Nonce)
	if err != nil {
		return response.NewServerError(err)
	}

	if email == "" {
		return response.NewBadRequestError(response.ErrInvalidNonce)
	}

	if err := s.checkUniqueUserEmail(input.Email); err != nil {
		return err
	}

	hashedPassword, err := password.Hash(input.Password)
	if err != nil {
		return response.NewServerError(err)
	}

	newUser := models.NewUser(
		input.Name,
		input.Surname,
		input.Patronymic,
		input.Email,
		hashedPassword,
	)

	err = s.Repository.CreateUser(newUser)
	if err != nil {
		return response.NewServerError(err)
	}

	return nil
}

func (s *Service) checkUniqueUserEmail(email string) error {
	existingUser, err := s.Repository.GetUserByEmail(email)
	if err != nil {
		return response.NewServerError(err)
	}

	if existingUser != nil {
		return response.NewBadRequestError(response.ErrEmailAlreadyInUse)
	}

	return nil
}
