package main

type Observable interface {
	subscribe(observer Observer) (bool, error)
	unsubscribe(observer Observer) (bool, error)
	sendAll() (bool, error)
}
