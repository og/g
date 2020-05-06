package gcrypto

import (
	"log"
)

var warningLog = func(message string) {
	log.Print("gcrypto:" + message)
}
