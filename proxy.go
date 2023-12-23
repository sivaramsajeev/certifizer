package main

type IProxy interface {
	setUp() error
	retstart() error
}
