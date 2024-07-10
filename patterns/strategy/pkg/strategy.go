package pkg

type Strategy interface {
	Route(startPoint, endPoint int)
}