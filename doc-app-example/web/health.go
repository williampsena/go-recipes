package web

import (
	"fmt"
	"net/http"
)

// Endpoint is responsible for responding to application health
func HealthCheckEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "I'm so healthy 😌!")
}
