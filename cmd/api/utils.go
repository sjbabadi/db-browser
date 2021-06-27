package main

import "log"

func (app *application) check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
