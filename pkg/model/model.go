// Package model defines some types used within the ent schemas.
package model

// License is a License on the GitHub API. In the schema, this is
// nullable-license-simple.
type License struct {
	Key    string `json:"key"`
	Name   string `json:"name"`
	URL    string `json:"url"`
	NodeID string `json:"node_id"`
	SpdxID string `json:"spdx_id"`
}
