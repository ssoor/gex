package tool

import "path/filepath"

// Tool represents a go package of a tool dependency.
type Tool struct {
	path   string
	Global bool
}

// Name returns an executable name.
func (t Tool) String() string {
	return t.path
}

// Name returns an executable name.
func (t Tool) Name() string {
	return filepath.Base(t.path)
}
