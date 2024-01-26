package generate_service

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io/fs"
	"letgoV2/system_code/pkg/logging"
	"letgoV2/system_code/pkg/setting"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
)

// 这个util负责计数文件夹，创建文件夹

// ConvZzza2int 将四位 aaaa 是 45_6975 |  1 zzzz 是0 ! 这并非十六进制
func ConvZzza2int(abc string) (error, uint32) {
	if len(abc) != 4 {
		errStr := fmt.Sprintf("字符串长度不合法")
		logging.Info(errStr)

		return errors.New(errStr), 0
	}

	var result uint32
	base := uint32('a') // 'f' 对应的数字减一作为基准 a = 97

	for _, char := range abc {
		if char < 'a' || char > 'z' {
			// 字符不在合法范围内
			errStr := fmt.Sprintf("字符不在合法范围内 [f - z].")
			logging.Info(errStr)
			return errors.New(errStr), 0
		}

		// 将字符映射为数字并累加到结果中
		result = result*26 + (25 - (uint32(char) - base))
	}

	return nil, result
}

// ConvInt2zzza 将四位 1 转为数字 fffa
func ConvInt2zzza(num uint32) (error, string) {
	// 数字范围应该是 0 到 45_6976-1
	if num < 0 || num >= 45_6976 {
		errStr := fmt.Sprintf("数字不合法")
		logging.Info(errStr)

		return errors.New(errStr), ""
	}

	var result strings.Builder
	base := uint32('a')

	for num >= 1 {
		per := num % 26
		num = num / 26
		result.WriteByte(byte(base + (25 - per)))
	}

	length := result.Len()
	for i := 0; i < 4-length; i++ {
		result.WriteByte('z')
	}

	str := result.String()

	result.Reset()
	for i := len(str) - 1; i >= 0; i-- {
		result.WriteByte(str[i])
	}

	return nil, result.String()
}

func GetLastDirID() (error, uint32) {
	targetPath := setting.CodePlace

	codeDirList := make([]string, 0)
	dNameList := make([]string, 0)
	filepath.WalkDir(targetPath, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			return nil
		}

		matched, _ := regexp.Match(`^ID[a-z]{4}_([a-z_-]+)$`, []byte(d.Name()))
		if !matched {
			return nil
		}

		codeDirList = append(codeDirList, path)
		dNameList = append(dNameList, d.Name())
		return nil
	})

	var maxId uint32 = 1
	var lastDir string

	for _, dirName := range dNameList {
		// 定义正则表达式
		re := regexp.MustCompile(`^ID([a-z]{4})_`)

		// 使用正则表达式查找匹配项
		match := re.FindStringSubmatch(dirName)
		if len(match) < 2 {
			continue
		}
		matchStr := match[1]
		err, id := ConvZzza2int(matchStr)
		if err != nil {
			return err, 0
		}

		if id > maxId {
			maxId = id
			lastDir = dirName
		}
	}

	logging.Info(fmt.Sprintf("查询到最后创建的文件夹[%s]ID[%d]", lastDir, maxId))
	return nil, maxId
}

func GetNewDirName(lastId uint32, titleSlug string) (error, string) {
	err, str := ConvInt2zzza(lastId)
	if err != nil {
		return err, str
	}

	titleSlug = strings.ReplaceAll(titleSlug, "-", "_")

	return nil, fmt.Sprintf("ID%s_%s", str, titleSlug)
}

func replaceStructParam(strTemplate string, any interface{}) (error, string) {
	sValue := reflect.ValueOf(any)
	if sValue.Kind() != reflect.Struct {
		errStr := fmt.Sprintf("非Struct类型")
		logging.Error(errStr)
		return errors.New(errStr), ""
	}

	replaceMap := getReplaceMap(sValue)
	for replaceName, value := range replaceMap {
		strTemplate = strings.ReplaceAll(strTemplate, fmt.Sprintf("${%s}", replaceName), value)
	}

	return nil, strTemplate
}

func getReplaceMap(value reflect.Value) map[string]string {
	result := make(map[string]string)
	tV := value.Type()

	for i := 0; i < value.NumField(); i++ {
		fieldValue := value.Field(i)
		fieldType := tV.Field(i)
		tag := fieldType.Tag.Get("replace")
		// Check if the field has a "replace" tag
		if tag != "" {
			result[tag] = fieldValue.String()
		}
	}

	return result
}

// UpdateGoCodeSnippet 更新方法，添加panic 避免编译报错
// 这个方法负责在未编写完成的方法的方法体中添加panic("implement me")避免编译报错
// 正则表达式完善代码
// !!!此方法存在bug @2024年1月25日!!!
func UpdateGoCodeSnippet(codeFunc string) (error, string) {

	// 定义正则表达式来匹配函数体
	//funcRegex := regexp.MustCompile(`func\s+[a-z1-9A-Z_]\w*\s*\([^)]*\)\s*\{([^{]*)}`)
	// 构建动态的正则表达式，匹配任意函数名的函数体
	funcRegex := regexp.MustCompile(`func\s+%s\s*\([^)]*\)\s*\{([^}]*)\}`)

	// 查找匹配的函数体
	match := funcRegex.FindStringSubmatch(codeFunc)
	if len(match) >= 2 {
		// 在函数体最后一行加上 panic("implement me ")
		updatedFunc := fmt.Sprintf("%s\n\t%s", match[1], "panic(\"implement me \")")

		// 替换原始函数体
		updatedCodeFunc := funcRegex.ReplaceAllString(codeFunc, updatedFunc)
		return nil, updatedCodeFunc
	}

	// 如果没有匹配到函数体，返回原始字符串
	return errors.New("未找到函数体"), codeFunc
}

// UpdateGoCodeSnippetByAst 更新方法，添加panic 避免编译报错
// codeContent 要求带上package 字段
func UpdateGoCodeSnippetByAst(codeContent string) (error, string) {
	// 解析代码字符串
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "", codeContent, parser.ParseComments)
	if err != nil {
		return err, ""
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

	return nil, modifiedCode.String()
}

func FindUnusedFunctionNames(code string) (error, []string) {
	// 解析代码
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "", code, parser.ParseComments)
	if err != nil {
		return err, nil
	}

	// 保存函数调用的名称
	calledFunctions := make(map[string]bool)

	// 遍历所有函数调用，并记录调用的函数名
	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.CallExpr:
			switch fun := x.Fun.(type) {
			case *ast.Ident:
				calledFunctions[fun.Name] = true
			case *ast.SelectorExpr:
				calledFunctions[fun.Sel.Name] = true
			}
		}
		return true
	})

	// 保存未被调用的函数名
	unusedFunctions := make([]string, 0)

	// 遍历所有函数声明，检查是否被调用
	for _, decl := range node.Decls {
		if fd, ok := decl.(*ast.FuncDecl); ok {
			funcName := fd.Name.Name
			if !calledFunctions[funcName] {
				unusedFunctions = append(unusedFunctions, funcName)
			}
		}
	}

	return nil, unusedFunctions
}

// isBuiltinType checks if the given type is a built-in Go type
func isBuiltinType(typeName string) bool {
	typeName = strings.ReplaceAll(typeName, "map", "")
	switch strings.Trim(typeName, "[]*") {
	case "bool", "byte", "complex64", "complex128", "error", "float32", "float64",
		"int", "int8", "int16", "int32", "int64", "rune", "string", "uint", "uint8",
		"uint16", "uint32", "uint64", "uintptr":
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

// GetNeedImportPackageList codeContent 要求是一个完整的代码内容 包括package
func GetNeedImportPackageList(codeContent string) (result []string) {
	// 解析代码
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "", codeContent, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	// 遍历所有函数声明，检查是否被调用
	for _, decl := range node.Decls {
		if funcDecl, ok := decl.(*ast.FuncDecl); ok {
			//funcName := funcDecl.Name.Name
			// 输出函数或方法的名称
			logging.Info(fmt.Sprintf("Function/Method Name: %s\n", funcDecl.Name.Name))

			// 检查函数或方法是否有参数
			if funcDecl.Type.Params != nil {
				// 遍历参数列表
				for _, field := range funcDecl.Type.Params.List {
					// 输出参数类型
					for _, paramName := range field.Names {
						// 获取参数的类型名称
						typeName := getTypeName(field.Type)

						logging.Info(fmt.Sprintf("Parameter Name: %s, Type: %s ", paramName.Name, typeName))

						if !isBuiltinType(typeName) {
							result = append(result, typeName)
						}
					}
				}
			} else {
				logging.Warn(fmt.Sprintf("No parameters for this function/method."))
			}
		}
	}

	return
}

// 这个方法负责冲代码中获取一个且仅一个未被调用的函数，用来进行测试代码编写
func getOneFuncName(snippet string) (error, string) {
	completeSnippet := fmt.Sprintf("package main\n\n%s", snippet)
	err, funNameList := FindUnusedFunctionNames(completeSnippet)
	if err != nil {
		return err, ""
	}

	if len(funNameList) == 0 {
		errStr := fmt.Sprintf("未找到任何可调用的函数 代码内容:%s", snippet)
		logging.Error(errStr)
		return errors.New(errStr), ""
	}

	if len(funNameList) > 1 {
		errStr := fmt.Sprintf("太多的未调用函数 代码内容:%s", snippet)
		logging.Error(errStr)
		return errors.New(errStr), ""
	}

	return nil, funNameList[0]
}

// 添加引用
func addImportIfNeed(snippet string) string {
	needImportPackageList := GetNeedImportPackageList(snippet)
	replaceFlag := `import (
	_ "fmt"
)`

	importSnippet := ``
	if len(needImportPackageList) == 0 {
		importSnippet = `
import (
	. "letgoV2/system_code/pkg/common"
)`
		logging.Info(fmt.Sprintf("添加了import字段"))
	}

	snippet = strings.ReplaceAll(snippet, replaceFlag, importSnippet)
	return snippet
}
