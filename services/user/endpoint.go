package user

import (
    "context"

    "github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
    GetAllUsersEndpoint endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
    return Endpoints{
        GetAllUsersEndpoint: MakeGetAllUsersEndpoint(s),
    }
}

func MakeGetAllUsersEndpoint(s Service) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        return s.GetAllUsers(ctx)
    }
}
