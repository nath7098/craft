package templating

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/imdario/mergo"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
)

// FuncMap returns a minimal template.FuncMap.
//
// Available functions are fromYaml, map, prefix, suffix and toYaml.
func FuncMap() template.FuncMap {
	return template.FuncMap{
		"map":    mergeMaps,
		"toYaml": toYAML,
	}
}

func mergeMaps(dst map[string]any, src ...any) map[string]any {
	for i, in := range src {
		var cast map[string]any
		if err := mapstructure.Decode(in, &cast); err != nil {
			dst[fmt.Sprint(i, "_decode_error")] = err.Error()
			continue
		}
		if err := mergo.Merge(&dst, cast); err != nil {
			dst[fmt.Sprint(i, "_merge_error")] = err.Error()
			continue
		}
	}
	return dst
}

// toYAML takes an interface, marshals it to yaml, and returns a string. It will
// always return a string, even on marshal error (empty string).
//
// This is designed to be called from a template.
// Copy of https://github.com/helm/helm/blob/main/pkg/engine/funcs.go#L82.
func toYAML(v any) string {
	data, err := yaml.Marshal(v)
	if err != nil {
		// Swallow errors inside of a template.
		return ""
	}
	return strings.TrimSuffix(string(data), "\n")
}
