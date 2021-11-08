package handlers

import "net/http"

func (d *Deps) DeleteFood(w http.ResponseWriter, r *http.Request) {
	foodID := r.URL.Query().Get("id")
	if foodID == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("content-type", "application/json")
		w.Write([]byte(`{"error": "missing id"}`))
		return
	}

	// Continue the logic of deleting a food here
	// Then send the response back to the client
	// of the food item that has been deleted.
}
