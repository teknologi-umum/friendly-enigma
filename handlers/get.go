package handlers

import (
	"encoding/json"
	"net/http"
	"refrigerator/business"

	"github.com/georgysavva/scany/sqlscan"
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

	ctx := r.Context()
	member, ok := ctx.Value("user").(business.Member)
	if !ok {
		return
	}

	rows, err := c.QueryContext(
		r.Context(),
		"SELECT * FROM refrigerator WHERE id = ?",
		member.ID,
	)
	if err != nil {
		return
	}
	defer rows.Close()

	var foods []business.Food
	err = sqlscan.ScanAll(&foods, rows)
	if err != nil {
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
