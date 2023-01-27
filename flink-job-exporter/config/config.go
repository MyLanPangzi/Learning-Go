package config

type AppConfig struct {
	ScrapeInterval string `json:"scrapeInterval"`
	Http           struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"http"`
	Db struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Name     string `json:"name"`
	} `json:"db"`
	Yarn struct {
		Url string `json:"url"`
	} `json:"yarn"`
	Monitor struct {
		Tag string `json:"tag"`
	} `json:"monitor"`
}
