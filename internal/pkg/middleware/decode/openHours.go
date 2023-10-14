package decode

import (
	"encoding/json"

	"github.com/SicParv1sMagna/mdhh_backend/internal/model"
)

func UnmarshalOpenHours(data []byte) ([]model.OpenHoursType, error) {
	var openHours []model.OpenHoursType
	if err := json.Unmarshal(data, &openHours); err != nil {
		return nil, err
	}
	return openHours, nil
}
