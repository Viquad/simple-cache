package cache

import (
	"testing"
	"time"
)

func TestMemoryCache(t *testing.T) {
	testTableSetCases := []struct {
		key           string
		value         interface{}
		ttl           time.Duration
		expectedError string
	}{
		{
			key:           "test_key",
			value:         "test_value",
			expectedError: "",
		},
		{
			key:           "test_key2",
			value:         "test_value",
			expectedError: "",
			ttl:           time.Millisecond * 10,
		},
		{
			key:           "test_key3",
			value:         "test_value",
			expectedError: "",
			ttl:           time.Millisecond * 10,
		},
		{
			key:           "test_key3",
			value:         "test_value",
			expectedError: "",
			ttl:           time.Millisecond * 100,
		},
	}

	testTableGetCases := []struct {
		key           string
		expectedValue interface{}
		expectedError string
	}{
		{
			key:           "test_key",
			expectedValue: "test_value",
			expectedError: "",
		},
		{
			key:           "test_key2",
			expectedValue: nil,
			expectedError: "missed value by test_key2 key",
		},
		{
			key:           "test_key3",
			expectedValue: "test_value",
			expectedError: "",
		},
		{
			key:           "unknown_key",
			expectedValue: nil,
			expectedError: "missed value by unknown_key key",
		},
	}

	testTableDeleteCases := []struct {
		key           string
		expectedError string
	}{
		{
			key:           "test_key",
			expectedError: "",
		},
		{
			key:           "test_key",
			expectedError: "attempt to delete missed value by test_key key",
		},
	}

	cache := NewMemoryCache()

	for _, testCase := range testTableSetCases {
		resultError := ""
		if err := cache.Set(testCase.key, testCase.value, testCase.ttl); err != nil {
			resultError = err.Error()
		}

		if resultError != testCase.expectedError {
			t.Errorf("Incorrect result. Expected %v, got %v", testCase.expectedError, resultError)
		}
	}

	time.Sleep(time.Millisecond * 20)

	for _, testCase := range testTableGetCases {
		resultError := ""
		resultValue, err := cache.Get(testCase.key)
		if err != nil {
			resultError = err.Error()
		}

		if resultError != testCase.expectedError || resultValue != testCase.expectedValue {
			t.Errorf("Incorrect result. Expected value: %v, error %v, got value: %v, error %v", testCase.expectedValue, testCase.expectedError, resultValue, resultError)
		}
	}

	for _, testCase := range testTableDeleteCases {
		resultError := ""
		if err := cache.Delete(testCase.key); err != nil {
			resultError = err.Error()
		}

		if resultError != testCase.expectedError {
			t.Errorf("Incorrect result. Expected %v, got %v", testCase.expectedError, resultError)
		}
	}
}
