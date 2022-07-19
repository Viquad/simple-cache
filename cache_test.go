package simplecache

import "testing"

func TestMemoryCache(t *testing.T) {
	testTableSetCases := []struct {
		key           string
		value         interface{}
		expectedError string
	}{
		{
			key:           "",
			value:         2,
			expectedError: "empty key",
		},
		{
			key:           "test_key",
			value:         "test_value",
			expectedError: "",
		},
	}

	testTableGetCases := []struct {
		key           string
		expectedValue interface{}
		expectedError string
	}{
		{
			key:           "",
			expectedValue: nil,
			expectedError: "empty key",
		},
		{
			key:           "test_key",
			expectedValue: "test_value",
			expectedError: "",
		},
		{
			key:           "unknown_key",
			expectedValue: nil,
			expectedError: "missed value by unknown_key key in memory cache",
		},
	}

	testTableDeleteCases := []struct {
		key           string
		expectedError string
	}{
		{
			key:           "",
			expectedError: "empty key",
		},
		{
			key:           "test_key",
			expectedError: "",
		},
		{
			key:           "test_key",
			expectedError: "attempt to delete missed value by test_key key in memory cache",
		},
	}

	cache := NewMemoryCache()

	for _, testCase := range testTableSetCases {
		resultError := ""
		if err := cache.Set(testCase.key, testCase.value); err != nil {
			resultError = err.Error()
		}

		if resultError != testCase.expectedError {
			t.Errorf("Incorrect result. Expected %v, got %v", testCase.expectedError, resultError)
		}
	}

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
