package flyweight

type Context struct {
	id string
	star int
}

func NewContext(id string, star int) *Context {
	return &Context{id: id, star: star}
}
