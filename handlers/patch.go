package handlers

import "net/http"

func (d *Deps) UpdateFood(w http.ResponseWriter, r *http.Request) {
	foodID := r.URL.Query().Get("id")
	if foodID == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("content-type", "application/json")
		w.Write([]byte(`{"error": "missing id"}`))
		return
	}

	// Continue the logic of updating a food here.
	// Returns a JSON response that consist of an array with
	// two objects:
	// 1. The previous state of the food item
	// 2. The updated state of the food item
}
