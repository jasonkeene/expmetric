package expmetric

import (
	"expvar"
	"fmt"
)

type Counter struct {
	expvar.Int
}

func (c *Counter) String() string {
	return fmt.Sprintf(`{"type":"counter","value":%s}`, c.Int.String())
}

func NewCounter(name string) *Counter {
	v := &Counter{}
	expvar.Publish(name, v)
	return v
}
