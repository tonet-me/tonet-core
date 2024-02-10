package userparam

import "github.com/tonet-me/tonet-core/entity"

type LoginOrRegisterRequest struct {
	Token string
}
type LoginOrRegisterResponse struct {
	User   entity.User
	Tokens Tokens
}
