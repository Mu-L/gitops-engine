package json

// https://github.com/ksonnet/ksonnet/blob/master/pkg/kubecfg/diff.go
func removeFields(config, live any) any {
	switch c := config.(type) {
	case map[string]any:
		l, ok := live.(map[string]any)
		if ok {
			return RemoveMapFields(c, l)
		}
		return live
	case []any:
		l, ok := live.([]any)
		if ok {
			return RemoveListFields(c, l)
		}
		return live
	default:
		return live
	}
}

// RemoveMapFields remove all non-existent fields in the live that don't exist in the config
func RemoveMapFields(config, live map[string]any) map[string]any {
	result := map[string]any{}
	for k, v1 := range config {
		v2, ok := live[k]
		if !ok {
			continue
		}
		if v2 != nil {
			v2 = removeFields(v1, v2)
		}
		result[k] = v2
	}
	return result
}

func RemoveListFields(config, live []any) []any {
	// If live is longer than config, then the extra elements at the end of the
	// list will be returned as-is so they appear in the diff.
	result := make([]any, 0, len(live))
	for i, v2 := range live {
		if len(config) > i {
			if v2 != nil {
				v2 = removeFields(config[i], v2)
			}
			result = append(result, v2)
		} else {
			result = append(result, v2)
		}
	}
	return result
}
