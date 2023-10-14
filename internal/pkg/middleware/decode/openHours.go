package decode

import (
	"encoding/json"

	"github.com/SicParv1sMagna/mdhh_backend/internal/model"
)

func UnmarshalOpenHours(data []byte) ([]model.OpenHoursType, error) {
	if len(data) == 0 || data[0] == '{' && data[1] == '}' {
		return []model.OpenHoursType{}, nil
	}

	var openHours []model.OpenHoursType
	if err := json.Unmarshal(data, &openHours); err != nil {
		return nil, err
	}
	return openHours, nil
}
