package matching_test

import (
	"testing"

	"github.com/mokevnin/crossmatch/internal/matching"
)

func TestCompatible(t *testing.T) {
	base := matching.Channel{Topic: "tech", Language: "ru", AvgReach: 1000}

	tests := []struct {
		name  string
		other matching.Channel
		want  bool
	}{
		{"same reach", matching.Channel{Topic: "tech", Language: "ru", AvgReach: 1000}, true},
		{"double reach", matching.Channel{Topic: "tech", Language: "ru", AvgReach: 2000}, true},
		{"half reach", matching.Channel{Topic: "tech", Language: "ru", AvgReach: 500}, true},
		{"too large", matching.Channel{Topic: "tech", Language: "ru", AvgReach: 2001}, false},
		{"too small", matching.Channel{Topic: "tech", Language: "ru", AvgReach: 499}, false},
		{"other topic", matching.Channel{Topic: "food", Language: "ru", AvgReach: 1000}, false},
		{"other language", matching.Channel{Topic: "tech", Language: "en", AvgReach: 1000}, false},
		{"no reach", matching.Channel{Topic: "tech", Language: "ru", AvgReach: 0}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matching.Compatible(base, tt.other); got != tt.want {
				t.Errorf("Compatible(%+v, %+v) = %v, want %v", base, tt.other, got, tt.want)
			}
			if got := matching.Compatible(tt.other, base); got != tt.want {
				t.Errorf("Compatible is not symmetric for %+v", tt.other)
			}
		})
	}
}
