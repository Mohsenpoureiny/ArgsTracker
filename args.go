package argstracker

import (
	"os"
	"regexp"
)

type CallbackInfo struct {
	CommandName  string
	CallBackFunc func([]string)
}

func GetArgs() []string {
	var args []string = os.Args
	return args
}
func GetAllSingleFlags(args []string) []string {
	flags := []string{}
	for i := 0; i < len(args); i++ {
		if b, _ := regexp.MatchString("^-[a-zA-Z]+", args[i]); b {
			flags = append(flags, args[i])
		}
	}
	return flags
}
func GetAllDoubleFlags(args []string) []string {
	flags := []string{}
	for i := 0; i < len(args); i++ {
		if b, _ := regexp.MatchString("^--[a-zA-Z]+", args[i]); b {
			flags = append(flags, args[i])
		}
	}
	return flags
}
func GetQuote(args []string, forFlag string, offset ...int) string {
	for i := 0; i < len(args); i++ {
		if args[i] == forFlag {
			if len(args) > i+1 {
				return args[i+1]
			}
		}
	}
	return "NO_QUOTE"
}
func GetMainCommand(args []string, cbis []CallbackInfo, notFound func()) {
	mainCommand := args[1]
	for i := 0; i < len(cbis); i++ {
		if cbis[i].CommandName == mainCommand {
			cbis[i].CallBackFunc(args)
			return
		}
	}
	notFound()
}
