package gqlschema

type ClusterAsset struct {
	Name       string      `json:"name"`
	Type       string      `json:"type"`
	Status     AssetStatus `json:"status"`
	Metadata   JSON        `json:"metadata"`
	Parameters JSON        `json:"parameters"`
}
