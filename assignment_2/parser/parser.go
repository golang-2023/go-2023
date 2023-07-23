package parser

import (
	"strconv"
)

func IntParser(args []string) ([]int, error) {
	result := make([]int, len(args))
	for i, arg := range args {
		val, err := strconv.Atoi(arg)
		if err != nil {
			return nil, err
		}
		result[i] = val
	}
	return result, nil
}

func FloatParser(args []string) ([]float64, error) {
	result := make([]float64, len(args))
	for i, arg := range args {
		val, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			return nil, err
		}
		result[i] = val
	}
	return result, nil
}

func MixedParser(args []string) ([]interface{}, error) {
	result := make([]interface{}, len(args))
	for i, arg := range args {
		if val, err := strconv.Atoi(arg); err == nil {
			result[i] = val
		} else if val, err := strconv.ParseFloat(arg, 64); err == nil {
			result[i] = val
		} else {
			result[i] = arg
		}
	}
	return result, nil
}
