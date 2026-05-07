package transport

import "net/http"

type message struct {
	StatusCode int    `json:"status_code"`
	Error      string `json:"error"`
}
type response struct {
	Message message `json:"message"`
}

type request struct {
	VirtualEndpoint string `json:"virtual_endpoint"`
	ID              int64  `json:"id"`
}

type createPeerResponse struct {
	Message   message `json:"message"`
	PublicKey string  `json:"public_key"`
}

func newCreatePeer(publicKey string) createPeerResponse {
	return createPeerResponse{
		Message: message{
			StatusCode: http.StatusCreated,
		},
		PublicKey: publicKey,
	}
}

func newResp(statusCode int, err error) response {
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}
	return response{
		Message: message{
			StatusCode: statusCode,
			Error:      errMsg,
		},
	}
}
