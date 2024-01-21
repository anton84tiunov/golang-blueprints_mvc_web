package logprint

import (
	"fmt"
	"log"

	config "github.com/anton84tiunov/golang-blueprints_mvc_web/internal/config"
	// "os"
)

func Print(msg string) {
	if config.GLOBAL_CONFIG.Server.Debug {
		fmt.Println(msg)
	} else {
		log.Println(msg)
	}
}
