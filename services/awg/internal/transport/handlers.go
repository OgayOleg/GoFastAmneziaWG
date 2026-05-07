package transport

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/gorilla/mux"
)

// awg is an interface for interacting with the AWG service.
type awg interface {
	AddPeer(fileName, virtualEndpoint string) (string, string, error)
	DeletePeer(peerPublicKeyStr string) error
}

// handlers is a struct that contains the AWG service and handles HTTP requests.
type handlers struct {
	awg         awg
	storagePath string
}

// DeletePeer handles the DELETE request to delete a peer.
// use endpoint with publicKey VAR parameter
func (h *handlers) DeletePeer(w http.ResponseWriter, r *http.Request) {

	// Read publicKey in URL
	vars := mux.Vars(r)
	publicKey := vars["publicKey"]

	// awg delete peer and get process status
	if err := h.awg.DeletePeer(publicKey); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := newResp(http.StatusInternalServerError, fmt.Errorf("failed to delete peer: %w", err))

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			fmt.Printf("failed to encode response: %v", err)
		}

		return
	}

	w.WriteHeader(http.StatusOK)
	resp := newResp(http.StatusOK, nil)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		fmt.Printf("failed to encode response: %v", err)
	}

}

// AddPeer handles the POST request to add a peer.
// use json body with publicKey, id parameters
func (h *handlers) AddPeer(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	var req request

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Printf("failed to decode request: %v", err)
	}

	// check if file name and virtual endpoint are empty
	if req.ID == 0 || req.VirtualEndpoint == "" {
		w.WriteHeader(http.StatusBadRequest)
		resp := newResp(http.StatusBadRequest, fmt.Errorf("file name and virtual endpoint are required"))

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			fmt.Printf("failed to encode response: %v", err)
		}

		return
	}

	fileName := strconv.FormatInt(req.ID, 10)

	_, publicKey, err := h.awg.AddPeer(fileName, req.VirtualEndpoint)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := newResp(http.StatusInternalServerError, fmt.Errorf("failed to add peer: %w", err))

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			fmt.Printf("failed to encode response: %v", err)
		}

		return
	}

	w.WriteHeader(http.StatusCreated)
	resp := newCreatePeer(publicKey)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		fmt.Printf("failed to encode response: %v", err)
	}

}

// handler for sending configuration file by VAR=id
func (h *handlers) SendConfFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	filePath := path.Join(h.storagePath, id+".conf")

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		w.WriteHeader(http.StatusNotFound)
		resp := newResp(http.StatusNotFound, fmt.Errorf("file not exist"))

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			fmt.Printf("failed to encode response: %v", err)
		}

		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := newResp(http.StatusInternalServerError, fmt.Errorf("failed to check file existence: %w", err))

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			fmt.Printf("failed to encode response: %v", err)
		}

		return
	}

	http.ServeFile(w, r, filePath)

}
