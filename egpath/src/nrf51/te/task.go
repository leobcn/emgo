package te

import (
	"mmio"

	"nrf51/internal"
)

// Task represents nRF51 task.
type Task struct {
	reg *mmio.U32
}

// GetTask returns n-th te task.
func GetTask(pe *internal.Pheader, n int) Task {
	return Task{&pe.Tasks[n]}
}

// Trigger starts action corresponding to task t.
func (t Task) Trigger() {
	t.reg.Store(1)
}
