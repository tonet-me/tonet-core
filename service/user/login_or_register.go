package userservice

import (
	"context"
	"fmt"
	"github.com/tonet-me/tonet-core/entity"
	userparam "github.com/tonet-me/tonet-core/param/user"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
)

func (s Service) LoginOrRegister(ctx context.Context, req userparam.LoginOrRegisterRequest) (*userparam.LoginOrRegisterResponse, error) {
	const op = richerror.OP("userservice.LoginOrRegister")
	userInfoFromToken, vErr := s.oAuthSvc.ValidationAndGetInfoFromToken(ctx, req.ProviderID, req.Token)
	if vErr != nil {
		return nil, richerror.New(richerror.WithOp(op),
			richerror.WithInnerError(vErr))
	}

	isExisted, takenUser, gErr := s.repo.IsUserExistByEmail(ctx, userInfoFromToken.Email)
	if gErr != nil {
		return nil, richerror.New(richerror.WithOp(op),
			richerror.WithInnerError(gErr))
	}

	var user entity.User
	var isNewUser bool
	if isExisted {
		user = takenUser
	} else {
		newUser, cErr := s.repo.CreateNewUser(ctx, entity.User{
			FirstName:       userInfoFromToken.FirstName,
			LastName:        userInfoFromToken.LastName,
			Email:           userInfoFromToken.Email,
			ProfilePhotoURL: userInfoFromToken.ProfilePhotoURL,
			Status:          entity.UserStatusActive,
			EmailVerified:   true,
		})
		if cErr != nil {
			return nil, richerror.New(richerror.WithOp(op),
				richerror.WithInnerError(cErr))
		}
		user = newUser
		isNewUser = true
	}

	fmt.Println("user id service", user.ID)
	authenticate := entity.Authenticable{
		ID: user.ID,
	}
	accessToken, caErr := s.authGenerator.CreateAccessToken(authenticate)
	if caErr != nil {
		return nil, richerror.New(richerror.WithOp(op),
			richerror.WithInnerError(caErr))
	}

	refreshToken, crEre := s.authGenerator.CreateRefreshToken(authenticate)
	if crEre != nil {
		return nil, richerror.New(richerror.WithOp(op),
			richerror.WithInnerError(crEre))
	}

	return &userparam.LoginOrRegisterResponse{
		User: user,
		Tokens: userparam.Tokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
		NewUser: isNewUser,
	}, nil

}
