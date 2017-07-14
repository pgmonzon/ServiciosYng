package core

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Responde en formato JSON
func RespuestaJSON(w http.ResponseWriter, r *http.Request, start time.Time, response []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	log.Printf("%s\t%s\t%s\t%s\t%d\t%d\t%s",
		r.RemoteAddr,
		r.Method,
		r.RequestURI,
		r.Proto,
		code,
		len(response),
		time.Since(start),
	)
	if string(response) != "" {
		w.Write(response)
	}
}

// Respuesta de error en formato JSON
func ErrorJSON(w http.ResponseWriter, r *http.Request, start time.Time, message string, code int) {
	j := map[string]string{"mensaje": message}
	response, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	log.Printf("%s\t%s\t%s\t%s\t%d\t%d\t%s",
		r.RemoteAddr,
		r.Method,
		r.RequestURI,
		r.Proto,
		code,
		len(response),
		time.Since(start),
	)
	w.Write(response)
}
