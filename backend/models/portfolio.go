package models

type Portfolio struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Biography string `json:"biography"`
	Films     string `json:"films"`
	Awards    string `json:"awards"`
}
