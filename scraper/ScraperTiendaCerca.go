package scraper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julioolivares90/scraper_tienda_cercaV2/models"
)

const (
	access_token = "pk.eyJ1IjoianVsaW9vbGl2YXJlczkwIiwiYSI6ImNrYm52dG1veDEzZjMyb282aXpxZTdvdTkifQ.XceucE2ir4LW8HS2Fhr_7g"
)

func FindLugar(lugar string) models.Lugar {
	urlFind := fmt.Sprintf("https://api.mapbox.com/geocoding/v5/mapbox.places/%s.json?access_token=%s", lugar, access_token)

	resp, err := http.Get(urlFind)
	if err != nil {
		log.Fatal()
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var lugarStruct models.Lugar
	json.Unmarshal(bodyBytes, &lugarStruct)

	return lugarStruct
}
