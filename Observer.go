package main

type Observer interface {
	HandleEvent(vacanies []string)
}
