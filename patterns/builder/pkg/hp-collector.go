package pkg

// те же самые поля, что и у Computer
// реализация интерфейса сборщика hp collector
type HpCollector struct {
	Core       int
	Brand      string
	Memory     int
	Monitor    int
	GraphicCard int
}

func (collector *HpCollector) SetCore() {
	collector.Core = 4
}

func (collector *HpCollector) SetBrand() {
	collector.Brand = "Hp"
}

func (collector *HpCollector) SetMemory() {
	collector.Memory = 16
}


func (collector *HpCollector) SetMonitor() {
	collector.Monitor = 2
}

func (collector *HpCollector) SetGraphicCard() {
	collector.GraphicCard = 1
}


func (collector *HpCollector) GetComputer() Computer {
	return Computer{
		Core: collector.Core,
		Brand: collector.Brand,
		Memory: collector.Memory,
		Monitor: collector.Monitor,
		GraphicCard: collector.GraphicCard,
	}
}


