package authorcontroller

import (
	"encoding/json"
	"net/http"
	"restapi/config"
	"restapi/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var author []models.Author
	
	if err := config.DB.Find(&author).Error; err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	res, _ := json.Marshal(author)
	w.Write(res)

}
