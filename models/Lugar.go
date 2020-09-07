package models

type Lugar struct {
	Type     string   `json:"type"`
	Query    []string `json:"query"`
	Features []struct {
		ID         string   `json:"id"`
		Type       string   `json:"type"`
		PlaceType  []string `json:"place_type"`
		Relevance  int      `json:"relevance"`
		Properties struct {
			Wikidata string `json:"wikidata"`
		} `json:"properties,omitempty"`
		Text      string    `json:"text"`
		PlaceName string    `json:"place_name"`
		Bbox      []float64 `json:"bbox,omitempty"`
		Center    []float64 `json:"center"`
		Geometry  struct {
			Type        string    `json:"type"`
			Coordinates []float64 `json:"coordinates"`
		} `json:"geometry"`
		Context []struct {
			ID        string `json:"id"`
			ShortCode string `json:"short_code"`
			Wikidata  string `json:"wikidata"`
			Text      string `json:"text"`
		} `json:"context"`
	} `json:"features"`
	Attribution string `json:"attribution"`
}
