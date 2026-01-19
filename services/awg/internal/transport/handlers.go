package transport

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// awg is an interface for interacting with the AWG service.
type awg interface {
	AddPeer(fileName, virtualEndpoint string) (string, error)
	DeletePeer(peerPublicKeyStr string) error
}

// handlers is a struct that contains the AWG service and handles HTTP requests.
type handlers struct {
	awg awg
}

// DeletePeer handles the DELETE request to delete a peer.
// use endpoint with publicKey query parameter
func (h *handlers) DeletePeer(w http.ResponseWriter, r *http.Request) {

	// Read publicKey in URL
	publicKey := r.URL.Query().Get("publicKey")
	// check if publicKey query parameter is empty
	if publicKey == "" {
		w.WriteHeader(http.StatusBadRequest)
		resp := newResp(http.StatusBadRequest, fmt.Errorf("public key is required"))
		json.NewEncoder(w).Encode(resp)

		return
	}

	// awg delete peer and get process status
	if err := h.awg.DeletePeer(publicKey); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := newResp(http.StatusInternalServerError, fmt.Errorf("failed to delete peer: %w", err))
		json.NewEncoder(w).Encode(resp)

		return
	}

	w.WriteHeader(http.StatusOK)
	resp := newResp(http.StatusOK, nil)
	json.NewEncoder(w).Encode(resp)

}

func (h *handlers) AddPeer(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	var req request
	json.NewDecoder(r.Body).Decode(&req)

	// check if file name and virtual endpoint are empty
	if req.FileName == "" || req.VirtualEndpoint == "" {
		w.WriteHeader(http.StatusBadRequest)
		resp := newResp(http.StatusBadRequest, fmt.Errorf("file name and virtual endpoint are required"))
		json.NewEncoder(w).Encode(resp)

		return
	}

	publicKey, err := h.awg.AddPeer(req.FileName, req.VirtualEndpoint)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := newResp(http.StatusInternalServerError, fmt.Errorf("failed to add peer: %w", err))
		json.NewEncoder(w).Encode(resp)

		return
	}

	resp := newCreatePeer(publicKey, req.FileName+".conf")
	json.NewEncoder(w).Encode(resp)

}
