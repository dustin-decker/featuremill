package extractor

// feature index determinestic by murmur32(uniqueHashPrefixStr+<user field>+<interval>)

import (
	"fmt"

	"github.com/araddon/dateparse"
	"github.com/spaolacci/murmur3"
)

// TransformDate returns a slice of 2 seasonality vectors: day of week, and month of year
// for the deterministic feature ID
func TransformDate(field, date string) ([]string, error) {
	var out []string

	dt, err := dateparse.ParseAny(date)
	if err != nil {
		return out, err
	}

	dayOfWeek := int(dt.Weekday())
	dayOfWeekFeatureID := murmur3.Sum32([]byte(uniqueHashPrefixStr + field + "dayOfWeek"))
	out = append(out, fmt.Sprintf("%d:%f", dayOfWeekFeatureID, float32(dayOfWeek)/6))

	monthOfYear := dt.Month()
	monthOfYearFeatureID := murmur3.Sum32([]byte(uniqueHashPrefixStr + field + "monthOfYear"))
	out = append(out, fmt.Sprintf("%d:%f", monthOfYearFeatureID, float32(monthOfYear)/12))

	return out, nil
}
