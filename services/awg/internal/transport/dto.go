package transport

import "net/http"

type message struct {
	StatusCode int   `json:"status_code"`
	Error      error `json:"error"`
}
type response struct {
	Message message `json:"message"`
}

type request struct {
	VirtualEndpoint string `json:"virtual_endpoint"`
	FileName        string `json:"file_name"`
}

type createPeer struct {
	Message   message `json:"message"`
	PublicKey string  `json:"public_key"`
	FilePath  string  `json:"file_path"`
}

func newCreatePeer(publicKey, filePath string) createPeer {
	return createPeer{
		Message: message{
			StatusCode: http.StatusCreated,
		},
		PublicKey: publicKey,
		FilePath:  filePath,
	}
}

func newResp(statusCode int, error error) response {
	return response{
		Message: message{
			StatusCode: statusCode,
			Error:      error,
		},
	}
}
