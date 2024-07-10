package pkg

import "fmt"

type Notebook struct {
	Type    string
	Core    int
	Memory  int
	Monitor bool
}

func NewNotebook() Computer {
	return Notebook{
		Type:    NotebookType,
		Core:    4,
		Memory:  8,
		Monitor: true,
	}
}

func (pc Notebook) GetType() string {
	return pc.Type
}

func (pc Notebook) PrintDetails() {
	fmt.Printf("%s Core:[%d] Mem:[%d]\n", pc.Type, pc.Core, pc.Memory)
}
