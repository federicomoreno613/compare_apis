package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var (
	counter uint64 = 0
	mutex   sync.Mutex
)

func main() {
	http.HandleFunc("/date", dateHandler)
	http.HandleFunc("/counter", counterHandler)

	fmt.Println("Listening on port :8080")
	http.ListenAndServe(":8080", nil)
}

// Handler para la ruta /date
func dateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	fullFormat, err := strconv.ParseBool(r.URL.Query().Get("full_format"))
	if err != nil {
		http.Error(w, "Invalid parameter value", http.StatusBadRequest)
		return
	}

	// Enviamos la solicitud al canal para limitar la cantidad de solicitudes simultáneas
	c := make(chan string)
	go handleDateRequest(c, fullFormat)

	// Esperamos a que la goroutine complete la solicitud
	dateStr := <-c

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(dateStr))
}

// Goroutine para manejar la solicitud de fecha
func handleDateRequest(c chan string, fullFormat bool) {
	var dateStr string
	if fullFormat {
		dateStr = time.Now().Format("2006-01-02 15:04:05")
	} else {
		dateStr = time.Now().Format("2006-02-01")
	}
	c <- dateStr
}

func counterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	// Incrementamos el contador con un Mutex para evitar data races
	mutex.Lock()
	counter++
	count := counter
	mutex.Unlock()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.FormatUint(count, 10)))
}
