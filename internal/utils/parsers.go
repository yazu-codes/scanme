package utils

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

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

func StringToIntArray(menuIds string) ([]int64, error) {
	var menuIDs []int64

	menuIds = strings.TrimSpace(menuIds)

	if menuIds == "" {
		return menuIDs, nil
	}

	for _, value := range strings.Split(menuIds, ",") {
		value = strings.TrimSpace(value)

		if value == "" {
			continue
		}

		id, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, fmt.Errorf(
				"invalid menu id %q: %w",
				value,
				err,
			)
		}

		menuIDs = append(menuIDs, id)
	}

	return menuIDs, nil
}
