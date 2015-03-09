package uuid

import "testing"

type test struct {
	in, out string
}

var uuidTests = []test{
	// Invalid
	{
		in:  "", // Blank
		out: "",
	},
	{
		in:  "95706b34-xxxx-4655-8a5f-5f586fa0a37b", // Invalid chars
		out: "",
	},
	{
		in:  "fa94d30-5b9e-4449-90fd-b816ea08f", // Too short
		out: "",
	},
	{
		in:  "6471412c-9592-441a-8c5fb5-99a6e19645",
		out: "",
	},

	// Valid
	{
		in:  "6471412c-9592-441a-8c5f-b599a6e19645",
		out: "6471412c-9592-441a-8c5f-b599a6e19645",
	},
	{
		in:  "6471412C-9592-441A-8C5F-B599A6E19645",
		out: "6471412c-9592-441a-8c5f-b599a6e19645",
	},
}

func BenchmarkRegexp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, uuid := range uuidTests {
			_, _ = Regex(uuid.in)
		}
	}
}

func BenchmarkRunes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, uuid := range uuidTests {
			_, _ = Runes(uuid.in)
		}
	}
}

func BenchmarkBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, uuid := range uuidTests {
			_, _ = Bytes(uuid.in)
		}
	}
}

func BenchmarkUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, uuid := range uuidTests {
			_, _ = UUID(uuid.in)
		}
	}
}

func BenchmarkFuncs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, uuid := range uuidTests {
			_, _ = Funcs(uuid.in)
		}
	}
}

func TestRegexp(t *testing.T) {
	for _, tt := range uuidTests {
		result, err := Regex(tt.in)
		if tt.out != result {
			t.Errorf("normalize Regex(%s) = \n%s, wanted \n%s", tt.in, result, tt.out)
		}
		if err != nil && len(result) > 0 {
			t.Errorf("normalize Regex(%s) should not return results when an error is present")
		}
	}
}

func TestRunes(t *testing.T) {
	for _, tt := range uuidTests {
		result, err := Runes(tt.in)
		if tt.out != result {
			t.Errorf("normalize Runes(%s) = \n%s, wanted \n%s", tt.in, result, tt.out)
		}
		if err != nil && len(result) > 0 {
			t.Errorf("normalize Runes(%s) should not return results when an error is present")
		}
	}
}

func TestBytes(t *testing.T) {
	for _, tt := range uuidTests {
		result, err := Bytes(tt.in)
		if tt.out != result {
			t.Errorf("normalize Bytes(%s) = \n%s, wanted \n%s", tt.in, result, tt.out)
		}
		if err != nil && len(result) > 0 {
			t.Errorf("normalize Bytes(%s) should not return results when an error is present")
		}
	}
}

func TestUUID(t *testing.T) {
	for _, tt := range uuidTests {
		result, err := UUID(tt.in)
		if tt.out != result {
			t.Errorf("normalize UUID(%s) = \n%s, wanted \n%s", tt.in, result, tt.out)
		}
		if err != nil && len(result) > 0 {
			t.Errorf("normalize UUID(%s) should not return results when an error is present")
		}
	}
}

func TestFuncs(t *testing.T) {
	for _, tt := range uuidTests {
		result, err := Funcs(tt.in)
		if tt.out != result {
			t.Errorf("normalize Funcs(%s) = \n%s, wanted \n%s", tt.in, result, tt.out)
		}
		if err != nil && len(result) > 0 {
			t.Errorf("normalize Funcs(%s) should not return results when an error is present")
		}
	}
}
