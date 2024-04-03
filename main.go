package main

import (
	"context"
	"net/http"

	"github.com/google/uuid"
//	"github.com/traefik/traefik/v3/pkg/plugins"
)

// Config estrutura para a configuração do plugin.
type Config struct{}

// CreateConfig cria e inicializa a configuração do plugin.
func CreateConfig() *Config {
	return &Config{}
}

type addIDs struct {
	next http.Handler
	name string
}

// New cria e retorna uma nova instância do seu plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &addIDs{
		next: next,
		name: name,
	}, nil
}

func (a *addIDs) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	traceID := uuid.New().String() // Gera um TraceID único
	spanID := uuid.New().String()  // Gera um SpanID único

	// Adiciona o TraceID e o SpanID ao cabeçalho da requisição no formato especificado
	req.Header.Add("X-Request-ID", traceID+"-"+spanID)

	// Mantém os headers originais para TraceID e SpanID para fins de debug ou tracing posterior
	req.Header.Add("X-Trace-ID", traceID)
	req.Header.Add("X-Span-ID", spanID)

	// Chama o próximo handler na cadeia
	a.next.ServeHTTP(rw, req)
}

