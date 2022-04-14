package spec

import "strings"

type php struct {
	Version   string
	Framework *framework
}

func (p *php) hasValidVersion() bool {
	parts := strings.Split(p.Version, ".")

	if len(parts) > 3 || len(parts) < 2 {
		return false
	}

	major := parts[0]

	return major == "5" || major == "7"
}
