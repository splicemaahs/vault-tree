package main

import (
	"log"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func mergeMaps(a, b map[string]interface{}) map[string]interface{} {
	out := make(map[string]interface{}, len(a))
	for k, v := range a {
		out[k] = v
	}
	for k, v := range b {
		if v, ok := v.(map[string]interface{}); ok {
			if bv, ok := out[k]; ok {
				if bv, ok := bv.(map[string]interface{}); ok {
					out[k] = mergeMaps(bv, v)
					continue
				}
			}
		}
		out[k] = v
	}
	return out
}

// CoalesceTables merges a source map into a destination map.
//
// dest is considered authoritative.
func CoalesceTables(dst, src map[string]interface{}) map[string]interface{} {
	if dst == nil || src == nil {
		return src
	}
	// Because dest has higher precedence than src, dest values override src
	// values.
	for key, val := range src {
		if istable(val) {
			switch innerdst, ok := dst[key]; {
			case !ok:
				dst[key] = val
			case istable(innerdst):
				CoalesceTables(innerdst.(map[string]interface{}), val.(map[string]interface{}))
			default:
				log.Printf("warning: cannot overwrite table with non table for %s (%v)", key, val)
			}
		} else if dv, ok := dst[key]; ok && istable(dv) {
			log.Printf("warning: destination for %s is a table. Ignoring non-table value %v", key, val)
		} else if !ok { // <- ok is still in scope from preceding conditional.
			dst[key] = val
		}
	}
	return dst
}

// istable is a special-purpose function to see if the present thing matches the definition of a YAML table.
func istable(v interface{}) bool {
	_, ok := v.(map[string]interface{})
	return ok
}

func copyMap(src map[string]interface{}) map[string]interface{} {
	m := make(map[string]interface{}, len(src))
	for k, v := range src {
		m[k] = v
	}
	return m
}
