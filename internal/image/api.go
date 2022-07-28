package image

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(mux *http.ServeMux, service Service) {
	res := resource{service}

	mux.HandleFunc("/image-of-the-day", res.get)
	mux.HandleFunc("/images", res.list)
}

type resource struct {
	service Service
}

func (r resource) get(rw http.ResponseWriter, req *http.Request) {
	dateParam, ok := req.URL.Query()["date"]
	if !ok || len(dateParam[0]) < 1 {
		replyError(rw, http.StatusBadRequest, "invalid request was sent")
		return
	}

	date, err := time.Parse("2006-01-02", dateParam[0])
	if err != nil {
		replyError(rw, http.StatusBadRequest, "invalid request format was sent")
		return
	}

	img, err := r.service.GetByDate(req.Context(), date)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			replyError(rw, http.StatusNotFound, "image for the given date not found")
		default:
			replyError(rw, http.StatusInternalServerError, "error occured while retreiving the image")
		}
		return
	}

	replyJson(rw, img)
}

func (r resource) list(rw http.ResponseWriter, req *http.Request) {
	images, err := r.service.List(req.Context())
	if err != nil {
		replyError(rw, http.StatusInternalServerError, "error occured while retreiving the images")
		return
	}

	replyJson(rw, images)
}

func replyError(rw http.ResponseWriter, status int, format string, args ...interface{}) {
	http.Error(rw, fmt.Sprintf(format, args...), status)
}

func replyJson(rw http.ResponseWriter, model interface{}) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(rw).Encode(model); err != nil {
		log.Print(err.Error())
	}
}
