package ast

import (
	"encoding/json"
	"fmt"
	"os"
	"sysy/ast/types"
	"testing"
)

func TestDraw(t *testing.T) {
	prog := Program{
		Children: []any{
			&Declare{
				Definitions: []Define{
					{
						Type:         &types.Base{Type: types.Int},
						Identifier:   "a",
						InitialValue: nil,
					},
				},
			},
			&FuncDef{
				Identifier: "main",
				Params: []*FuncParam{
					{
						Identifier: "a",
						Type:       &types.Base{Type: types.Int},
					}, {
						Identifier: "b",
						Type:       &types.Base{Type: types.Float},
					},
				},
				Return: &types.Base{Type: types.Int},
				Body:   nil,
			},
		},
	}

	tree := prog.Draw()
	str, err := json.Marshal(tree)
	if err != nil {
		t.Error(err)
		return
	}
	os.WriteFile("test.html", []byte(fmt.Sprintf(tmpl, str)), 0666)
}

const tmpl = `<!DOCTYPE html>
<html>
<head>
    <!-- Import style -->
    <link rel="stylesheet" href="//unpkg.com/element-plus/dist/index.css" />
    <!-- Import Vue 3 -->
    <script src="https://unpkg.com/vue@3/dist/vue.global.js"></script>
    <!-- Import component library -->
    <script src="//unpkg.com/element-plus"></script>
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
