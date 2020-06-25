package main

import (
    "flag"
    "fmt"
    "github.com/iancoleman/strcase"
    "strings"
)

func main() {
    n := flag.String("n", "StructName", "Specification struct name")
    fs := flag.String("f", "field1 field2", "Fields")
    flag.Parse()
    farr := strings.Split(*fs, " ")
    var fields []string
    for _, f := range farr {
        c := strcase.ToCamel(f)
        fields = append(fields, fmt.Sprintf("%s     string    `json:\"%s\"`",  c, f))
    }
    sfs := strings.Join(fields, "\n")
    format := `type %s struct {
%s
}`
    f := fmt.Sprintf(format, strcase.ToCamel(*n), sfs)
    fmt.Println(f)
}
