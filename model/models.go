package model

type KV struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type MutateReq struct {
	Version string              `json:"version"`
	Kind    string              `json:"kind"`
	Region  map[string]Region   `json:"region"`
	Data    map[string][]string `json:"data"`
}

type Region struct {
	Cluster  []string
	Selector map[string]string
}
