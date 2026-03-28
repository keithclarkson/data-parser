package data_parser

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type JSONResponse struct {
	JSON map[string]interface{} `json:"json"`
}

func GetJSONResponse(url string) (*JSONResponse, error) {
	resp, err := http.Get(url)
	if err!= nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode!= http.StatusOK {
		return nil, fmt.Errorf("invalid status code: %d", resp.StatusCode)
	}

	var jsonResp JSONResponse
	if err := json.NewDecoder(resp.Body).Decode(&jsonResp); err!= nil {
		return nil, err
	}

	return &jsonResp, nil
}

func GetJSONValue(data interface{}, path...string) (interface{}, error) {
	var currentValue interface{}
	for _, key := range path {
		if currentValue == nil {
			return nil, fmt.Errorf("path is empty")
		}
		switch value := currentValue.(type) {
		case map[string]interface{}:
			if value == nil {
				return nil, fmt.Errorf("path is empty")
			}
			if value, ok := value[key]; ok {
				currentValue = value
			} else {
				return nil, fmt.Errorf("key %s not found in path", key)
			}
		case []interface{}:
			if value == nil {
				return nil, fmt.Errorf("path is empty")
			}
			if i, ok := value[key]; ok {
				currentValue = i
			} else {
				return nil, fmt.Errorf("key %s not found in path", key)
			}
		default:
			return nil, fmt.Errorf("value type %T is not supported", value)
		}
	}
	return currentValue, nil
}

func IsJSON(data []byte) bool {
	var jsonResp JSONResponse
	err := json.Unmarshal(data, &jsonResp)
	return err == nil
}

func IsDate(dateStr string) bool {
	_, err := time.Parse(time.RFC3339, dateStr)
	return err == nil
}

func TrimString(s string) string {
	return strings.TrimSpace(s)
}

func ReadFile(filePath string) (string, error) {
	data, err := io.ReadFile(filePath)
	if err!= nil {
		return "", err
	}
	return string(data), nil
}