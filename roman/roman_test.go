package roman

import (
	"fmt"
	"log"
	"testing"
	"testing/quick"
)

var cases = []struct {
	Description string
	Arabic      int
	Roman       string
}{
	{
		Description: "1 gets converted to I",
		Arabic:      1,
		Roman:       "I",
	},

	{
		Description: "2 gets converted to II",
		Arabic:      2,
		Roman:       "II",
	},
	{
		Description: "3 gets converted to III",
		Arabic:      3,
		Roman:       "III",
	},
	{
		Description: "4 gets converted to IV",
		Arabic:      4,
		Roman:       "IV",
	},
	{
		Description: "5 gets converted to V",
		Arabic:      5,
		Roman:       "V",
	},
	{
		Description: "6 gets converted to VI",
		Arabic:      6,
		Roman:       "VI",
	},
	{
		Description: "7 gets converted to VII",
		Arabic:      7,
		Roman:       "VII",
	},
	{
		Description: "8 gets converted to VIII",
		Arabic:      8,
		Roman:       "VIII",
	},
	{
		Description: "9 gets converted to IX",
		Arabic:      9,
		Roman:       "IX",
	},
	{
		Description: "10 gets converted to X",
		Arabic:      10,
		Roman:       "X",
	},
	{
		Description: "14 gets converted to XIV",
		Arabic:      14,
		Roman:       "XIV",
	},
	{
		Description: "18 gets converted to XVIII",
		Arabic:      18,
		Roman:       "XVIII",
	},
	{
		Description: "20 gets converted to XX",
		Arabic:      20,
		Roman:       "XX",
	},
	{
		Description: "39 gets converted to XXXIX",
		Arabic:      39,
		Roman:       "XXXIX",
	},
	{
		Description: "40 gets converted to XL",
		Arabic:      40,
		Roman:       "XL",
	},
	{
		Description: "41 gets converted to XLI",
		Arabic:      41,
		Roman:       "XLI",
	},
	{
		Description: "48 gets converted to XLVIII",
		Arabic:      48,
		Roman:       "XLVIII",
	},
	{
		Description: "49 gets converted to XLIX",
		Arabic:      49,
		Roman:       "XLIX",
	},
	{
		Description: "50 gets converted to L",
		Arabic:      50,
		Roman:       "L",
	},
	{
		Description: "100 gets converted to C",
		Arabic:      100,
		Roman:       "C",
	},
	{
		Description: "100 gets converted to XC",
		Arabic:      90,
		Roman:       "XC",
	},
	{
		Description: "400 gets converted to CD",
		Arabic:      400,
		Roman:       "CD",
	},
	{
		Description: "500 gets converted to D",
		Arabic:      500,
		Roman:       "D",
	},
	{
		Description: "900 gets converted to CM",
		Arabic:      900,
		Roman:       "CM",
	},
	{
		Description: "1000 gets converted to M",
		Arabic:      1000,
		Roman:       "M",
	},
	{
		Description: "1984 gets converted to MCMLXXXIV",
		Arabic:      1984,
		Roman:       "MCMLXXXIV",
	},
	{
		Description: "3999 gets converted to MMMCMXCIX",
		Arabic:      3999,
		Roman:       "MMMCMXCIX",
	},
	{
		Description: "2014 gets converted to MMXIV",
		Arabic:      2014,
		Roman:       "MMXIV",
	},
	{
		Description: "1006 gets converted to MVI",
		Arabic:      1006,
		Roman:       "MVI",
	},
	{
		Description: "798 gets converted to DCCXCVIII",
		Arabic:      798,
		Roman:       "DCCXCVIII",
	},
}

func TestRomanNumerals(t *testing.T) {

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			want := test.Roman

			if got != want {
				t.Errorf("got '%s', want '%s'", got, want)
			}
		})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)

			if got != test.Arabic {
				t.Errorf("got '%d', want '%d'", got, test.Arabic)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic int) bool {
		if arabic < 0 || arabic > 3999 {
			log.Println(arabic)
			return true
		}
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}
	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("failed checks", err)
	}
}
