package handlers

import (
	"encoding/json"
	"net/http"
	"refrigerator/business"
)

func (d *Deps) GetFood(w http.ResponseWriter, r *http.Request) {
	c, err := d.DB.Conn(r.Context())
	if err != nil {
		// Write your own error response here!
		// And also on everything else where the if err != nil
		// returns nothing.
		return
	}
	defer c.Close()

	id := r.URL.Query().Get("id")
	if id != "" {
		// Handle request if id was provided.
		// Write a new function on /business/food.go
		// that searches for a food by id.
		return
	}

	foods, err := business.GetAllFoods(c, r.Context())
	if err != nil {
		// Write your own error handling logic!
		return
	}

	data, err := json.Marshal(&foods)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "application/json")
	w.Write(data)
}
