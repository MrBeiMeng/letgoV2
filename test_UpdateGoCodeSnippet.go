package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"strings"
)

//
//func main() {
//
//	funcContent := `package main
//
//func lengthOfLongestSubstring(s string) int {
//}`
//
//	// 解析代码字符串
//	fset := token.NewFileSet()
//	node, err := parser.ParseFile(fset, "", funcContent, parser.ParseComments)
//	if err != nil {
//		panic(err)
//	}
//
//	// 遍历 AST
//	ast.Inspect(node, func(n ast.Node) bool {
//		if funDecl, ok := n.(*ast.FuncDecl); ok {
//			// 找到目标函数
//			if funDecl.Name.Name == "lengthOfLongestSubstring" {
//				// 检查函数是否有返回值
//				if funDecl.Type.Results != nil || len(funDecl.Type.Results.List) != 0 {
//					// 在函数末尾添加 panic("implement me")
//					stmt := &ast.ExprStmt{
//						X: &ast.CallExpr{
//							Fun: &ast.Ident{Name: "panic"},
//							Args: []ast.Expr{
//								&ast.BasicLit{Kind: token.STRING, Value: "\"implement me\""},
//							},
//						},
//					}
//					funDecl.Body.List = append(funDecl.Body.List, stmt)
//
//					// 打印修改后的代码
//					//printer.Fprint(os.Stdout, fset, node)
//					return false
//				}
//			}
//		}
//		return true
//	})
//
//	fmt.Println("\n\nModified Code:")
//	printer.Fprint(os.Stdout, fset, node)
//}

func main() {
	funcContent := `package main

func lengthOfLongestSubstring(s string) int {
}

func someFunction() {
	// This function has a return value
	return 42
}

func anotherFunction() {
	// This function has no return value
	fmt.Println("Hello")
}`

	// 解析代码字符串
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "", funcContent, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	// 遍历 AST
	ast.Inspect(node, func(n ast.Node) bool {
		if funDecl, ok := n.(*ast.FuncDecl); ok {
			// 检查函数是否有返回值
			if funDecl.Type.Results != nil && len(funDecl.Type.Results.List) > 0 {
				// 在函数末尾添加 panic("implement me")
				stmt := &ast.ExprStmt{
					X: &ast.CallExpr{
						Fun: &ast.Ident{Name: "panic"},
						Args: []ast.Expr{
							&ast.BasicLit{Kind: token.STRING, Value: "\"implement me\""},
						},
					},
				}
				funDecl.Body.List = append(funDecl.Body.List, stmt)
			}
		}
		return true
	})

	// 获取修改后的代码字符串
	// 使用 strings.Builder 构建字符串
	var modifiedCode strings.Builder
	// 打印修改后的代码
	//fmt.Println("Modified Code:")
	//printer.Fprint(os.Stdout, fset, node)
	printer.Fprint(&modifiedCode, fset, node)

	fmt.Printf("Modified Code: %s", modifiedCode.String())

}
