package auth

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

// func TestGetAPIKey(t *testing.T) {
// 	Header := http.Header{
// 		"Accept-Encoding": {"gzip, deflate"},
// 		"Accept-Language": {"en-us"},
// 		"Authorization":   {"ApiKey ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb"},
// 	}
// 	got, err := GetAPIKey(Header)
// 	if err != nil {
// 		fmt.Printf("Error anjing %v", err)
// 	}
// 	want := "ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb"
// 	if !reflect.DeepEqual(want, got) {
// 		t.Fatalf("expected: %v, got: %v", want, got)
// 	}
// }

// package auth

// import (
// 	"fmt"
// 	"net/http"
// 	"strings"
// 	"testing"
// )

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		key       string
		value     string
		expect    string
		expectErr string
	}{
		{
			expectErr: "no authorization header",
		},
		{
			key:       "Authorization",
			expectErr: "no authorization header",
		},
		{
			key:       "Authorization",
			value:     "-",
			expectErr: "malformed authorization header",
		},
		{
			key:       "Authorization",
			value:     "Bearer xxxxxx",
			expectErr: "malformed authorization header"},
		{
			key:       "Authorization",
			value:     "ApiKey xxxxxx",
			expect:    "xxxxxx",
			expectErr: "not expecting an error",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestGetAPIKey Case #%v:", i), func(t *testing.T) {
			header := http.Header{}
			header.Add(test.key, test.value)

			output, err := GetAPIKey(header)
			if err != nil {
				if strings.Contains(err.Error(), test.expectErr) {
					return
				}
				t.Errorf("Unexpected: TestGetAPIKey:%v\n", err)
				return
			}

			if output != test.expect {
				t.Errorf("Unexpected: TestGetAPIKey:%s", output)
				return
			}
		})
	}
}
