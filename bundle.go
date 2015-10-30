package bundle

import "net/http"

func Bundle(handlers ...http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, h := range handlers {
			bw := responseWriter{ResponseWriter: w}

			h.ServeHTTP(http.ResponseWriter(&bw), r)

			if bw.status != 0 {
				break
			}
		}
	})
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (w *responseWriter) WriteHeader(s int) {
	w.status = s
	w.ResponseWriter.WriteHeader(s)
}
