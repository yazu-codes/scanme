package utils

import "encoding/json"

func JsonStringToArray(menuIds string) ([]int64, error) {
	var menuIDs []int64

	if err := json.Unmarshal(
		[]byte(menuIds),
		&menuIDs,
	); err != nil {
		return nil, err
	}

	return menuIDs, nil
}
