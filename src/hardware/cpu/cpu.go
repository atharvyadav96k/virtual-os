package cpu

import (
	"fmt"

	"github.com/atharvyadav96k/virtual-os/hardware/cpu/flags"
	"github.com/atharvyadav96k/virtual-os/hardware/ram"
	"github.com/atharvyadav96k/virtual-os/hardware/storage/register"
	"github.com/atharvyadav96k/virtual-os/hardware/storage/values"
	"github.com/atharvyadav96k/virtual-os/opcode/parser"
)

func NewCpu() CPU {
	return CPU{
		general_purpose_register: register.NewRegisters(7),
		programCounter:           0,
		stackPointer:             0,
		flags:                    flags.NewFlags(),
	}
}

func (c *CPU) SetGeneralPurposeRegisterValue(idx int, value values.Value) error {
	if idx < 0 || idx >= len(c.general_purpose_register) {
		return fmt.Errorf("register R%d out of range", idx)
	}
	c.general_purpose_register[idx].SetValue(value)
	return nil
}

func (c *CPU) GetGeneralPurposeRegisterValue(idx int) values.Value {
	if idx < 0 || idx >= len(c.general_purpose_register) {
		return values.NewNull()
	}
	return c.general_purpose_register[idx].GetValue()
}

func (c *CPU) IncrementProgramCounter() {
	c.programCounter++
}

func (c *CPU) JumpProgramCounter(jump int) {
	c.programCounter += jump
}

// LoadProgram writes each byte of the program into RAM as Character values starting at startAddr.
func (c *CPU) LoadProgram(r *ram.Ram, startAddr int, program []byte) {
	for i, b := range program {
		r.Write(startAddr+i, values.NewCharacter(b))
	}
}

func (c *CPU) Step(r *ram.Ram) (bool, error) {
	var buf []byte
	for {
		if c.programCounter < 0 || c.programCounter >= r.Size() {
			return false, fmt.Errorf("program counter %d out of RAM bounds", c.programCounter)
		}
		cell := r.Read(c.programCounter)
		c.programCounter++
		b, err := cell.GetCharacter()
		if err != nil {
			return false, fmt.Errorf("PC %d: expected byte, got %v", c.programCounter-1, cell.Type)
		}
		if b == '\n' || b == 0 {
			break
		}
		buf = append(buf, b)
	}

	instr := string(buf)
	code, err := parser.Parse(instr, c.general_purpose_register, r)
	if err != nil {
		return false, fmt.Errorf("parse error: %w", err)
	}

	if code == nil {
		return false, nil
	}

	c.instruction = code
	code.Execute()
	return true, nil
}
