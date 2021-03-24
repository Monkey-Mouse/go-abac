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

### import 
``` go
import 	"github.com/Monkey-Mouse/go-abac/abac"

var ac AccessControl
```
### construct rule

``` go
type DemoRule struct {
	id string	`json:"id" example:"u2020"`
}

func (r *DemoRule) ProcessContext(ctx abac.ContextType)  {
	// implement ProcessContext() to use params in context
	r.id=ctx.Value("id").(string)
}
func (r *DemoRule)JudgeRule()(bool,error) {
        // you can define your own rule here
	if r.id == "u2020"{
		return true,nil
	}else {
		return false,nil
	}
}

```

### config access rule 
look up more way to add rule [here](docs/rules.md)
``` go
grants := abac.GrantsType{
    "role1": {
        "resource1": {
            "create:any": []abac.RuleType{&DemoRule{}},
            "read:own":   abac.RulesType{},
        },
        "resource2": {
            "create:any": []abac.RuleType{},
            "update:own": []abac.RuleType{},
        },
    },
}
ac.Grant(grants)
```
### judge access rule
``` go
	resFail:=ac.CanAnd(abac.IQueryInfo{
		Subject:  "role1",
		Action:   "create:any",
		Resource: "resource1",
		Context:  DemoContext{"id":"u3030"},
	})
	if resFail==true{
		t.Errorf("should fail")
	}
	resPass:=ac.CanAnd(abac.IQueryInfo{
		Subject:  "role1",
		Action:   "create:any",
		Resource: "resource1",
		Context:  DemoContext{"id":"u2020"},
	})
	if resPass==false{
		t.Errorf("should pass")
	}
```

## Related
- [docs](docs)

##License
go-abac is MIT licensed. See the [LICENSE](LICENSE) file for details.






