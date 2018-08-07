package util

import "sync"

type GracefulMonitor interface {
	StartRoutine()
	FinishRoutine()
	Wait()
}

func NewGraceful() GracefulMonitor {
	return &Graceful{}
}

type Graceful struct {
	rwMux         sync.RWMutex
	wg            sync.WaitGroup
	routinesCount int
}

func (g *Graceful) StartRoutine() {
	g.rwMux.Lock()
	defer g.rwMux.Unlock()
	g.wg.Add(1)
	g.routinesCount++
}

func (g *Graceful) FinishRoutine() {
	g.rwMux.Lock()
	defer g.rwMux.Unlock()
	g.wg.Done()
	g.routinesCount--
}

func (g *Graceful) RoutinesCount() int {
	g.rwMux.RLock()
	defer g.rwMux.RUnlock()
	return g.routinesCount
}

func (g *Graceful) Wait() {
	g.wg.Wait()
}
