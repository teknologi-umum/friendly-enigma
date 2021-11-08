package handlers

import (
	"encoding/json"
	"net/http"
	"refrigerator/business"
)

func (d *Deps) AddFood(w http.ResponseWriter, r *http.Request) {
	c, err := d.DB.Conn(r.Context())
	if err != nil {
		// Write your own error response here!
		// And also on everything else where the if err != nil
		// returns nothing.
		return
	}
	defer c.Close()

	// Read request body as JSON
	var food business.Food
	err = json.NewDecoder(r.Body).Decode(&food)
	if err != nil {
		// Write your own error response here!
		return
	}

	// Add food to database. Try to implement your own
	// SQL query that might be better from this one.
	//
	// Because I'm sure this one wouldn't work haha!
	_, err = business.AddNewFood(c, r.Context(), food)
	if err != nil {
		// Write your own error response here!
		return
	}

	// Write your response!
}
