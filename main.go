package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"log"
	"os"
	"sysy/parser"
	"sysy/visitor"
)

var (
	code   string
	output string
)

func main() {
	flag.StringVar(&code, "f", "./test/demo1.sysy", "code file path")
	flag.StringVar(&output, "o", "./output.html", "output file path")
	flag.Parse()

	input, err := antlr.NewFileStream(code)
	if err != nil {
		log.Fatalln(err)
	}

	lexer := parser.NewsysyLexer(input)
	if lexer.HasError() {
		log.Fatalln(lexer.GetError().GetMessage())
	}

	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewsysyParser(stream)
	if p.HasError() {
		log.Fatalln(p.GetError().GetMessage())
	}

	program, err := visitor.NewContext().Comp(p.Comp())
	if err != nil {
		log.Fatalln(err)
	}

	str, err := json.Marshal(program.Draw())
	if err != nil {
		log.Fatalln(err)
	}

	os.WriteFile("output.html", []byte(fmt.Sprintf(tmpl, str)), 0666)
}

const tmpl = `<!DOCTYPE html>
<html>
<head>
    <!-- Import style -->
    <link rel="stylesheet" href="https://unpkg.com/element-plus/dist/index.css" />
    <!-- Import Vue 3 -->
    <script src="https://unpkg.com/vue@3/dist/vue.global.js"></script>
    <!-- Import component library -->
    <script src="https://unpkg.com/element-plus"></script>
</head>
<body>
    <div id="app">
        <el-tree :data="data" :props="defaultProps" @node-click="handleNodeClick" />
    </div>
    <script>
        const { createApp, reactive } = Vue

        createApp({
            setup() {
                const data = reactive([%s])
                return {
                    data
                }
            }
        }).use(ElementPlus).mount('#app')
    </script>
</body>
</html>`
