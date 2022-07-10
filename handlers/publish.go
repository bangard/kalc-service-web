package handlers

import (
	"dobledcloud.com/consumers/models"
	"dobledcloud.com/consumers/repository"
	"dobledcloud.com/consumers/server"
	"encoding/json"
	"net/http"
)

func PublishesHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		publishes, err := repository.GetFilesPublishedByEmission(r.Context(), 3)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		pubData := []models.PublishJSON{}
		for i := 0; i < len(publishes); i++ {
			jdata := models.PublishDate{}
			json.Unmarshal([]byte(publishes[i].Date), &jdata)
			pubjson := models.PublishJSON{jdata.Days, publishes[i].DateRange, publishes[i].Hours, publishes[i].Md5, publishes[i].Position, publishes[i].TimeToAir, publishes[i].Url}
			pubData = append(pubData, pubjson)
		}
		data, _ := json.Marshal(models.PublishContent{pubData})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}
