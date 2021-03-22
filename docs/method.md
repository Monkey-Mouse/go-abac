## AccessControl: class

Construct an `AccessControl` instance by 
- `Grant()` passing a `GrantsType` object
- `SetGrant()` passing a `IAccessInfo` object
- `SetGrants()`array of `IAccessInfo` objects
### Grant()
``` go
	grants := abac.GrantsType{
		"role1": {
			"resource1": {
				"create:any": []abac.RuleType{},
				"read:own":   abac.RulesType{},
			},
			"resource2": {
				"create:any": []abac.RuleType{},
				"update:own": []abac.RuleType{},
			},
		},
		"role1": {...},
	}

	ac :=abac.AccessControl{}
	ac.Grant(grants)
```
### SetGrant()
``` go
	var ac AccessControl
	ac.Grants = make(GrantsType)

	var info IAccessInfo
	info = IAccessInfo{
		Subject:  "sub",
		Action:   "act",
		Resource: "res",
		Rules:    RulesType{FooRule{tips: "1"}, FooRule{tips: "2"}, FooRule{tips: "3"}},
	}
	ac.SetGrant(info)
```