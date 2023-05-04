package hanlders

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	mailers_adapaters "github.com/karlozz157/stori/app/mailers/adapters"
	"github.com/karlozz157/stori/app/models"
	"github.com/karlozz157/stori/app/services"
	services_transactions_adapters "github.com/karlozz157/stori/app/services/transactions/adapters"
)

type transactionHandler struct {
	summaryService *services.SummaryService
}

func NewTransactionHandler() *transactionHandler {
	mailerService := mailers_adapaters.NewStmpMailer()
	transactionService := services_transactions_adapters.NewTransactionStorageService()
	transactionService.Init()
	summaryService := services.NewSummaryService(mailerService, transactionService)

	return &transactionHandler{summaryService: summaryService}
}

func (h *transactionHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/", h.helloWorld).Methods(http.MethodGet)
	r.HandleFunc("/summary", h.createSummary).Methods(http.MethodPost)
	r.HandleFunc("/summary", h.getSummary).Methods(http.MethodGet)
	r.HandleFunc("/summary/send", h.createAndSendSummary).Methods(http.MethodPost)
}

func (h *transactionHandler) helloWorld(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Response{
		Message: "Hello World!",
	})
}

func (h *transactionHandler) createSummary(w http.ResponseWriter, r *http.Request) {
	err := h.summaryService.CreateSummary()
	h.handleResponse(w, err)
}

func (h *transactionHandler) createAndSendSummary(w http.ResponseWriter, r *http.Request) {
	err := h.summaryService.CreateAndSendSummary()
	h.handleResponse(w, err)
}

func (h *transactionHandler) getSummary(w http.ResponseWriter, r *http.Request) {
	summary := h.summaryService.GetSummary()
	json.NewEncoder(w).Encode(summary)
}

func (h *transactionHandler) handleResponse(w http.ResponseWriter, err error) {
	statusCode := http.StatusOK
	response := models.Response{
		Message: "ok",
	}

	if err != nil {
		statusCode = http.StatusInternalServerError
		response.Message = err.Error()
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
