package bytecode

func ReadByteCode(bytecode []byte) []Instr {
    var instrs []Instr
    idx := 0
    for idx < len(bytecode) {
        opcode := Opcode(bytecode[idx])
        opLength := 0
        instr := Instr{ Opcode: opcode }
        switch opcode {
            case Bipush:
                opLength = 1
            case Sipush:
                opLength = 2
            case Ldc:
                opLength = 1
            case LdcW:
                opLength = 2
            case Ldc2W:
                opLength = 2 
            case Iload:
                opLength = 1
            case Lload:
                opLength = 1
            case Dload:
                opLength = 1
            case Fload:
                opLength = 1
            case Aload:
                opLength = 1
            case Istore:
                opLength = 1
            case Lstore:
                opLength = 1
            case Fstore:
                opLength = 1
            case Dstore:
                opLength = 1
            case Astore:
                opLength = 1
            case Iinc:
                opLength = 2
            case Ifeq:
                opLength = 2
            case Ifne:
                opLength = 2
            case Iflt:
                opLength = 2
            case Ifgt:
                opLength = 2
            case Ifle:
                opLength = 2
            case Ifge:
                opLength = 2
            case IfIcmpeq:
                opLength = 2
            case IfIcmpne:
                opLength = 2
            case IfIcmplt:
                opLength = 2
            case IfIcmple:
                opLength = 2
            case IfIcmpgt:
                opLength = 2
            case IfIcmpge:
                opLength = 2
            case Goto:
                opLength = 2
            case Jsr:
                opLength = 2
            case Ret:
                opLength = 1
            case Tableswitch:
                panic("Tableswitch not implemented")
            case Lookupswitch:
                panic("Lookupswitch not implemented")
            case Getstatic:
                opLength = 2
            case Putstatic:
                opLength = 2
            case Getfield:
                opLength = 2
            case Putfield:
                opLength = 2
            case Invokevirtual:
                opLength = 2
            case Invokespecial:
                opLength = 2
            case Invokestatic:
                opLength = 2
            case Invokeinterface:
                opLength = 4
            case Invokedynamic:
                opLength = 4
            case New:
                opLength = 2
            case Newarray:
                opLength = 1
            case Anewarray:
                opLength = 2
            case Checkcast:
                opLength = 2
            case Instanceof:
                opLength = 2
            case Wide:
                panic("Wide opcode not supported yet")
            case Multianewarray:
                opLength = 3
            case Ifnull:
                opLength = 2
            case Ifnonnull:
                opLength = 2
            case Gotow:
                opLength = 4
            case Jsrw:
                opLength = 4
        }
        idx++
        if opLength > 0 {
            instr.Operands = bytecode[idx:idx + opLength]
        }
        instrs = append(instrs, instr)
        idx += opLength
    }
    return instrs
}

