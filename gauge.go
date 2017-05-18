package expmetric

import (
	"expvar"
	"fmt"
)

type Gauge struct {
	expvar.Float
	unit string
}

func (g *Gauge) String() string {
	return fmt.Sprintf(`{"type":"gauge","unit":"%s","value":%s}`, g.unit, g.Float.String())
}

func NewGauge(name, unit string) *Gauge {
	v := &Gauge{
		unit: unit,
	}
	expvar.Publish(name, v)
	return v
}
