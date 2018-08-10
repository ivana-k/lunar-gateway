package model

type KVS struct {
	RegionID  string `json:"regionId"`
	ClusterID string `json:"clusterid"`
	Labels    []KV   `json:"labels"`
	Data      []KV   `json:"data"`
}

type KV struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
