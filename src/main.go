package main

import (
	"fmt"
	"os"
)

const (
	helpPrompt = `Usage: colorconverter [-h | --help] <command> [-i [<color format>]] [-o [<color format>]] color

There are the supported commands:
 convert  Used to covert an input color format into the specified color format:
          Eg. colorconverter convert -i hex -o rgb '#5fafff'
              rgb(95,175,255)
	
 info     Used to show all conversions on th supported color formats:
          Eg. colorconverter all -i hex '#5fafff'
              Color conversions:
              Hex     #5FAFFF
              RGB     rgb(95,175,255)
              X11     Recommended closest value: 75

Supported color formats:
  hex 
  rgb
  x11
`
)

func main() {
	args := os.Args

	if args[1] == "-h" || args[1] == "--help" {
		showHelp()
	}

}

func showHelp() {
	fmt.Print(helpPrompt)
}
