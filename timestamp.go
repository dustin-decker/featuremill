package extractor

// feature index determinestic by murmur32(uniqueHashPrefixStr+<user field>+<interval>)

import (
	"fmt"

	"github.com/araddon/dateparse"
	"github.com/spaolacci/murmur3"
)

// TransformTimestamp returns as 3 seasonality vectors: minute of hour, hour of day, day of week
func TransformTimestamp(field, timestamp string) ([]string, error) {
	var out []string

	dt, err := dateparse.ParseAny(timestamp)
	if err != nil {
		return out, err
	}

	dayOfWeek := int(dt.Weekday())
	dayOfWeekFeatureID := murmur3.Sum32([]byte(uniqueHashPrefixStr + field + "dayOfWeek"))
	out = append(out, fmt.Sprintf("%d:%d", dayOfWeekFeatureID, dayOfWeek/6))

	hourOfDay := dt.Hour()
	hourOfDayFeatureID := murmur3.Sum32([]byte(uniqueHashPrefixStr + field + "hourOfDay"))
	out = append(out, fmt.Sprintf(" %d:%d", hourOfDayFeatureID, hourOfDay/23))

	minuteOfHour := dt.Minute()
	minuteOfHourFeatureID := murmur3.Sum32([]byte(uniqueHashPrefixStr + field + "minuteOfHour"))
	out = append(out, fmt.Sprintf(" %d:%d", minuteOfHourFeatureID, minuteOfHour/59))

	return out, nil
}
