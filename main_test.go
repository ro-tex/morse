package main

import "testing"

func TestEncode(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected string
	}{
		"empty": {input: "", expected: ""},
		"hello": {input: "hello world", expected: ".... . .-.. .-.. ---  .-- --- .-. .-.. -.."},
		"SOS":   {input: "SOS", expected: "... --- ..."},
		"mixed": {
			input:    "this $%^ is =-a mixed 123 me$$age",
			expected: "- .... .. ...  .. ...  .-  -- .. -..- . -..  ..--- ...-- ....-  -- .  .- --. .",
		},
	}

	for name, tt := range tests {
		out := Encode(tt.input)
		if out != tt.expected {
			t.Errorf("Test %s: expected '%s', got '%s'", name, tt.expected, out)
		}
	}
}

func TestDecode(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected string
	}{
		"empty": {input: "", expected: ""},
		"hello": {input: ".... . .-.. .-.. ---  .-- --- .-. .-.. -..", expected: "hello world"},
		"SOS":   {input: "... --- ...", expected: "sos"},
		"mixed": {
			input:    "- .... .. ... / .. ... $ .-  -- .. -..- . -.. @ ..--- ...-- ....-  -- .  .- --. .  ",
			expected: "this is a mixed 123 me age",
		},
	}

	for name, tt := range tests {
		out := Decode(tt.input)
		if out != tt.expected {
			t.Errorf("Test %s: expected '%s', got '%s'", name, tt.expected, out)
		}
	}
}
