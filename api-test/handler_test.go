package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestDateHandler(t *testing.T) {
	req, err := http.NewRequest("POST", "/date?full_format=true", nil)
	if err != nil {
		t.Fatal(err)
	}
	// creamos un ResponseRecorder para obtener la respuesta
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(dateHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Date handler returned wrong status code: got %v want %v",
			status, http.StatusOK)

	}
	expectedDate := time.Now().Format("2006-01-02 15:04:05")
	if rr.Body.String() != expectedDate {
		t.Errorf("Date handler returned unexpected body: got %v want %v",
			rr.Body.String(), expectedDate)
	}
	req, err = http.NewRequest("POST", "/date", nil)
	if err != nil {
		t.Fatal(err)

	}
	req, err = http.NewRequest("GET", "/date?full_format=true", nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCounterHandler(t *testing.T) {
	counter = 0
	// creamos un WaitGroup para esperar a que finalicen las solicitudes concurrentes
	var wg sync.WaitGroup

	concurrentRequests := 10

	for i := 0; i < concurrentRequests; i++ {
		wg.Add(1)
		go func() {
			req, err := http.NewRequest("GET", "/counter", nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(counterHandler)
			handler.ServeHTTP(rr, req)
			// comprobamos el código de estado
			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	if counter != uint64(concurrentRequests) {
		t.Errorf("handler returned unexpected counter value: got %v want %v",
			counter, concurrentRequests)
	}

	expectedCounter := strconv.FormatUint(counter, 10)
	if counter != uint64(concurrentRequests) {
		t.Errorf("handler returned unexpected counter value: got %v want %v",
			counter, expectedCounter)
	}
}
