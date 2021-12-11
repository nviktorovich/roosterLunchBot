package processors

import (
	"fmt"
	"time"
)

type Representer interface {
	Show() string
	Renew()
}

type LogRow struct {
	T   time.Time // должно быть текущее время
	Log error
}

func (L *LogRow) Show() string {
	date := L.T.Format("2006/01/02---15:04:05.000")
	return fmt.Sprintf("%v, %s", date, L.Log)
}

func (L *LogRow) Renew() {
	L.T = time.Now()
}

type CloseMessage struct {
	T    time.Time
	Info string
}

func (c *CloseMessage) Show() string {
	date := c.T.Format("2006.January.02  15:04")
	return fmt.Sprintf("%v, \n%s", date, c.Info)
}

func (c *CloseMessage) Renew() {
	c.T = time.Now()
}
