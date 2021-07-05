package MainController

type Doer interface {
	Do()
}

type contrtoller struct {
	name string
}
