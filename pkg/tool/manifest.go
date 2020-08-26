package tool

import (
	"sort"

	"github.com/ssoor/gex/pkg/manager"
)

// Manifest contains tool list
type Manifest struct {
	toolMap     map[string]Tool
	managerType manager.Type
}

// NewManifest creates a new Manifest instance.
func NewManifest(tools []Tool, mType manager.Type) *Manifest {
	toolMap := make(map[string]Tool, len(tools))
	for _, t := range tools {
		toolMap[t.Name()] = t
	}
	return &Manifest{toolMap: toolMap, managerType: mType}
}

func (m *Manifest) ManagerType() manager.Type { return m.managerType }

// AddTool adds a new tool to the manifest.
func (m *Manifest) AddTool(tool Tool) {
	m.toolMap[tool.Name()] = tool
}

// FindTool returns a tool by a name.
func (m *Manifest) FindTool(name string) (t Tool, ok bool) {
	t, ok = m.toolMap[name]
	return
}

// Tools returns a tool list.
func (m *Manifest) Tools() []Tool {
	n := len(m.toolMap)
	s := make([]Tool, 0, n)
	for _, t := range m.toolMap {
		s = append(s, t)
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i].String() > s[j].String()
	})
	ts := make([]Tool, n, n)
	for i, t := range s {
		ts[i] = t
	}
	return ts
}

func (m *Manifest) addTool(t Tool) {
	m.toolMap[t.Name()] = t
}
