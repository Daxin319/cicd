package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		input http.Header
		want  string
	}{
		"normal":              {input: http.Header{"Authorization": []string{"ApiKey CorrectHorseBatteryStaple"}}, want: "CorrectHorseBatteryStaple"},
		"malformedNoSpace":    {input: http.Header{"Authorization": []string{"ApiKeyCorrectHorseBatteryStaple"}}, want: ""},
		"noAuthHeader":        {input: http.Header{"Auth": []string{"ApiKey CorrectHorseBatteryStaple"}}, want: ""},
		"malformedWrongTitle": {input: http.Header{"Authorization": []string{"ApiKay CorrectHorseBatteryStaple"}}, want: ""},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, _ := GetAPIKey(tc.input)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}
