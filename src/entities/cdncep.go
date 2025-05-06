package entities

type CdnAPICep struct {
	Status   int    `json:"status"`
	Message  string `json:"message"`
	Code     string `json:"code"`
	State    string `json:"state"`
	City     string `json:"city"`
	District string `json:"district"`
	Address  string `json:"address"`
}
