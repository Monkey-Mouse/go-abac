# go-abac

翻译: [English](README.md)|[简体中文](README_zh.md)



go语言实现基于属性的访问控制（Attribute Based Access Control）的模块，
## 特性
- 自定义访问规则，易于扩展
- 最小依赖，无需外部模块

- 级联方法，便于调用api
- 与不同语言同类项目风格一致
- 简单易用，文档详细

主要参考node.js实现此功能的开源项目
[accessControl](https://github.com/onury/accesscontrol) ，最大的不同在于遵循abac的基本特性，
泛化请求者的属性多样

## 安装

```
go get github.com/Monkey-Mouse/go-abac
```
## 使用

``` go
import 	"github.com/Monkey-Mouse/go-abac/abac"

var ac AccessControl

ac.AddRules(IAccessInfo{
    Subject:  "foo",
    Action:   ActionUpdate,
    Resource: "bar",
    Rules: RulesType{MyRule{
        S: "dili",
        R: "dala",
    },},
})

res := ac.Can(IQueryInfo{
    Subject:  "foo",
    Action:   ActionUpdate,
    Resource: "bar",
})
// (bool)res = true/false
```
## Related
- [docs](docs)

##License
go-abac is MIT licensed. See the [LICENSE](LICENSE) file for details.


