package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloat(strings []string) ([]float64, error) {
	floats := make([]float64, len(strings))
	for i, v := range strings {
		float, err := strconv.ParseFloat(v, 64)

		if err != nil {
			return nil, errors.New("float conversion to string failed")
		}

		floats[i] = float
	}
	return floats, nil
}
