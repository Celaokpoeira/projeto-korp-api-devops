package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Resposta struct {
	Nome    string `json:"nome"`
	Horario string `json:"horario"`
}

var (
	requestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Volume total de requisições recebidas.",
		},
		[]string{"path"},
	)
	serviceUp = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "service_up",
			Help: "Status do serviço (1 = Up, 0 = Down).",
		},
	)
)

func init() {
	prometheus.MustRegister(requestsTotal)
	prometheus.MustRegister(serviceUp)
	serviceUp.Set(1)
}

func projetoKorpHandler(w http.ResponseWriter, r *http.Request) {
	requestsTotal.WithLabelValues(r.URL.Path).Inc()

	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	horarioAtual := time.Now().UTC().Format(time.RFC3339)
	resposta := Resposta{Nome: "Projeto Korp", Horario: horarioAtual}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resposta)
}

func main() {
	http.HandleFunc("/projeto-korp", projetoKorpHandler)
	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Servidor rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
