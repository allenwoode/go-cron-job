package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	c := cron.New(cron.WithSeconds())
	//g := New(c)
	for i := 0; i < 2; i++ {
		g := New(c)
		spec := fmt.Sprintf("*/%d * * * * ?", i+1)
		id, _ := c.AddJob(spec, g)
		g.SetId(id)
	}

	c.Start()
	select {}
}

type Game struct {
	Beg time.Time
	End time.Time
	c   *cron.Cron
	id  cron.EntryID
}

func New(c *cron.Cron) *Game {
	return &Game{
		Beg: time.Now().Add(10 * time.Second),
		End: time.Now().Add(50 * time.Second),
		c:   c,
	}
}

func (g *Game) SetId(id cron.EntryID) {
	g.id = id
}

func (g *Game) Run() {
	//g.i += 1
	//fmt.Println("run job", g.id)
	// 执行业务前
	if time.Now().Before(g.Beg) {
		fmt.Println("时间未到")
		return
	}

	g.send()

	// 执行业务后
	if time.Now().After(g.End) {
		fmt.Println("remove job", g.id)
		g.c.Remove(g.id)
	}
}

func (g *Game) send() {
	fmt.Println("Game send message id", g.id)
}
