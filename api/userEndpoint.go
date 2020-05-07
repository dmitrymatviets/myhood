package api

import (
	"github.com/dmitrymatviets/myhood/api/dto"
	"github.com/dmitrymatviets/myhood/core/contract"
	baseHTTP "github.com/dmitrymatviets/myhood/infrastructure/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserEndpoint struct {
	*Server
	contract.IUserService
}

func NewUserEndpoint(s *Server, userService contract.IUserService) *UserEndpoint {
	endpoint := &UserEndpoint{s, userService}
	endpoint.Server = s

	s.AddRoutes(&baseHTTP.Route{
		Method:      http.MethodPost,
		Path:        "v1/user/getUser",
		HandleFuncs: []gin.HandlerFunc{endpoint.GetUserV1},
	})

	s.AddRoutes(&baseHTTP.Route{
		Method:      http.MethodPost,
		Path:        "v1/user/getFriends",
		HandleFuncs: []gin.HandlerFunc{endpoint.GetFriendsV1},
	})

	s.AddRoutes(&baseHTTP.Route{
		Method:      http.MethodPost,
		Path:        "v1/user/addFriend",
		HandleFuncs: []gin.HandlerFunc{endpoint.AddFriendV1},
	})

	s.AddRoutes(&baseHTTP.Route{
		Method:      http.MethodPost,
		Path:        "v1/user/removeFriend",
		HandleFuncs: []gin.HandlerFunc{endpoint.RemoveFriendV1},
	})

	s.AddRoutes(&baseHTTP.Route{
		Method:      http.MethodPost,
		Path:        "v1/user/search",
		HandleFuncs: []gin.HandlerFunc{endpoint.SearchV1},
	})

	/*
		s.AddRoutes(&baseHTTP.Route{
			Method:      http.MethodPost,
			Path:        "v1/user/saveUser",
			HandleFuncs: []gin.HandlerFunc{endpoint.SaveUserV1},
		})*/

	s.AddRoutes(&baseHTTP.Route{
		Method:      http.MethodPost,
		Path:        "v1/user/getRecommendations",
		HandleFuncs: []gin.HandlerFunc{endpoint.GetRecommendationsV1},
	})

	return endpoint
}

func (ue *UserEndpoint) GetUserV1(ctx *gin.Context) {
	var requestDto dto.GetUserRequest
	ue.ApiMethod(ctx, &requestDto, func() (interface{}, error) {
		user, err := ue.GetById(ctx, requestDto.Session, requestDto.UserId)
		if err != nil {
			return nil, err
		}
		return dto.GetUserResponse{
			User: user,
		}, nil
	})
}

func (ue *UserEndpoint) GetFriendsV1(ctx *gin.Context) {
	var requestDto dto.GetFriendsRequest
	ue.ApiMethod(ctx, &requestDto, func() (interface{}, error) {
		friends, err := ue.GetFriends(ctx, requestDto.Session, requestDto.UserId)
		if err != nil {
			return nil, err
		}
		return dto.GetFriendsResponse{
			Friends: friends,
		}, nil
	})
}

func (ue *UserEndpoint) AddFriendV1(ctx *gin.Context) {
	var requestDto dto.AddFriendRequest
	ue.ApiMethod(ctx, &requestDto, func() (interface{}, error) {
		err := ue.AddFriend(ctx, requestDto.Session, requestDto.FriendId)
		if err != nil {
			return nil, err
		}
		return dto.GetFriendsResponse{}, nil
	})
}

func (ue *UserEndpoint) RemoveFriendV1(ctx *gin.Context) {
	var requestDto dto.RemoveFriendRequest
	ue.ApiMethod(ctx, &requestDto, func() (interface{}, error) {
		err := ue.RemoveFriend(ctx, requestDto.Session, requestDto.FriendId)
		if err != nil {
			return nil, err
		}
		return dto.RemoveFriendResponse{}, nil
	})
}

func (ue *UserEndpoint) SaveUserV1(ctx *gin.Context) {
	var requestDto dto.SaveUserRequest
	ue.ApiMethod(ctx, &requestDto, func() (interface{}, error) {
		user, err := ue.SaveUser(ctx, requestDto.Session, requestDto.User)
		if err != nil {
			return nil, err
		}
		return dto.SaveUserResponse{
			User: user,
		}, nil
	})
}

func (ue *UserEndpoint) GetRecommendationsV1(ctx *gin.Context) {
	var requestDto dto.GetRecommendationsRequest
	ue.ApiMethod(ctx, &requestDto, func() (interface{}, error) {
		recs, err := ue.GetRecommendations(ctx, requestDto.Session)
		if err != nil {
			return nil, err
		}
		return dto.GetRecommendationsResponse{
			Recommendations: recs,
		}, nil
	})
}

func (ue *UserEndpoint) SearchV1(ctx *gin.Context) {
	var requestDto dto.SearchRequest
	ue.ApiMethod(ctx, &requestDto, func() (interface{}, error) {
		users, err := ue.Search(ctx, requestDto.Session, requestDto.SearchDto)
		if err != nil {
			return nil, err
		}
		return dto.SearchResponse{
			Users: users,
		}, nil
	})
}
