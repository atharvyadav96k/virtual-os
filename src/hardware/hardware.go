package hardware

import (
	"github.com/atharvyadav96k/virtual-os/hardware/cpu"
	"github.com/atharvyadav96k/virtual-os/hardware/ram"
)

type Hardware struct {
	cpu cpu.CPU
	ram ram.Ram
}

func NewHardware() Hardware {
	return Hardware{
		cpu: cpu.NewCpu(),
		ram: ram.NewRam(1024),
	}
}

func (h *Hardware) Ram() *ram.Ram {
	return &h.ram
}

func (h *Hardware) Cpu() *cpu.CPU {
	return &h.cpu
}

func (h *Hardware) LoadProgram(startAddr int, program []byte) {
	h.cpu.LoadProgram(&h.ram, startAddr, program)
}

func (h *Hardware) Run() error {
	for {
		cont, err := h.cpu.Step(&h.ram)
		if err != nil {
			return err
		}
		if !cont {
			return nil
		}
	}
}
