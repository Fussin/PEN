package plugin

type Plugin interface {
	Name() string
	Run(target string) (string, error)
}

type Manager struct {
	plugins []Plugin
}

func NewManager() *Manager {
	return &Manager{
		plugins: make([]Plugin, 0),
	}
}

func (m *Manager) Register(plugin Plugin) {
	m.plugins = append(m.plugins, plugin)
}

func (m *Manager) Run(target string) []string {
	results := make([]string, 0)
	for _, p := range m.plugins {
		result, err := p.Run(target)
		if err != nil {
			// In a real implementation, we would handle the error.
			continue
		}
		results = append(results, result)
	}
	return results
}
