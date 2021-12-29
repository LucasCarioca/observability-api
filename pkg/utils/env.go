package utils

import (
	"flag"
	"fmt"
	"os"
)

//GetEnv returns the current application environment
//will look in the local system variables for a 'ENV' first then the 'e' command line flag
//defaults to 'dev' if neither is found
func GetEnv() string {
	env := os.Getenv("ENV")
	if env == "" {
		envFlag := flag.String("e", "dev", "")
		flag.Usage = func() {
			fmt.Println("Usage: server -e {mode}")
			os.Exit(1)
		}
		flag.Parse()
		env = *envFlag
	}
	return env
}
