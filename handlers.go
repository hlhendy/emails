package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/emails/internal/messages"
	"github.com/emails/models"
	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	Message messages.Message
}

// NewHandler returns a new Email Handler
func NewHandler(m messages.Message) *Handler {
	return &Handler{
		Message: m,
	}
}

// Send handles requests to send an email message
func (e *Handler) Send(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("%+v", err)
		http.Error(w, "INTERNAL_ERROR", http.StatusInternalServerError)
		return
	}

	var request models.EmailRequest
	if err := json.Unmarshal(body, &request); err != nil {
		log.Printf("%+v", err)
		http.Error(w, "INTERNAL_ERROR", http.StatusInternalServerError)
		return
	}

	if &request == nil {
		http.Error(w, "BAD_REQUEST", http.StatusBadRequest)
		return
	}

	switch {
	case request.RecipientName == "":
		http.Error(w, "INVALID_ARG_TO_NAME_REQUIRED", http.StatusBadRequest)
		return
	case request.RecipientEmail == "":
		http.Error(w, "INVALID_ARG_TO_REQUIRED", http.StatusBadRequest)
		return
	case request.SenderName == "":
		http.Error(w, "INVALID_ARG_FROM_NAME_REQUIRED", http.StatusBadRequest)
		return
	case request.SenderEmail == "":
		http.Error(w, "INVALID_ARG_FROM_REQUIRED", http.StatusBadRequest)
		return
	case request.HTMLBody == "":
		http.Error(w, "INVALID_ARG_BODY_REQUIRED", http.StatusBadRequest)
		return
	case request.Subject == "":
		http.Error(w, "INVALID_ARG_SUBJECT_REQUIRED", http.StatusBadRequest)
		return
	}

	err = request.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = e.Message.Send(&request)
	if err != nil {
		log.Printf("%+v", err)
		http.Error(w, "INTERNAL_ERROR", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("200 OK"))
}
