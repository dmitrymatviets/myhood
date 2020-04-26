package api

import (
	"github.com/dmitrymatviets/myhood/api/dto"
	"github.com/dmitrymatviets/myhood/core/contract"
	baseHTTP "github.com/dmitrymatviets/myhood/infrastructure/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CityEndpoint struct {
	*Server
	cityService contract.ICityService
}

func NewCityEndpoint(s *Server, cityService contract.ICityService) *CityEndpoint {
	endpoint := &CityEndpoint{s, cityService}
	endpoint.Server = s

	s.AddRoutes(&baseHTTP.Route{
		Method:      http.MethodPost,
		Path:        "v1/city/list",
		HandleFuncs: []gin.HandlerFunc{endpoint.CityListV1},
	})
	return endpoint
}

func (e *CityEndpoint) CityListV1(ctx *gin.Context) {
	var requestDto dto.CityListRequest
	e.ApiMethod(ctx, &requestDto, func() (interface{}, error) {
		cities, err := e.cityService.GetCities(ctx)
		if err != nil {
			return nil, err
		}
		return dto.CityListResponse{
			Cities: cities,
		}, nil
	})
}
