package bytecode

import (
	"fmt"
	"strconv"
)

type Instr struct {
    Opcode
    Operands []byte
}
func (inst Instr) String() string {
    base := fmt.Sprintf("%s", inst.Opcode)
    for _, operand := range inst.Operands {
        base += " "
        base += strconv.FormatUint(uint64(operand), 10)
    }
    return base
}

type Opcode byte
const (
    Nop Opcode = 0x00
    AconstNull Opcode = 0x01
    IconstM1 Opcode = 0x02
    Iconst0 Opcode = 0x03
    Iconst1 Opcode = 0x04
    Iconst2 Opcode = 0x05
    Iconst3 Opcode = 0x06
    Iconst4 Opcode = 0x07
    Iconst5 Opcode = 0x08
    Lconst0 Opcode = 0x09
    Lconst1 Opcode = 0x0a
    Fconst0 Opcode = 0x0b
    Fconst1 Opcode = 0x0c
    Fconst2 Opcode = 0x0d
    Dconst0 Opcode = 0x0e
    Dconst1 Opcode = 0x0f
    Bipush Opcode = 0x10
    Sipush Opcode = 0x11
    Ldc Opcode = 0x12
    LdcW Opcode = 0x13
    Ldc2W Opcode = 0x14
    Iload Opcode = 0x15
    Lload Opcode = 0x16
    Fload Opcode = 0x17
    Dload Opcode = 0x18
    Aload Opcode = 0x19
    Iload0 Opcode = 0x1a
    Iload1 Opcode = 0x1b
    Iload2 Opcode = 0x1c
    Iload3 Opcode = 0x1d
    Lload0 Opcode = 0x1e
    Lload1 Opcode = 0x1f
    Lload2 Opcode = 0x20
    Lload3 Opcode = 0x21
    Fload0 Opcode = 0x22
    Fload1 Opcode = 0x23
    Fload2 Opcode = 0x24
    Fload3 Opcode = 0x25
    Dload0 Opcode = 0x26
    Dload1 Opcode = 0x27
    Dload2 Opcode = 0x28
    Dload3 Opcode = 0x29
    Aload0 Opcode = 0x2a
    Aload1 Opcode = 0x2b
    Aload2 Opcode = 0x2c
    Aload3 Opcode = 0x2d
    Iaload Opcode = 0x2e
    Laload Opcode = 0x2f
    Faload Opcode = 0x30
    Daload Opcode = 0x31
    Aaload Opcode = 0x32
    Baload Opcode = 0x33
    Caload Opcode = 0x34
    Saload Opcode = 0x35
    Istore Opcode = 0x36
    Lstore Opcode = 0x37
    Fstore Opcode = 0x38
    Dstore Opcode = 0x39
    Astore Opcode = 0x3a
    Istore0 Opcode = 0x3b
    Istore1 Opcode = 0x3c
    Istore2 Opcode = 0x3d
    Istore3 Opcode = 0x3e
    Lstore0 Opcode = 0x3f
    Lstore1 Opcode = 0x40
    Lstore2 Opcode = 0x41
    Lstore3 Opcode = 0x42
    Fstore0 Opcode = 0x43
    Fstore1 Opcode = 0x44
    Fstore2 Opcode = 0x45
    Fstore3 Opcode = 0x46
    Dstore0 Opcode = 0x47
    Dstore1 Opcode = 0x48
    Dstore2 Opcode = 0x49
    Dstore3 Opcode = 0x4a
    Astore0 Opcode = 0x4b
    Astore1 Opcode = 0x4c
    Astore2 Opcode = 0x4d
    Astore3 Opcode = 0x4e
    Iastore Opcode = 0x4f
    Lastore Opcode = 0x50
    Fastore Opcode = 0x51
    Dastore Opcode = 0x52
    Aastore Opcode = 0x53
    Bastore Opcode = 0x54
    Castore Opcode = 0x55
    Sastore Opcode = 0x56
    Pop Opcode = 0x57
    Pop2 Opcode = 0x58
    Dup Opcode = 0x59
    DupX1 Opcode = 0x5a
    DupX2 Opcode = 0x5b
    Dup2 Opcode = 0x5c
    Dup2X1 Opcode = 0x5d
    Dup2X2 Opcode = 0x5e
    Swap Opcode = 0x5f
    Iadd Opcode = 0x60
    Ladd Opcode = 0x61
    Fadd Opcode = 0x62
    Dadd Opcode = 0x63
    Isub Opcode = 0x64
    Lsub Opcode = 0x65
    Fsub Opcode = 0x66
    Dsub Opcode = 0x67
    Imul Opcode = 0x68
    Lmul Opcode = 0x69
    Fmul Opcode = 0x6a
    Dmul Opcode = 0x6b
    Idiv Opcode = 0x6c
    Ldiv Opcode = 0x6d
    Fdiv Opcode = 0x6e
    Ddiv Opcode = 0x6f
    Irem Opcode = 0x70
    Lrem Opcode = 0x71
    Frem Opcode = 0x72
    Drem Opcode = 0x73
    Ineg Opcode = 0x74
    Lneg Opcode = 0x75
    Fneg Opcode = 0x76
    Dneg Opcode = 0x77
    Ishl Opcode = 0x78
    Lshl Opcode = 0x79
    Ishr Opcode = 0x7a
    Lshr Opcode = 0x7b
    Iushr Opcode = 0x7c
    Lushr Opcode = 0x7d
    Iand Opcode = 0x7e
    Land Opcode = 0x7f
    Ior Opcode = 0x80
    Lor Opcode = 0x81
    Ixor Opcode = 0x82
    Lxor Opcode = 0x83
    Iinc Opcode = 0x84
    I2l Opcode = 0x85
    I2f Opcode = 0x86
    I2d Opcode = 0x87
    L2i Opcode = 0x88
    L2f Opcode = 0x89
    L2d Opcode = 0x8a
    f2i Opcode = 0x8b
    f2l Opcode = 0x8c
    f2d Opcode = 0x8d
    d2i Opcode = 0x8e
    d2l Opcode = 0x8f
    d2f Opcode = 0x90
    i2b Opcode = 0x91
    i2c Opcode = 0x92
    i2s Opcode = 0x93
    lcmp Opcode = 0x94
    fcmpl Opcode = 0x95
    fcmpg Opcode = 0x96
    dcmpl Opcode = 0x97
    dcmpg Opcode = 0x98
    Ifeq Opcode = 0x99
    Ifne Opcode = 0x9a
    Iflt Opcode = 0x9b
    Ifge Opcode = 0x9c
    Ifgt Opcode = 0x9d
    Ifle Opcode = 0x9e
    IfIcmpeq Opcode = 0x9f
    IfIcmpne Opcode = 0xa0
    IfIcmplt Opcode = 0xa1
    IfIcmpge Opcode = 0xa2
    IfIcmpgt Opcode = 0xa3
    IfIcmple Opcode = 0xa4
    Ifacmpeq Opcode = 0xa5
    Ifacmpne Opcode = 0xa6
    Goto Opcode = 0xa7
    Jsr Opcode = 0xa8
    Ret Opcode = 0xa9
    Tableswitch Opcode = 0xaa
    Lookupswitch Opcode = 0xab
    Ireturn Opcode = 0xac
    Lreturn Opcode = 0xad
    Freturn Opcode = 0xae
    Deturn Opcode = 0xaf
    Areturn Opcode = 0xb0
    Return Opcode = 0xb1
    Getstatic Opcode = 0xb2
    Putstatic Opcode = 0xb3
    Getfield Opcode = 0xb4
    Putfield Opcode = 0xb5
    Invokevirtual Opcode = 0xb6
    Invokespecial Opcode = 0xb7
    Invokestatic Opcode = 0xb8
    Invokeinterface Opcode = 0xb9
    Invokedynamic Opcode = 0xba
    New Opcode = 0xbb
    Newarray Opcode = 0xbc
    Anewarray Opcode = 0xbd
    Arraylength Opcode = 0xbe
    Athrow Opcode = 0xbf
    Checkcast Opcode = 0xc0
    Instanceof Opcode = 0xc1
    Monitorenter Opcode = 0xc2
    Monitorexit Opcode = 0xc3
    Wide Opcode = 0xc4
    Multianewarray Opcode = 0xc5
    Ifnull Opcode = 0xc6
    Ifnonnull Opcode = 0xc7
    Gotow Opcode = 0xc8
    Jsrw Opcode = 0xc9
    Breakpoint Opcode = 0xca
    Impdep1 Opcode = 0xfe
    Impdep2 Opcode = 0xff
)
func (o Opcode) String() string {
    switch o {
        case Nop: 
            return "Nop"
        case AconstNull: 
            return "aconstnull"
        case IconstM1: 
            return "iconstm1"
        case Iconst0: 
            return "iconst0"
        case Iconst1: 
            return "iconst1"
        case Iconst2: 
            return "iconst2"
        case Iconst3: 
            return "iconst3"
        case Iconst4: 
            return "iconst4"
        case Iconst5: 
            return "iconst5"
        case Lconst0: 
            return "lconst0"
        case Lconst1: 
            return "lconst1"
        case Fconst0: 
            return "fconst0"
        case Fconst1: 
            return "fconst1"
        case Fconst2: 
            return "fconst2"
        case Dconst0: 
            return "dconst0"
        case Dconst1: 
            return "dconst1"
        case Bipush: 
            return "bipush"
        case Sipush: 
            return "sipush"
        case Ldc: 
            return "ldc"
        case LdcW: 
            return "ldcw"
        case Ldc2W: 
            return "ldc2w"
        case Iload: 
            return "iload"
        case Lload: 
            return "lload"
        case Fload: 
            return "fload"
        case Dload: 
            return "dload"
        case Aload: 
            return "aload"
        case Iload0: 
            return "iload0"
        case Iload1: 
            return "iload1"
        case Iload2: 
            return "iload2"
        case Iload3: 
            return "iload3"
        case Lload0: 
            return "lload0"
        case Lload1: 
            return "lload1"
        case Lload2: 
            return "lload2"
        case Lload3: 
            return "lload3"
        case Fload0: 
            return "fload0"
        case Fload1: 
            return "fload1"
        case Fload2: 
            return "fload2"
        case Fload3: 
            return "fload3"
        case Dload0: 
            return "dload0"
        case Dload1: 
            return "dload1"
        case Dload2: 
            return "dload2"
        case Dload3: 
            return "dload3"
        case Aload0: 
            return "aload0"
        case Aload1: 
            return "aload1"
        case Aload2: 
            return "aload2"
        case Aload3: 
            return "aload3"
        case Iaload: 
            return "iaload"
        case Laload: 
            return "laload"
        case Faload: 
            return "faload"
        case Daload: 
            return "daload"
        case Aaload: 
            return "aaload"
        case Baload: 
            return "baload"
        case Caload: 
            return "caload"
        case Saload: 
            return "saload"
        case Istore: 
            return "istore"
        case Lstore: 
            return "lstore"
        case Fstore: 
            return "fstore"
        case Dstore: 
            return "dstore"
        case Astore: 
            return "astore"
        case Istore0: 
            return "istore0"
        case Istore1: 
            return "istore1"
        case Istore2: 
            return "istore2"
        case Istore3: 
            return "istore3"
        case Lstore0: 
            return "lstore0"
        case Lstore1: 
            return "lstore1"
        case Lstore2: 
            return "lstore2"
        case Lstore3: 
            return "lstore3"
        case Fstore0: 
            return "fstore0"
        case Fstore1: 
            return "fstore1"
        case Fstore2: 
            return "fstore2"
        case Fstore3: 
            return "fstore3"
        case Dstore0: 
            return "dstore0"
        case Dstore1: 
            return "dstore1"
        case Dstore2: 
            return "dstore2"
        case Dstore3: 
            return "dstore3"
        case Astore0: 
            return "astore0"
        case Astore1: 
            return "astore1"
        case Astore2: 
            return "astore2"
        case Astore3: 
            return "astore3"
        case Iastore: 
            return "iastore"
        case Lastore: 
            return "lastore"
        case Fastore: 
            return "fastore"
        case Dastore: 
            return "dastore"
        case Aastore: 
            return "aastore"
        case Bastore: 
            return "bastore"
        case Castore: 
            return "castore"
        case Sastore: 
            return "sastore"
        case Pop: 
            return "pop"
        case Pop2: 
            return "pop2"
        case Dup: 
            return "dup"
        case DupX1: 
            return "dupx1"
        case DupX2: 
            return "dupx2"
        case Dup2: 
            return "dup2"
        case Dup2X1: 
            return "dup2x1"
        case Dup2X2: 
            return "dup2x2"
        case Swap: 
            return "swap"
        case Iadd: 
            return "iadd"
        case Ladd: 
            return "ladd"
        case Fadd: 
            return "fadd"
        case Dadd: 
            return "dadd"
        case Isub: 
            return "isub"
        case Lsub: 
            return "lsub"
        case Fsub: 
            return "fsub"
        case Dsub: 
            return "dsub"
        case Imul: 
            return "imul"
        case Lmul: 
            return "lmul"
        case Fmul: 
            return "fmul"
        case Dmul: 
            return "dmul"
        case Idiv: 
            return "idiv"
        case Ldiv: 
            return "ldiv"
        case Fdiv: 
            return "fdiv"
        case Ddiv: 
            return "ddiv"
        case Irem: 
            return "irem"
        case Lrem: 
            return "lrem"
        case Frem: 
            return "frem"
        case Drem: 
            return "drem"
        case Ineg: 
            return "ineg"
        case Lneg: 
            return "lneg"
        case Fneg: 
            return "fneg"
        case Dneg: 
            return "dneg"
        case Ishl: 
            return "ishl"
        case Lshl: 
            return "lshl"
        case Ishr: 
            return "ishr"
        case Lshr: 
            return "lshr"
        case Iushr: 
            return "iushr"
        case Lushr: 
            return "lushr"
        case Iand: 
            return "iand"
        case Land: 
            return "land"
        case Ior: 
            return "ior"
        case Lor: 
            return "lor"
        case Ixor: 
            return "ixor"
        case Lxor: 
            return "lxor"
        case Iinc: 
            return "iinc"
        case I2l: 
            return "i2l"
        case I2f: 
            return "i2f"
        case I2d: 
            return "i2d"
        case L2i: 
            return "l2i"
        case L2f: 
            return "l2f"
        case L2d: 
            return "l2d"
        case f2i: 
            return "f2i"
        case f2l: 
            return "f2l"
        case f2d: 
            return "f2d"
        case d2i: 
            return "d2i"
        case d2l: 
            return "d2l"
        case d2f: 
            return "d2f"
        case i2b: 
            return "i2b"
        case i2c: 
            return "i2c"
        case i2s: 
            return "i2s"
        case lcmp: 
            return "lcmp"
        case fcmpl: 
            return "fcmpl"
        case fcmpg: 
            return "fcmpg"
        case dcmpl: 
            return "dcmpl"
        case dcmpg: 
            return "dcmpg"
        case Ifeq: 
            return "ifeq"
        case Ifne: 
            return "ifne"
        case Iflt: 
            return "iflt"
        case Ifge: 
            return "ifge"
        case Ifgt: 
            return "ifgt"
        case Ifle: 
            return "ifle"
        case IfIcmpeq: 
            return "ificmpeq"
        case IfIcmpne: 
            return "ificmpne"
        case IfIcmplt: 
            return "ificmplt"
        case IfIcmpge: 
            return "ificmpge"
        case IfIcmpgt: 
            return "ificmpgt"
        case IfIcmple: 
            return "ificmple"
        case Ifacmpeq: 
            return "ifacmpeq"
        case Ifacmpne: 
            return "ifacmpne"
        case Goto: 
            return "goto"
        case Jsr: 
            return "jsr"
        case Ret: 
            return "ret"
        case Tableswitch: 
            return "tableswitch"
        case Lookupswitch: 
            return "lookupswitch"
        case Ireturn: 
            return "ireturn"
        case Lreturn: 
            return "lreturn"
        case Freturn: 
            return "freturn"
        case Deturn: 
            return "deturn"
        case Areturn: 
            return "areturn"
        case Return: 
            return "return"
        case Getstatic: 
            return "getstatic"
        case Putstatic: 
            return "putstatic"
        case Getfield: 
            return "getfield"
        case Putfield: 
            return "putfield"
        case Invokevirtual: 
            return "invokevirtual"
        case Invokespecial: 
            return "invokespecial"
        case Invokestatic: 
            return "invokestatic"
        case Invokeinterface: 
            return "invokeinterface"
        case Invokedynamic: 
            return "invokedynamic"
        case New: 
            return "new"
        case Newarray: 
            return "newarray"
        case Anewarray: 
            return "anewarray"
        case Arraylength: 
            return "arraylength"
        case Athrow: 
            return "athrow"
        case Checkcast: 
            return "checkcast"
        case Instanceof: 
            return "instanceof"
        case Monitorenter: 
            return "monitorenter"
        case Monitorexit: 
            return "monitorexit"
        case Wide: 
            return "wide"
        case Multianewarray: 
            return "multianewarray"
        case Ifnull: 
            return "ifnull"
        case Ifnonnull: 
            return "ifnonnull"
        case Gotow: 
            return "gotow"
        case Jsrw: 
            return "jsrw"
        case Breakpoint: 
            return "breakpoint"
        case Impdep1: 
            return "impdep1"
        case Impdep2: 
            return "impdep2"
    }
    return strconv.FormatUint(uint64(o), 16)
}
