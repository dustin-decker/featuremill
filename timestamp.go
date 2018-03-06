package featuremill

// feature index determinestic by murmur32(uniqueHashPrefixStr+<user field>+<interval>)

import (
	"fmt"

	"github.com/araddon/dateparse"
	"github.com/spaolacci/murmur3"
)

// ExtractTimestamp returns a slice of 3 seasonality vectors: minute of hour, hour of day, day of week
func ExtractTimestamp(field, timestamp string) ([]string, error) {
	var out []string

	dt, err := dateparse.ParseAny(timestamp)
	if err != nil {
		return out, err
	}

	dayOfWeek := int(dt.Weekday())
	dayOfWeekFeatureID := int32(murmur3.Sum32([]byte(uniqueHashPrefixStr + field + "dayOfWeek")))
	out = append(out, fmt.Sprintf("%d:%f", dayOfWeekFeatureID, float32(dayOfWeek)/6))

	hourOfDay := dt.Hour()
	hourOfDayFeatureID := int32(murmur3.Sum32([]byte(uniqueHashPrefixStr + field + "hourOfDay")))
	out = append(out, fmt.Sprintf("%d:%f", hourOfDayFeatureID, float32(hourOfDay)/23))

	minuteOfHour := dt.Minute()
	minuteOfHourFeatureID := int32(murmur3.Sum32([]byte(uniqueHashPrefixStr + field + "minuteOfHour")))
	out = append(out, fmt.Sprintf("%d:%f", minuteOfHourFeatureID, float32(minuteOfHour)/59))

	return out, nil
}
