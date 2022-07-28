package main

type Database interface {
	GetUser() string
}

type DefaultDatabase struct{}

func (db DefaultDatabase) GetUser() string {
	return "bob"
}


