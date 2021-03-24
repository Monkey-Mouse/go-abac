# rules

for better extensible, the type of rules in the model(`RulesType`) is designed as array of interface   
we suggest user implementing handler method `JudgeRule()` to check if the request of subject satisfy the rule  
and hereby, this package provides method `Can()` to return whether the request meet the rules above

> [problem] not enough info for json serialize

## RulesType

``` go
type RulesType []RuleType
type RuleType interface {
	JudgeRule() (bool, error)
	ProcessContext(ctx ContextType)
}
```


### JudgeRule()(bool,error)

## AddRules()

``` go
	ac.AddRules(IAccessInfo{
		Subject:  "foo",
		Action:   ActionUpdate,
		Resource: "bar",
		Rules: RulesType{MyRule{
			S: "dili",
			R: "dala",
		}},
	})

```

above is equivalent to:
``` go
	ac.Grant(abac.GrantsType{
		"foo": {
			"bar": {
				abac.ActionCreateAny: []abac.RuleType{},
			},
		},
	})

```
(however, ` reflect.deepEqual(acAbove,acBelow))` return false, for the array of rule is `nil` above but empty below)

## Can() 
### logic

**!!! or**

check each item in the rules array, if **any** rule satisfy(`JudgeRule()`return`true,nil`), 
the request have right to access the resources

