package pkg

import "fmt"

type PublicTransportStrategy struct {}

func (r *PublicTransportStrategy) Route(startPoint, endPoint int) {
	avgSpeed := 40
	total := endPoint - startPoint
	totalTime := total * 40
	fmt.Printf("Public transport A:[%d] to B:[%d] Avg speed:[%d] Total:[%d] Total time:[%d] min\n", startPoint, endPoint,
	avgSpeed, total, totalTime)
}