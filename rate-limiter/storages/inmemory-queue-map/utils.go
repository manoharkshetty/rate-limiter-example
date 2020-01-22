package inmemory_queue_map

import (
	"log"
)

func DoPanicRecover() {
	if r := recover(); r != nil {
		log.Print( "panics_recover", r)
	}
}