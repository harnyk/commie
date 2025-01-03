package banner

import (
	"fmt"
	"strings"

	"github.com/mgutz/ansi"
)

// From: https://manytools.org/hacker-tools/ascii-banner/ font: "Rowan Cap"
var banner = `
   .aMMMb  .aMMMb  dMMMMMMMMb  dMMMMMMMMb  dMP dMMMMMP 
  dMP"VMP dMP"dMP dMP"dMP"dMP dMP"dMP"dMP amr dMP      
 dMP     dMP dMP dMP dMP dMP dMP dMP dMP dMP dMMMP     
dMP.aMP dMP.aMP dMP dMP dMP dMP dMP dMP dMP dMP        
VMMMP"  VMMMP" dMP dMP dMP dMP dMP dMP dMP dMMMMMP     
`

func efr(s string) string {
	runes := []rune(s)
	head := string(runes[0])
	tail := string(runes[1:])
	return ansi.Color(head, "green+b") + tail
}

func GetBanner() string {
	sloganWords := []string{
		efr("Code"),
		efr("Operations"),
		efr("Made"),
		efr("More"),
		efr("Intelligent"),
		"&",
		efr("Efficient"),
	}
	return fmt.Sprintf("%s\n%s\n", banner, strings.Join(sloganWords, " "))
}
