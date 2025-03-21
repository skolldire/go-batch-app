package ping

import (
	"net/http"
)

type service struct {
}

var _ Service = (*service)(nil)

func NewService() *service {
	return &service{}
}

// Apply godoc
// @Summary      Validate that the application is working
// @Description  get string by string
// @Tags         pong
// @Accept       json
// @Produce      json
// @Success      200  {object}  string
// @Router       /ping [get]
func (s *service) Apply() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("pong"))
		if err != nil {
			return
		}
	}
}
