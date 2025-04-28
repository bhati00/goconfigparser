package models

type Config struct {
	Sections map[string]map[string]string `json:"sections"`
}
