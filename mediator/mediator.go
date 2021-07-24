package mediator

type Mediator interface {
	CanLand(train Train) bool
	NotifyFree()
}
