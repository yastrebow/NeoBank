package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

type gzipWriter struct {
	http.ResponseWriter
	writer io.Writer
}

func (w gzipWriter) Write(b []byte) (int, error) {
	return w.writer.Write(b)
}

func CompressResponse(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Compress!!!!")

		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		gz, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		defer gz.Close()

		w.Header().Set("Content-Encoding", "gzip")
		next.ServeHTTP(gzipWriter{ResponseWriter: w, writer: gz}, r)
	})
}
