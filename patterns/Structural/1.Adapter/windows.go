package main

import "fmt"

// Unknow service

type Windows struct {
}

func (w *Windows) InsertIntoLightningPort() {
	fmt.Println("USB connector is plugged into windows machine")
}
