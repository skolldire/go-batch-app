package ping

import "net/http"

type Service interface {
	Apply() http.HandlerFunc
}
