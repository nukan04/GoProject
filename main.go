package main

func main() {
	hhkz := JobSite{}
	Bob := Person{}
	Bob.set("Bob")
	hhkz.addVacancies("Senior Golang Developer")
	hhkz.subscribe(&Bob)
	hhkz.addVacancies("Junior Golang Developer")
	hhkz.removeVacancies("Senior Golang Developer")

}
