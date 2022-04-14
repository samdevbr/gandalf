package spec

import "strings"

type Php struct {
	Version   string
	Framework *Framework
}

func (p *Php) hasValidVersion() bool {
	parts := strings.Split(p.Version, ".")

	if len(parts) > 3 || len(parts) < 2 {
		return false
	}

	major := parts[0]

	return major == "5" || major == "7"
}
