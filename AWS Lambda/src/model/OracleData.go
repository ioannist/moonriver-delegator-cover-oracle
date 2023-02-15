package model

type DelegationsData struct {
	OwnerAccount string `json:"ownerAccount"`
	Amount       uint   `json:"amount"`
}

type CollatorData struct {
	CollatorAccount      string            `json:"collatorAccount"`
	Points               string            `json:"points"`
	Active               bool              `json:"active"`
	Bond                 string            `json:"bond"`
	DelegationsTotal     string            `json:"delegationsTotal"`
	TopActiveDelegations []DelegationsData `json:"topActiveDelegations"`
}

type OracleData struct {
	TotalStaked   string         `json:"totalStaked"`
	TotalSelected string         `json:"totalSelected"`
	OrbitersCount string         `json:"orbitersCount"`
	Round         string         `json:"round"`
	BlockHash     string         `json:"blockHash"`
	BlockNumber   string         `json:"blockNumber"`
	Awarded       string         `json:"awarded"`
	Collators     []CollatorData `json:"collators"`
	Finalize      bool           `json:"finalize"`
}
