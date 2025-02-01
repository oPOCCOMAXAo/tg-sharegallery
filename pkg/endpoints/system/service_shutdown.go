package system

import "net/http"

func (s *Service) Shutdown(w http.ResponseWriter, _ *http.Request) {
	defer s.cancel(nil)

	w.WriteHeader(http.StatusOK)
}
