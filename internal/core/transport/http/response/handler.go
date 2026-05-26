package core_http_response

import (
	"encoding/json"
	"fmt"
	"net/http"

	core_logger "github.com/eternalpaininside/golang-todo-app/internal/core/logger"
	"go.uber.org/zap"
)

type HTTPResponseHandler struct {
	log *core_logger.Logger
	rw  http.ResponseWriter
}

func NewHTTPResponseHandler(
	log *core_logger.Logger,
	rw http.ResponseWriter) *HTTPResponseHandler {
	return &HTTPResponseHandler{
		log: log,
	}
}

func (h *HTTPResponseHandler) PanicResponse(
	p any, message string) {
	statusCode := http.StatusInternalServerError
	err := fmt.Errorf("unexpected panic: %v", p)

	h.log.Error(message, zap.Error(err))
	h.rw.WriteHeader(statusCode)
	response := map[string]string{
		"message": message,
		"error":   err.Error(),
	}

	if err := json.NewEncoder(h.rw).Encode(response); err != nil {
		h.log.Error("write http response", zap.Error(err))
	}
}
