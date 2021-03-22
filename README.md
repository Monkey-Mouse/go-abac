## go-abac
translation: [简体中文](README_zh.md)|[English](README.md)

[![codecov](https://codecov.io/gh/Monkey-Mouse/go-abac/branch/main/graph/badge.svg?token=8X3HF5VFWT)](https://codecov.io/gh/Monkey-Mouse/go-abac)
![gobadge](https://github.com/Monkey-Mouse/go-abac/actions/workflows/.github/workflows/go.yml/badge.svg)
![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/Monkey-Mouse/go-abac.svg)](https://github.com/Monkey-Mouse/go-abac)
[![GoReportCard](https://goreportcard.com/badge/github.com/Monkey-Mouse/mo2)](https://goreportcard.com/report/github.com/Monkey-Mouse/mo2)
[![BCH compliance](https://bettercodehub.com/edge/badge/Monkey-Mouse/go-abac?branch=main)](https://bettercodehub.com/)
[![license](https://badgen.net/github/license/Monkey-Mouse/go-abac)](/LICENSE)  


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






