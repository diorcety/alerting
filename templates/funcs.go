package templates

import (
	"github.com/Masterminds/sprig/v3"
	"github.com/prometheus/alertmanager/template"
)

func MergeMaps[M ~map[K]V, K comparable, V any](src ...M) M {
	merged := make(M)
	for _, m := range src {
		for k, v := range m {
			merged[k] = v
		}
	}
	return merged
}

var ExtendedFuncs = template.FuncMap(sprig.FuncMap())

var DefaultFuncs = MergeMaps(template.DefaultFuncs, ExtendedFuncs)
