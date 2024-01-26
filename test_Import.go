package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

// isBuiltinType checks if the given type is a built-in Go type
func isBuiltinType(typeName string) bool {
	switch typeName {
	case "bool", "byte", "complex64", "complex128", "error", "float32", "float64",
		"int", "int8", "int16", "int32", "int64", "rune", "string", "uint", "uint8",
		"uint16", "uint32", "uint64", "uintptr", "main", "panic", "root":
		return true
	default:
		return false
	}
}

// 辅助函数：获取类型名称
func getTypeName(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s.%s", t.X, t.Sel.Name)
	case *ast.StarExpr:
		return "*" + getTypeName(t.X)
	case *ast.ArrayType:
		return "[]" + getTypeName(t.Elt)
	case *ast.MapType:
		return "map[" + getTypeName(t.Key) + "]" + getTypeName(t.Value)
	// 添加其他类型的处理逻辑，如结构体等
	default:
		return "UnknownType"
	}
}

func main() {
	src := `
package main

func testHelloWorld(root *TreeNode) []Person {
    panic("implement me")
}
`
	// 解析代码
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	// 遍历所有函数声明，检查是否被调用
	for _, decl := range node.Decls {
		if funcDecl, ok := decl.(*ast.FuncDecl); ok {
			//funcName := funcDecl.Name.Name
			// 输出函数或方法的名称
			fmt.Printf("Function/Method Name: %s\n", funcDecl.Name.Name)

			// 检查函数或方法是否有参数
			if funcDecl.Type.Params != nil {
				// 遍历参数列表
				for _, field := range funcDecl.Type.Params.List {
					// 输出参数类型
					for _, paramName := range field.Names {
						// 获取参数的类型名称
						typeName := getTypeName(field.Type)

						fmt.Printf("Parameter Name: %s, Type: %s\n", paramName.Name, typeName)
					}
				}
			} else {
				fmt.Println("No parameters for this function/method.")
			}
			fmt.Println()
		}
	}

}
