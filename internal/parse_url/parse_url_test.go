package parse_url

import "testing"

// tests defines table test cases for a ParseUrl function.
// The expected behavior is to remove only the fragment while preserving
// scheme, user info, host, port, escaping, and query parameters.
var tests = []struct {
	name      string // test name
	input     string // raw input url
	expected  string // expected clean url
	expectErr bool
}{
	{
		name:     "simple https with fragment",
		input:    "https://example.com/path#section",
		expected: "https://example.com/path",
	},
	{
		name:     "query and fragment",
		input:    "https://example.com/users?page=2&limit=10#top",
		expected: "https://example.com/users?page=2&limit=10",
	},
	{
		name:     "no fragment",
		input:    "https://example.com/api/v1/health",
		expected: "https://example.com/api/v1/health",
	},
	{
		name:     "root with fragment",
		input:    "https://example.com/#home",
		expected: "https://example.com/",
	},
	{
		name:     "port and fragment",
		input:    "http://localhost:8080/test#debug",
		expected: "http://localhost:8080/test",
	},
	{
		name:     "userinfo host fragment",
		input:    "https://user:pass@example.com/private#token",
		expected: "https://user:pass@example.com/private",
	},
	{
		name:     "ipv6 host",
		input:    "http://[::1]:8080/path?q=1#frag",
		expected: "http://[::1]:8080/path?q=1",
	},
	{
		name:     "escaped path fragment",
		input:    "https://example.com/a%2Fb/c#part-2",
		expected: "https://example.com/a%2Fb/c",
	},
	{
		name:     "query only fragment",
		input:    "https://example.com?x=1#y",
		expected: "https://example.com?x=1",
	},
	{
		name:     "empty fragment",
		input:    "https://example.com/path#",
		expected: "https://example.com/path",
	},
	{
		name:      "empty url",
		input:     "",
		expectErr: true,
		expected:  "",
	},
	{
		name:      "bad url",
		input:     "://bad-url",
		expectErr: true,
		expected:  "",
	},
	{
		name:      "whitespace input",
		input:     "   ",
		expectErr: true,
		expected:  "",
	},
	{
		name:      "unsupported scheme",
		input:     "ftp://example.com/path",
		expected:  "",
		expectErr: true,
	},
	{
		name:      "missing host",
		input:     "http:///path",
		expected:  "",
		expectErr: true,
	},
	{
		name:      "malformed URL",
		input:     "://missing-scheme.com",
		expected:  "",
		expectErr: true,
	},
}

func TestParseUrl(t *testing.T) {

	for i, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			t.Logf("%d): %s, expected: %s", i+1, tc.name, tc.expected)
			get, err := ParseUrl(tc.input)
			if get != tc.expected {
				t.Fatalf("TEST FAILED: name: %s, want: %s, get: %s", tc.name, tc.expected, get)
			}
			if tc.expectErr && err == nil {
				t.Fatalf("TEST FAILED: expected error, got nil: (name=%s)", tc.name)
			}
			if !tc.expectErr && err != nil {
				t.Fatalf("TEST FAILED: unexpected error: %v, (name=%s)", err, tc.name)
			}
		})
	}
}
