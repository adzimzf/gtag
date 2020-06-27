package formatter

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_sortByCount(t *testing.T) {
	tm := map[string]tagMap{
		"json": {
			MaxLength: 5,
			Count:     12,
		},
		"db": {
			MaxLength: 5,
			Count:     4,
		},
		"gorm": {
			MaxLength: 5,
			Count:     4,
		},
	}
	expectedKeys := []string{"json", "db", "gorm"}
	for i := 0; i < 20; i++ {
		t.Run(fmt.Sprintf("test_i_%d", i), func(t *testing.T) {
			actual := sortTagMap(tm)
			if b := reflect.DeepEqual(actual, expectedKeys); !b {
				t.Fatalf("expected: %v\nactual: %v", expectedKeys, actual)
			}
		})
	}

	tm = map[string]tagMap{
		"gorm": {
			MaxLength: 7,
			Count:     10,
		},
		"json": {
			MaxLength: 7,
			Count:     13,
		},
	}
	expectedKeys = []string{"json", "gorm"}
	for i := 0; i < 20; i++ {
		t.Run(fmt.Sprintf("test_j_%d", i), func(t *testing.T) {
			actual := sortTagMap(tm)
			if b := reflect.DeepEqual(actual, expectedKeys); !b {
				t.Fatalf("expected: %v\nactual: %v", expectedKeys, actual)
			}
		})
	}

}
