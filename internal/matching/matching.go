// Package matching decides whether two channels fit each other for a mutual promotion.
package matching

// Channel is what the matching needs to know about a channel.
type Channel struct {
	Topic    string
	Language string
	AvgReach int
}

// MaxReachRatio is how far apart two channels' average reach may be: a channel of 1000
// is offered channels from 500 to 2000.
const MaxReachRatio = 2.0

// Compatible reports whether a and b may be offered to each other: the same topic and
// language, and an average reach within MaxReachRatio of each other.
func Compatible(a, b Channel) bool {
	if a.Topic != b.Topic || a.Language != b.Language {
		return false
	}
	if a.AvgReach <= 0 || b.AvgReach <= 0 {
		return false
	}
	low, high := a.AvgReach, b.AvgReach
	if low > high {
		low, high = high, low
	}
	return float64(high) <= float64(low)*MaxReachRatio
}
