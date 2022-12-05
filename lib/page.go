package lib

type PageMeta struct {
	Owner      string `json:"owner"`
	Domain     string `json:"domain"`
	LastUpdate int64  `json:"last_update"`
}
