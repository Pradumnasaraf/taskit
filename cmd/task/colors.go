package task

import "fmt"

const (
	ColorDefault = "\x1b[39m"

	ColorRed    = "\x1b[91m"
	ColorGreen  = "\x1b[32m"
	ColorBlue   = "\x1b[94m"
	ColorPurple = "\x1b[35m"
)

func red(s string) string {
	return fmt.Sprintf("%s%s%s", ColorRed, s, ColorDefault)
}

func green(s string) string {
	return fmt.Sprintf("%s%s%s", ColorGreen, s, ColorDefault)
}

func purple(s string) string {
	return fmt.Sprintf("%s%s%s", ColorPurple, s, ColorDefault)
}
