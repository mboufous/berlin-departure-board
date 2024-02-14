package hafas

import (
	"fmt"
	"math"
	"testing"
	"time"
)

// Mock current time for testing
var now = time.Date(2024, 02, 13, 22, 0, 0, 0, time.UTC) // Example: A weekday at 3 PM

func TestDepartureService_getAdaptiveTTL(t *testing.T) {
	s := &DepartureService{} // Assuming DepartureService struct exists

	tests := []struct {
		name          string
		departureTime time.Time
		want          time.Duration
	}{
		{
			name:          "Departure in the past",
			departureTime: now.Add(-1 * time.Hour),
			want:          time.Duration(math.Round(1.5 * 10 * float64(time.Second))),
		},
		{
			name:          "Departure in the next 1 second",
			departureTime: now.Add(time.Second),
			want:          time.Duration(math.Round(2 * float64(time.Second))),
		},
		{
			name:          "Departure in the next 30 second",
			departureTime: now.Add(30 * time.Second),
			want:          time.Duration(math.Round(2 * float64(time.Second))),
		},
		{
			name:          "Departure in the 1 minute",
			departureTime: now.Add(time.Minute),
			want:          time.Duration(math.Round(2 * float64(time.Second))),
		},
		{
			name:          "Departure in the next 5 minutes",
			departureTime: now.Add(5 * time.Minute),
			want:          calculateExpectedTTL(5*60, 1.5, 3, now.Add(5*time.Minute)), // Use the same logic as in getAdaptiveTTL for calculation
		},
		{
			name:          "Departure in the next 30 minutes",
			departureTime: now.Add(30 * time.Minute),
			want:          calculateExpectedTTL(30*60, 1.5, 3, now.Add(30*time.Minute)),
		},
		{
			name:          "Weekend adjustment",
			departureTime: time.Date(2024, 2, 17, 12, 0, 0, 0, time.UTC),                                          // A Saturday
			want:          calculateExpectedTTL(5*60, 1.5*1.2, 3, time.Date(2022, 10, 15, 12, 5, 0, 0, time.UTC)), // Assuming 5 minutes into the future with weekend multiplier
		},
		{
			name:          "Morning peak hours",
			departureTime: time.Date(2024, 2, 14, 8, 0, 0, 0, time.UTC),                                          // A weekday at 8 AM
			want:          calculateExpectedTTL(5*60, 1.5*1.1, 3, time.Date(2022, 10, 18, 8, 5, 0, 0, time.UTC)), // Assuming 5 minutes into the future with peak hours multiplier
		},
		{
			name:          "Evening peak hours",
			departureTime: time.Date(2024, 2, 14, 19, 0, 0, 0, time.UTC),                                          // A weekday at 6 PM
			want:          calculateExpectedTTL(5*60, 1.5*1.1, 3, time.Date(2022, 10, 18, 18, 5, 0, 0, time.UTC)), // Assuming 5 minutes into the future with peak hours multiplier
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(s.calculateAdaptiveTTL(tt.departureTime, now))
			//if got := s.getAdaptiveTTL(tt.departureTime, now); got != tt.want {
			//	t.Errorf("getAdaptiveTTL() = %v, want %v", got, tt.want)
			//}
		})
	}
}

// calculateExpectedTTL mimics the logic in getAdaptiveTTL for determining the expected TTL in tests
func calculateExpectedTTL(secs float64, multiplier, base float64, departureTime time.Time) time.Duration {
	// Adjust multiplier for weekends
	weekday := departureTime.Weekday()
	if weekday == time.Saturday || weekday == time.Sunday {
		multiplier *= 1.2 // Weekend adjustment
	}

	// Adjust multiplier based on the time of day
	hourOfDay := departureTime.Hour()
	if hourOfDay >= 7 && hourOfDay <= 10 { // Morning peak hours
		multiplier *= 1.1
	} else if hourOfDay >= 17 && hourOfDay <= 19 { // Evening peak hours
		multiplier *= 1.1
	}

	// Calculate and return the TTL
	return time.Duration(math.Round(multiplier * math.Max(base, math.Pow(secs, 0.5)) * float64(time.Second)))
}
