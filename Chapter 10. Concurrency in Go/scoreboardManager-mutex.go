package main

import (
	"fmt"
	"sync"
)

func main() {
	manager := NewMutexScoreboardManager()
	manager.Update("hello", 1)
	fmt.Println(manager.Read("hello"))
	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			manager.Inc1("hello")
		}()
	}
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			manager.Inc2("hello")
		}()
	}
	wg.Wait()
	fmt.Println(manager.Read("hello"))
}

type MutexScoreboardManager struct {
	l          sync.RWMutex
	scoreboard map[string]int
}

func NewMutexScoreboardManager() *MutexScoreboardManager {
	return &MutexScoreboardManager{scoreboard: map[string]int{}}
}

func (msm *MutexScoreboardManager) Update(name string, val int) {
	msm.l.Lock()
	defer msm.l.Unlock()
	msm.scoreboard[name] = val
}
func (msm *MutexScoreboardManager) Read(name string) (int, bool) {
	msm.l.RLock()
	defer msm.l.RUnlock()
	i, ok := msm.scoreboard[name]
	return i, ok
}
func (msm *MutexScoreboardManager) Inc1(name string) {
	msm.l.Lock()
	defer msm.l.Unlock()
	msm.scoreboard[name]++
}
func (msm *MutexScoreboardManager) Inc2(name string) {
	msm.l.Lock()
	defer msm.l.Unlock()
	msm.scoreboard[name]++
}
