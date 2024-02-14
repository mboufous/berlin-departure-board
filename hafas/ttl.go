package hafas

import (
	"math"
	"time"
)

func (s *DepartureService) getAdaptiveTTL(currentTime time.Time, lines []Line) time.Duration {

	imminentDepartureTime := getImminentDepartureTime(lines)
	return s.calculateAdaptiveTTL(imminentDepartureTime, currentTime)
}

func (s *DepartureService) calculateAdaptiveTTL(imminentDepartureTime, currentTime time.Time) time.Duration {
	multiplier := 1.5
	fallback := 0.0
	base := 0.0
	//
	//if imminentDepartureTime.Before(currentTime) {
	//	return 0
	//}

	secs := imminentDepartureTime.Sub(currentTime).Seconds()
	weekday := imminentDepartureTime.Weekday()

	// Adjust multiplier for weekends
	if weekday == time.Saturday || weekday == time.Sunday {
		multiplier *= 1.2
	}

	// Adjust multiplier based on the time of day
	hourOfDay := imminentDepartureTime.Hour()
	if hourOfDay >= 7 && hourOfDay <= 10 { // Morning peak hours
		multiplier *= 1.1
	} else if hourOfDay >= 17 && hourOfDay <= 19 { // Evening peak hours
		multiplier *= 1.1
	}

	if secs > 0 {
		return time.Duration(math.Round(
			multiplier *
				math.Max(base, math.Pow(secs, 0.5)) *
				float64(time.Second),
		))
	}
	return time.Duration(math.Round(multiplier * fallback * float64(time.Second)))
}

func getImminentDepartureTime(lines []Line) time.Time {
	if len(lines) == 0 {
		return time.Time{}
	}
	var imminentDepartureTime time.Time

	for _, line := range lines {
		for _, direction := range line.Directions {
			if imminentDepartureTime.IsZero() || direction.Departures[0].When.Before(imminentDepartureTime) {
				imminentDepartureTime = direction.Departures[0].When
			}
		}
	}

	return imminentDepartureTime
}
