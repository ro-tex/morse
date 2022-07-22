package main

import (
	"regexp"
	"strings"
)

var (
	alphabetEncode = map[string]string{
		"a": ".-",
		"b": "-...",
		"c": "-.-.",
		"d": "-..",
		"e": ".",
		"f": "..-.",
		"g": "--.",
		"h": "....",
		"i": "..",
		"j": ".---",
		"k": "-.-",
		"l": ".-..",
		"m": "--",
		"n": "-.",
		"o": "---",
		"p": ".--.",
		"q": "--.-",
		"r": ".-.",
		"s": "...",
		"t": "-",
		"u": "..-",
		"v": "...-",
		"w": ".--",
		"x": "-..-",
		"y": "-.--",
		"z": "--..",
		"0": ".----",
		"1": "..---",
		"2": "...--",
		"3": "....-",
		"4": ".....",
		"5": "-....",
		"6": "--...",
		"7": "---..",
		"8": "----.",
		"9": "-----",
	}

	alphabetDecode = map[string]string{
		".-":    "a",
		"-...":  "b",
		"-.-.":  "c",
		"-..":   "d",
		".":     "e",
		"..-.":  "f",
		"--.":   "g",
		"....":  "h",
		"..":    "i",
		".---":  "j",
		"-.-":   "k",
		".-..":  "l",
		"--":    "m",
		"-.":    "n",
		"---":   "o",
		".--.":  "p",
		"--.-":  "q",
		".-.":   "r",
		"...":   "s",
		"-":     "t",
		"..-":   "u",
		"...-":  "v",
		".--":   "w",
		"-..-":  "x",
		"-.--":  "y",
		"--..":  "z",
		".----": "0",
		"..---": "1",
		"...--": "2",
		"....-": "3",
		".....": "4",
		"-....": "5",
		"--...": "6",
		"---..": "7",
		"----.": "8",
		"-----": "9",
	}
)

// Encode a string using ITU morse code. Any non-alphanumeric characters are
// replaced by spaces. All strings are treated as lower case.
func Encode(ss string) string {
	res := strings.Builder{}
	trailingSpaces := 0
	for _, s := range strings.Split(strings.ToLower(ss), "") {
		if alphabetEncode[s] != "" {
			res.WriteString(alphabetEncode[s])
			trailingSpaces = 0
		}
		// Two spaces separate words. We don't need more of those.
		if trailingSpaces > 1 {
			continue
		}
		res.WriteString(" ")
		trailingSpaces++
	}
	return strings.Trim(res.String(), " ")
}

// Decode an ITU morse code message. Any characters, other than `-` and `.` are
// treated as spaces. A single space is treated as character separator, multiple
// spaces are treated as word separator.
func Decode(ss string) string {
	res := strings.Builder{}
	lastWasSpace := false
	for _, s := range strings.Split(clean(ss), " ") {
		if alphabetDecode[s] == "" {
			if !lastWasSpace {
				res.WriteString(" ")
				lastWasSpace = true
			}
			continue
		}
		res.WriteString(alphabetDecode[s])
		lastWasSpace = false
	}
	return strings.Trim(res.String(), " ")
}

// clean removes non-morse characters and replaces them with spaces.
func clean(s string) string {
	return regexp.MustCompile(`[^\.\-]`).ReplaceAllString(s, " ")
}

func main() {
	enc := Encode("hello world!")
	println(enc)
	println(Decode(enc))
}
