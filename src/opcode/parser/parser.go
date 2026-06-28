package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/atharvyadav96k/virtual-os/hardware/ram"
	"github.com/atharvyadav96k/virtual-os/hardware/storage"
	"github.com/atharvyadav96k/virtual-os/hardware/storage/immediate"
	"github.com/atharvyadav96k/virtual-os/hardware/storage/register"
	"github.com/atharvyadav96k/virtual-os/hardware/storage/values"
	"github.com/atharvyadav96k/virtual-os/opcode"
	"github.com/atharvyadav96k/virtual-os/opcode/add"
	"github.com/atharvyadav96k/virtual-os/opcode/load"
	"github.com/atharvyadav96k/virtual-os/opcode/mov"
)

func Parse(instruction string, regs []register.Register, r *ram.Ram) (opcode.Code, error) {
	tokens := strings.Fields(instruction)
	if len(tokens) == 0 {
		return nil, fmt.Errorf("empty instruction")
	}

	switch strings.ToUpper(tokens[0]) {
	case "HALT":
		return nil, nil

	case "MOV":
		if len(tokens) != 3 {
			return nil, fmt.Errorf("MOV requires 2 operands, got %d", len(tokens)-1)
		}
		destIdx, err := parseRegister(tokens[1])
		if err != nil {
			return nil, fmt.Errorf("MOV dest: %w", err)
		}
		if destIdx >= len(regs) {
			return nil, fmt.Errorf("MOV dest: register R%d out of range", destIdx)
		}
		src, err := parseSource(tokens[2], regs)
		if err != nil {
			return nil, fmt.Errorf("MOV src: %w", err)
		}
		op := mov.NewMov(destIdx, &regs[destIdx], src)
		return &op, nil

	case "ADD":
		if len(tokens) != 3 {
			return nil, fmt.Errorf("ADD requires 2 operands, got %d", len(tokens)-1)
		}
		destIdx, err := parseRegister(tokens[1])
		if err != nil {
			return nil, fmt.Errorf("ADD dest: %w", err)
		}
		if destIdx >= len(regs) {
			return nil, fmt.Errorf("ADD dest: register R%d out of range", destIdx)
		}
		src, err := parseSource(tokens[2], regs)
		if err != nil {
			return nil, fmt.Errorf("ADD src: %w", err)
		}
		op := add.NewAdd(destIdx, &regs[destIdx], src)
		return &op, nil

	case "LOAD":
		if len(tokens) != 3 {
			return nil, fmt.Errorf("LOAD requires 2 operands, got %d", len(tokens)-1)
		}
		destIdx, err := parseRegister(tokens[1])
		if err != nil {
			return nil, fmt.Errorf("LOAD dest: %w", err)
		}
		if destIdx >= len(regs) {
			return nil, fmt.Errorf("LOAD dest: register R%d out of range", destIdx)
		}
		addr, err := strconv.Atoi(tokens[2])
		if err != nil {
			return nil, fmt.Errorf("LOAD address must be an integer, got %q", tokens[2])
		}
		op := load.NewLoad(destIdx, &regs[destIdx], addr, r)
		return &op, nil

	default:
		return nil, fmt.Errorf("unknown opcode %q", tokens[0])
	}
}

func parseRegister(token string) (int, error) {
	upper := strings.ToUpper(token)
	if len(upper) < 2 || upper[0] != 'R' {
		return 0, fmt.Errorf("expected register like R0, got %q", token)
	}
	idx, err := strconv.Atoi(upper[1:])
	if err != nil {
		return 0, fmt.Errorf("invalid register %q", token)
	}
	return idx, nil
}

func parseSource(token string, regs []register.Register) (storage.Storage, error) {
	if strings.HasPrefix(token, "#") {
		n, err := strconv.Atoi(token[1:])
		if err != nil {
			return nil, fmt.Errorf("invalid immediate %q", token)
		}
		imm := immediate.NewImmediate(values.Integer)
		imm.SetValue(values.NewInt(n))
		return &imm, nil
	}

	idx, err := parseRegister(token)
	if err != nil {
		return nil, err
	}
	if idx >= len(regs) {
		return nil, fmt.Errorf("register R%d out of range", idx)
	}
	return &regs[idx], nil
}
