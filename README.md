# sysy-compiler

杭电编译原理实验，antlr-go 实现的 sysy 编译器

## 词法分析&语法分析

已完工，词法语法文件和生成脚本均位于 antlr 文件夹，但是依赖我这本地的 python 环境，如果要重新生成就需要重新配置一下环境

## 语义分析

施工中,目前缺少stmt的语义分析，其他部分已经完成

## 代码生成

未开工，预计采用 LLVM IR 作为中间代码，后续目标代码就靠LLVM了