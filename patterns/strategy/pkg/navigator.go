package pkg

type Navigator struct {
	Strategy
}

func (nav *Navigator) SetStrategy(str Strategy) {
	nav.Strategy = str
}