package bytecode

import (
	"bytes"
)

func Write(instrs []Instr) []byte {
    buf := bytes.Buffer{}

    for _, instr := range instrs {
       buf.WriteByte(byte(instr.Opcode)) 
       if len(instr.Operands) > 0 {
           buf.Write(instr.Operands)
       }
    }
    return buf.Bytes()
}
