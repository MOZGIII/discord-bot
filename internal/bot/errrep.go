package bot

import "log"

func reportErrorInAny(err error) {
	if err != nil {
		reportError(err)
	}
}

func reportError(err error) {
	log.Println(err)
}
