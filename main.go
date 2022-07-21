package main

import (
	"regexp"
	"strings"
)

var alphabetDecode = map[string]string{
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

var alphabetEncode = map[string]string{
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

func Encode(ss string) string {
	res := strings.Builder{}
	for _, s := range strings.Split(ss, "") {
		res.WriteString(alphabetEncode[s])
		res.WriteString(" ")
	}
	return res.String()
}

func Decode(ss string) string {
	res := strings.Builder{}
	for _, s := range strings.Split(clean(ss), " ") {
		if len(s) == 0 {
			res.WriteString(" ")
			continue
		}
		res.WriteString(alphabetDecode[s])
	}
	return res.String()
}

func clean(s string) string {
	return regexp.MustCompile(`[^\.\-\s]`).ReplaceAllString(s, " ")
}

func main() {
	s := ".. - .----. ... / .- / -... .. -. .- .-. -.-- / ... -.-- ... - . -- / --- ..-. / -.-. --- -- -- ..- -. .. -.-. .- - .. --- -."
	println(Decode(s))

	println(Decode(Encode("hello world!")))
}
