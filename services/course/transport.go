package course

import (
    "context"
    "encoding/json"
    "net/http"

    httptransport "github.com/go-kit/kit/transport/http"
)

func MakeHTTPHandler(endpoints Endpoints) http.Handler {
    mux := http.NewServeMux()
    mux.Handle("/courses", httptransport.NewServer(
        endpoints.GetAllCoursesEndpoint,
        decodeGetAllCoursesRequest,
        encodeResponse,
    ))
    return mux
}

func decodeGetAllCoursesRequest(_ context.Context, r *http.Request) (interface{}, error) {
    return nil, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
    return json.NewEncoder(w).Encode(response)
}
