package user

import (
    "context"
    "encoding/json"
    "net/http"

    httptransport "github.com/go-kit/kit/transport/http"
)

func MakeHTTPHandler(endpoints Endpoints) http.Handler {
    mux := http.NewServeMux()
    mux.Handle("/users", httptransport.NewServer(
        endpoints.GetAllUsersEndpoint,
        decodeGetAllUsersRequest,
        encodeResponse,
    ))
    return mux
}

func decodeGetAllUsersRequest(_ context.Context, r *http.Request) (interface{}, error) {
    return nil, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
    return json.NewEncoder(w).Encode(response)
}
