package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/pkg/errors"
)

const writeFileName = "consulConfig"
const suffix = "_gen.go"
const filePerm = 0644 // Unix permission bits

// tpl 生成代码需要用到模板
const tpl = `
// Code generated by launch-base-generate DO NOT EDIT
package {{.pkg}}

// ConsulConfig ...
type ConsulConfig struct {
	{{range $key, $value := .comments}}
	{{$key}} string ` + "`" + `consul:"{{$value}}"` + "`" + `
	{{end}}
}

// NewDefaultConsulConfig ...
func NewDefaultConsulConfig() *ConsulConfig {
	ret := &ConsulConfig{}

	return ret
}
`

func main() {
	fset := token.NewFileSet()
	curPath, _ := os.Getwd()
	filePath := fmt.Sprintf("%s/%s", curPath, "consulConstant.go")
	af, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	cmap := ast.NewCommentMap(fset, af, af.Comments)

	// 保存注释信息
	var comments = make(map[string]string)
	for node := range cmap {
		// 仅支持一条声明语句，一个常量的情况
		if spec, ok := node.(*ast.ValueSpec); ok && len(spec.Names) == 1 {
			// 仅提取常量的注释
			ident := spec.Names[0]
			if ident.Obj.Kind == ast.Con {
				// 获取注释信息
				comments[ident.Name] = getComment(ident.Name, spec.Doc)
			}
		}
	}

	ctxBytes, err := gen(comments)
	writeFilePath := fmt.Sprintf("%s/%s%s", curPath, writeFileName, suffix)

	_ = ioutil.WriteFile(writeFilePath, ctxBytes, filePerm)
}

// getComment 获取注释信息，来自AST标准库的summary方法
func getComment(name string, group *ast.CommentGroup) string {
	var buf bytes.Buffer

	for _, comment := range group.List {
		// 注释信息会以 // 参数名，开始，我们实际使用时不需要，去掉
		text := strings.TrimSpace(strings.TrimPrefix(comment.Text, fmt.Sprintf("// %s", name)))
		text = strings.Split(text, " ")[0]
		buf.WriteString(text)
	}

	// replace any invisible with blanks
	bytes := buf.Bytes()
	for i, b := range bytes {
		switch b {
		case '\t', '\n', '\r':
			bytes[i] = ' '
		}
	}

	return string(bytes)
}

// gen 生成代码
func gen(comments map[string]string) ([]byte, error) {
	var buf = bytes.NewBufferString("")

	data := map[string]interface{}{
		"pkg":      "api",
		"comments": comments,
	}

	t, err := template.New("").Parse(tpl)
	if err != nil {
		return nil, errors.Wrapf(err, "template init err")
	}

	err = t.Execute(buf, data)
	if err != nil {
		return nil, errors.Wrapf(err, "template data err")
	}

	return format.Source(buf.Bytes())
}
