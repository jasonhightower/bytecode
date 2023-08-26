package main

import (
	"flag"
	"fmt"
	"io"
	"os"
    "bytes"
    "github.com/jasonhightower/bytecode/shared"
)

func main() {
    classFile := flag.String("f", "", "Class file to read")
    printUsage := flag.Bool("h", false, "Help")

    flag.Parse()

    if *printUsage {
        flag.Usage()
        return
    }

    f, err := os.Open(*classFile)
    checkErr(err)
    defer f.Close()

    var r io.Reader
    r = f

    class, err := ReadClass(&r)
    if err != nil {
        fmt.Printf("%s\n", err)
        os.Exit(1)
    }
    fmt.Printf("%d:%d\n", class.Major, class.Minor)


    fmt.Printf("Java Class Version: %d.%d\n", class.Major, class.Minor)
    fmt.Printf("  IsPublic: %t\n", class.Flags.IsPublic())
    fmt.Printf("  this: #%d\n", class.ThisIndex)
    fmt.Printf("  super: #%d\n", class.SuperIndex)
    fmt.Printf("  interfaces: %d\n", len(class.Interfaces))
    fmt.Printf("  fields: %d\n", len(class.Fields))
    fmt.Printf("  methods: %d\n", len(class.Methods))
    for i := 0; i < len(class.Methods); i++ {
        fmt.Printf("    #%d\n", class.Methods[i].NameIndex)
        for j := 0; j < len(class.Methods[i].Attributes); j++ {
            fmt.Printf("      #%d\n", class.Methods[i].Attributes[j].NameIndex)
            cd := class.ConstantPool.Constants[class.Methods[i].Attributes[j].NameIndex - 1]
            var utf8 ConstUtf8
            utf8 = cd.(ConstUtf8)
            fmt.Printf("     %s\n", string(utf8.Data)) 
            if "Code" == string(utf8.Data) {
                var ior io.Reader
                ior = bytes.NewReader(class.Methods[i].Attributes[j].Info)
                var ca Code
                ReadCode(&ior, &ca)
//                readCodeAttribute(&ior, &ca)
                fmt.Println("       code:")    
                instrs:= ReadByteCode(ca.ByteCode)
                for k := 0; k < len(instrs); k++ {
                    fmt.Println(instrs[k])
                }
//                readBytecode(ca.Code)
            }        
        }
    }
    fmt.Printf("  attributes: %d\n", len(class.Attributes))
    fmt.Printf("  Constant pool count: %d\n", class.ConstantPool.Count())
    for i := 1; i < int(class.ConstantPool.Count()); i++ {
        c := class.ConstantPool.Get(CpIndex(i))
        fmt.Printf("   %d | %s\n", i, *c)
    }

    fmt.Println()
    fmt.Println()

    var out io.Writer
    out = os.Stdout

    fmt.Println("====================")
    Write(&out, class)
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
