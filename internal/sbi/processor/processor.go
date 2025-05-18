package processor

import "github.com/NYCU-CSCS20047-PoCaWN/lab4-af/pkg/app"

type ProcessorNf interface {
	app.App

	Processor() *Processor
}

type Processor struct {
	ProcessorNf
}

func NewProcessor(nf ProcessorNf) (*Processor, error) {
	p := &Processor{
		ProcessorNf: nf,
	}
	return p, nil
}
