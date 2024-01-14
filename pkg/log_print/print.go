package logprint

import (
	"fmt"
	"log"

	config "../../internal/config"
	// "os"
)

func Print(msg string) {
	if config.GLOBAL_CONFIG.Server.Debug {
		fmt.Println(msg)
	} else {
		log.Println(msg)
	}
}
