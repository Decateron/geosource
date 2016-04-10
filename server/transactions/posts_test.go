package transactions

import (
	"testing"

	"github.com/joshheinrichs/geosource/server/types"
)

func BenchmarkGetPosts(b *testing.B) {
	// Testing speed of queries
	limit := 20
	offset := 0
	postQueryParams := PostQueryParams{
		Limit:     &limit,
		Offset:    &offset,
		Flags:     nil,
		TimeRange: nil,
		LocationRange: &LocationRange{
			Min: types.Location{
				Latitude:  45,
				Longitude: 45,
			},
			Max: types.Location{
				Latitude:  46,
				Longitude: 46,
			},
		},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := GetPosts("admin", &postQueryParams)
		if err != nil {
			b.Error(err)
		}
	}
}
