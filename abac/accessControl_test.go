package abac

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"
	"time"
)

type FooRule struct {
	tips string
}

func (f FooRule) JudgeRule() (bool, error) {
	return true, nil
}
func (f FooRule) setTips(tips string) {
	f.tips = tips
}
func (f FooRule) getTips() string {
	return f.tips
}

type MyRule struct {
	S SubjectType
	R ResourceType
}

func (m MyRule) JudgeRule() (bool, error) {
	fmt.Println(m.R)
	return true, nil
}

type FailRule struct {
	S SubjectType
	R ResourceType
}

func (m FailRule) JudgeRule() (bool, error) {
	fmt.Println(m)
	return false, nil
}

type TimeConsume struct {
}

func (m TimeConsume) JudgeRule() (bool, error) {
	time.Sleep(time.Second * 1)
	return true, nil
}

func init() {

}

func ExampleAccessControl_SetGrant() {
	var ac AccessControl
	ac.Grants = make(GrantsType)

	var info IAccessInfo
	info = IAccessInfo{
		Subject:  "sub",
		Action:   "act",
		Resource: "res",
		Rules:    RulesType{FooRule{tips: "1"}, FooRule{tips: "2"}, FooRule{tips: "3"}},
	}
	res, err := json.Marshal(ac)
	if err != nil {
		log.Fatal(err)
	} else {
		os.Stdout.Write(res)
	}
	//encoding.TextMarshaler()

	ac.SetGrant(info)

	res, err = json.Marshal(ac)
	if err != nil {
		log.Fatal(err)
	} else {
		os.Stdout.Write(res)

	}

	fmt.Println(ac)

	//output:
	//{"Grants":{}}{"Grants":{"sub":{"res":{"act":[{},{},{}]}}}}{map[sub:map[res:map[act:[{1} {2} {3}]]]]}
}
func ExampleAccessControl_Grant() {
	var ac AccessControl
	ac.Grant(GrantsType{})
	fmt.Println(ac)
	grants := ac.Grants
	grants["test"] = ResourceGrantsType{}
	fmt.Println(ac)

	//output:
	//{map[]}
	//{map[test:map[]]}
}

func ExampleAccessControl_GetGrants() {
	var foo AccessControl
	foo.Grant(GrantsType{"account": ResourceGrantsType{"book": ActionGrantsType{ActionCreate: RulesType{FooRule{tips: "has user role"}}}},
		"role": ResourceGrantsType{"project": ActionGrantsType{ActionCreate: RulesType{FooRule{tips: "has primer user role"}}}},
	})
	fmt.Println(foo.GetGrants())

	//output:
	//map[account:map[book:map[create:[{has user role}]]] role:map[project:map[create:[{has primer user role}]]]]
}

func ExampleGrantsType_GetSubject() {
	var foo AccessControl
	var subject1 = SubjectType("account")
	foo.Grant(GrantsType{"account": ResourceGrantsType{"book": ActionGrantsType{ActionCreate: RulesType{FooRule{tips: "has user role"}}}},
		"role": ResourceGrantsType{"project": ActionGrantsType{ActionCreate: RulesType{FooRule{tips: "has primer user role"}}}},
	})
	fmt.Println(foo.GetGrants().GetSubject(subject1))

	//output:
	//map[book:map[create:[{has user role}]]]
}

func TestGrantsType_GetSubject(t *testing.T) {
	var foo AccessControl
	var subject1 = SubjectType("account")
	foo.Grant(GrantsType{"account": ResourceGrantsType{"book": ActionGrantsType{ActionCreate: RulesType{FooRule{tips: "has user role"}}}},
		"role": ResourceGrantsType{"project": ActionGrantsType{ActionCreate: RulesType{FooRule{tips: "has primer user role"}}}},
	})
	getBefore := foo.GetGrants().GetSubject(subject1)

	var res = foo.GetGrants().GetSubject(subject1)
	res["book"] = ActionGrantsType{ActionUpdate: RulesType{FooRule{tips: "has user role"}}}
	getAfter := foo.GetGrants().GetSubject(subject1)
	if !reflect.DeepEqual(getBefore, getAfter) {
		t.Errorf("get return reference of object, after %v != before %v", getAfter, getBefore)
	}

}

func ExampleResourceGrantsType_GetResource() {
	var foo AccessControl
	var subject1 = SubjectType("account")
	foo.Grant(GrantsType{"account": ResourceGrantsType{"book": ActionGrantsType{ActionCreate: RulesType{FooRule{tips: "has user role"}}}},
		"role": ResourceGrantsType{"project": ActionGrantsType{ActionCreate: RulesType{FooRule{tips: "has primer user role"}}}},
	})
	fmt.Println(foo.GetGrants().GetSubject(subject1).GetResource("book"))

	//output:
	//map[create:[{has user role}]]
}

func ExampleActionGrantsType_GetAction() {
	var foo AccessControl
	var subject1 = SubjectType("account")
	foo.Grant(GrantsType{"account": ResourceGrantsType{"book": ActionGrantsType{ActionCreate: RulesType{FooRule{tips: "has user role"}}}},
		"role": ResourceGrantsType{"project": ActionGrantsType{ActionCreate: RulesType{FooRule{tips: "has primer user role"}}}},
	})
	fmt.Println(foo.GetGrants().GetSubject(subject1).GetResource("book").GetAction(ActionDelete))
	fmt.Println(foo.GetGrants().GetSubject(subject1).GetResource("book").GetAction(ActionCreate))

	//output:
	//[]
	//[{has user role}]
}

func TestActionGrantsType_GetAction(t *testing.T) {
	var foo AccessControl
	foo.Grant(GrantsType{"account": ResourceGrantsType{"book": ActionGrantsType{ActionCreate: RulesType{FooRule{tips: "has user role"}}}},
		"role": ResourceGrantsType{"project": ActionGrantsType{ActionCreate: RulesType{FooRule{tips: "has primer user role"}}}},
	})
	newFoo := foo
	if !reflect.DeepEqual(newFoo, foo) {
		t.Errorf("%v not equal to %v", newFoo, foo)
	}
	if &newFoo == &foo {
		t.Errorf("%v's pointer is equal to %v's", newFoo, foo)
	}
	newFoo.AddRules(IAccessInfo{
		Subject:  "foo",
		Action:   ActionUpdate,
		Resource: "bar",
		Rules:    RulesType{},
	})
	action1 := newFoo.GetGrants().GetSubject("foo").GetResource("bar").GetAction(ActionUpdate)

	fmt.Println(action1)
	fmt.Println(reflect.TypeOf(action1))
	for _, action := range action1 {
		fmt.Println(reflect.TypeOf(action))
	}

}

func TestAccessControl_Can(t *testing.T) {
	var ac AccessControl

	ac.AddRules(IAccessInfo{
		Subject:  "foo",
		Action:   ActionUpdate,
		Resource: "bar",
		Rules: RulesType{MyRule{
			S: "dili",
			R: "dala",
		}, TimeConsume{}},
	})

	res := ac.Can(IQueryInfo{
		Subject:  "foo",
		Action:   ActionUpdate,
		Resource: "bar",
	})
	if res {
		fmt.Println("pass")
	} else {
		fmt.Println("deny")
	}
}

func ExampleAccessControl_Can() {
	var ac AccessControl

	ac.AddRules(IAccessInfo{
		Subject:  "foo",
		Action:   ActionUpdate,
		Resource: "bar",
		Rules: RulesType{MyRule{
			S: "dili",
			R: "dala",
		}},
	})

	res := ac.Can(IQueryInfo{
		Subject:  "foo",
		Action:   ActionUpdate,
		Resource: "bar",
	})
	if res {
		fmt.Println("pass")
	} else {
		fmt.Println("deny")
	}
	//output:
	//dala
	//pass
}

func TestAccessControl_GetRules(t *testing.T) {
	var foo AccessControl
	foo.Grant(GrantsType{"account": ResourceGrantsType{"book": ActionGrantsType{ActionCreate: RulesType{FooRule{tips: "has user role"}}}},
		"role": ResourceGrantsType{"project": ActionGrantsType{ActionCreate: RulesType{FooRule{tips: "has primer user role"}}}},
	})

	rules := foo.GetRules(IQueryInfo{
		Subject:  "account",
		Action:   ActionCreate,
		Resource: "book",
	})
	if rules == nil {
		t.Errorf("rules come out to nil")
	}
	rules = foo.GetRules(IQueryInfo{
		Subject:  "account",
		Action:   ActionDelete,
		Resource: "book",
	})
	if rules != nil {
		t.Errorf("rules come out to be %v", rules)
	}
}

func Test_processRule(t *testing.T) {
	type args struct {
		ctx   context.Context
		rules RulesType
	}
	tests := []struct {
		name     string
		args     args
		wantPass bool
	}{
		{name: "text process rule fail", args: args{ctx: context.TODO(), rules: []RuleType{
			FooRule{},
			FailRule{},
			FooRule{},
		}}, wantPass: false},
		{name: "text process rule pass", args: args{ctx: context.TODO(), rules: []RuleType{
			FooRule{},
			FooRule{},
			FooRule{},
		}}, wantPass: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPass := processRule(tt.args.ctx, tt.args.rules); gotPass != tt.wantPass {
				t.Errorf("processRule() = %v, want %v", gotPass, tt.wantPass)
			}
		})
	}
}
