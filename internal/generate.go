package internal

import (
	"errors"
	"flag"
	"path/filepath"

	"github.com/iancoleman/strcase"

	"github.com/thecodedproject/gopkg"

	"fmt"
)

var (
	structName = flag.String("struct_name", "", "name of the resource struct")
)

func Generate() error {

	flag.Parse()

	if *structName == "" {
		return errors.New("struct_name must be set")
	}

	outputDir := "."

	pkgFiles, err := gopkg.Parse(outputDir)
	if err != nil {
		return err
	}

	resourceStruct, pkgName, _ := findResourceStruct(pkgFiles, *structName)

	resourceStructType, ok := resourceStruct.Type.(gopkg.TypeStruct)
	if !ok {
		return errors.New(*structName + "is not a struct")
	}

	return gopkg.LintAndGenerate([]gopkg.FileContents{
		makeImplFile(outputDir, resourceStructType, *structName, pkgName),
	})
}

func findResourceStruct(
	pkgFiles []gopkg.FileContents,
	structName string,
) (gopkg.DeclType, string, string) {


	for _, f := range pkgFiles {
		for _, t := range f.Types {
			if t.Name == structName {
				return t, f.PackageName, f.PackageImportPath
			}
		}
	}

	return gopkg.DeclType{}, "", ""
}

func makeImplFile(
	outputDir string,
	resourceStruct gopkg.TypeStruct,
	structName string,
	pkgName string,
) gopkg.FileContents {

	funcs := make([]gopkg.DeclFunc, 0, len(resourceStruct.Fields))

	fmt.Println(resourceStruct)

	for _, field := range resourceStruct.Fields {

		fmt.Println("field:", field)

		funcs = append(funcs, gopkg.DeclFunc{
			Name: "Get" + strcase.ToCamel(field.Name),
			ReturnArgs: []gopkg.DeclVar{
				{
					Type: field.Type,
				},
			},
		})
	}

	file := gopkg.FileContents{
		PackageName: pkgName,
		Filepath: filepath.Join(outputDir, strcase.ToSnake(structName) + "_impl.go"),
		Types: []gopkg.DeclType{
			{
				Name: strcase.ToCamel(structName),
				Type: gopkg.TypeInterface{
					Funcs: funcs,
				},
			},
		},
	}


	return file
}
