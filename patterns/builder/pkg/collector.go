package pkg

const (
	AsusCollectorType = "asus"
	HpCollectorType   = "hp"
)

// интерфейс сборщика
type Collector interface {
	SetCore()
	SetBrand()
	SetMemory()
	SetGraphicCard()
	SetMonitor()
	GetComputer() Computer
}

func GetCollector(collectorType string) Collector {
	switch collectorType {
	default:
		return nil
	case AsusCollectorType:
		return &AsusCollector{}
	case HpCollectorType:
		return &HpCollector{}
	}
}
