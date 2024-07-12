package course

import (
    "context"

    "github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
    GetAllCoursesEndpoint endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
    return Endpoints{
        GetAllCoursesEndpoint: MakeGetAllCoursesEndpoint(s),
    }
}

func MakeGetAllCoursesEndpoint(s Service) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        return s.GetAllCourses(ctx)
    }
}
