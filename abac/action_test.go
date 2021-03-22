package abac

import (
	"fmt"
	"testing"
)

func TestResourceGrantsType_CreateAny(t *testing.T) {
	var ac = AccessControl{}
	ac.Grants = make(GrantsType)
	ac.Grants.Subject("foo").CreateAny(RulesType{}, "bar", "foo")

	fmt.Println(ac)
	//output:
	//{map[foo:map[bar:map[create:any:[]] foo:map[create:any:[]]]]}
}

func ExampleResourceGrantsType_Extend() {
	var ac = AccessControl{}
	ac.Grants = make(GrantsType)
	ac.Grants.Subject("foo").CreateAny(RulesType{}, "bar", "foo").Extend(ac, "bar")
	fmt.Println(ac)
	//output:
	//{map[bar:map[bar:map[create:any:[]] foo:map[create:any:[]]] foo:map[bar:map[create:any:[]] foo:map[create:any:[]]]]}
}
