package ass1

type Observable interface {
	subscribe(observer Observer) (bool, error)
	unsubscribe(observer Observer) (bool, error)
	sendAll() (bool, error)
}
