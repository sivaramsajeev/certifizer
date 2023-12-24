package main

type Domain struct {
	name  string
	email string
}

func (d *Domain) validate() {
	d.init()
	// d.confirmOwnership()
}

func (d *Domain) init() {
	config := readConfig()
	config.validate()
}
