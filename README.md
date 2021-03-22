## go-abac
translation: [简体中文](README_zh.md)|[English](README.md)

implement attribute based access control in golang
## Features
- implement access rules by self, extensible 
- minimal dependencies
- chained methods
- consistent style with same function project
- handy to use, detail docs

main reference：
[accessControl](https://github.com/onury/accesscontrol)

however, to simplify(?) the design, we decide to stick to the definition of abac, without the attribute role   
instead focus on the wider range of subject(including role/department/project)


## Install

```
go get github.com/Monkey-Mouse/go-abac
```
## Usage

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






