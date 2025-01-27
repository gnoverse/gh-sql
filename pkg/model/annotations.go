package model

import "entgo.io/ent/schema"

// NoCopyTemplate is an ent annotation that can be used to mark when a field
// shouldn't be placed in the "copy" template; ie. it should not be possible
// to copy it from an existing structure.
type NoCopyTemplate struct{}

var _ schema.Annotation = NoCopyTemplate{}

func (n NoCopyTemplate) Name() string {
	return "no_copy_template"
}
