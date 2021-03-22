package abac

import (
	"context"
	"log"
	"sync"
	"time"
)

type SubjectEntity interface{}
type ResourceEntity interface{}

type roleType string
type SubjectType string
type ResourceType string
type ActionType string
type RulesType []RuleType
type RuleType interface {
	JudgeRule() (bool, error)
}
type GrantsType map[SubjectType]ResourceGrantsType
type ResourceGrantsType map[ResourceType]ActionGrantsType
type ActionGrantsType map[ActionType]RulesType
type AccessControl struct {
	Grants GrantsType
}
type IAccessInfo struct {
	Subject  SubjectType  `json:"subject,omitempty" example:"user"`
	Action   ActionType   `json:"action,omitempty" example:"delete"`
	Resource ResourceType `json:"resource,omitempty" example:"blog"`
	Rules    RulesType    `json:"rules"`
}

type IQueryInfo struct {
	Subject  SubjectType  `json:"subject,omitempty" example:"user"`
	Action   ActionType   `json:"action,omitempty" example:"delete"`
	Resource ResourceType `json:"resource,omitempty" example:"blog"`
}

// zero return zero value of SubjectType
func (sT SubjectType) zero() SubjectType {
	return ""
}

// IsZero check whether a variable of SubjectType is zero
func (sT SubjectType) IsZero() bool {
	return sT == sT.zero()
}

// zero return zero value of ResourceType
func (rT ResourceType) zero() ResourceType {
	return ""
}

// IsZero check whether a variable of ResourceType is zero
func (rT ResourceType) IsZero() bool {
	return rT == rT.zero()
}

// zero return zero value of ResourceType
func (aT ActionType) zero() ActionType {
	return ""
}

// IsZero check whether a variable of ResourceType is zero
func (aT ActionType) IsZero() bool {
	return aT == aT.zero()
}

func (ac *AccessControl) Grant(grantsType2 GrantsType) *GrantsType {
	ac.Grants = grantsType2
	return &ac.Grants
}
func (ac *AccessControl) SetGrant(info IAccessInfo) *GrantsType {
	ac.Grants[info.Subject] = ResourceGrantsType{info.Resource: ActionGrantsType{info.Action: info.Rules}}
	return &ac.Grants
}

//GetGrants get all grants within a controller
func (ac *AccessControl) GetGrants() GrantsType {
	return ac.Grants
}

//GetSubject get all grants of a certain subject
func (g GrantsType) GetSubject(subject SubjectType) ResourceGrantsType {
	return g[subject]
}

//GetResource get grants of a certain resource
func (r ResourceGrantsType) GetResource(resource ResourceType) ActionGrantsType {
	return r[resource]
}

//GetAction get all rules of a certain action
func (a ActionGrantsType) GetAction(action ActionType) RulesType {
	return a[action]
}

// EnsureMap check if the map to visit nil, if nil, make new one
func (ac *AccessControl) EnsureMap(info IAccessInfo) *AccessControl {
	if ac.Grants == nil {
		ac.Grants = make(GrantsType)
	}
	if !info.Subject.IsZero() && ac.Grants[info.Subject] == nil {
		ac.Grants[info.Subject] = make(ResourceGrantsType)
	}
	if !info.Resource.IsZero() && ac.Grants[info.Subject][info.Resource] == nil {
		ac.Grants[info.Subject][info.Resource] = make(ActionGrantsType)
	}
	return ac
}

// EnsureMap check if the map to visit nil, if nil, make new one
func (ac *AccessControl) CheckMap(info IQueryInfo) bool {

	if ac.Grants == nil || ac.Grants[info.Subject] == nil || ac.Grants[info.Subject][info.Resource] == nil {
		return false
	}
	return true
}

// AddRules append rules to subject
func (ac *AccessControl) AddRules(info IAccessInfo) *GrantsType {
	// in case of nil map
	ac.EnsureMap(info)

	ac.Grants[info.Subject][info.Resource][info.Action] = append(ac.Grants[info.Subject][info.Resource][info.Action], info.Rules...)
	return &ac.Grants
}

// GetRules get rules to subject
func (ac *AccessControl) GetRules(info IQueryInfo) RulesType {
	// in case of nil map
	if ac.CheckMap(info) {
		return ac.Grants[info.Subject][info.Resource][info.Action]
	}
	return nil
}

func (a *IAccessInfo) Set(info IAccessInfo) *IAccessInfo {
	a = &IAccessInfo{
		Subject:  info.Subject,
		Action:   info.Action,
		Resource: info.Resource,
		Rules:    info.Rules,
	}
	return a
}
func (a *IAccessInfo) SetSubject(subject SubjectType) *IAccessInfo {
	a.Subject = subject
	return a
}
func (a *IAccessInfo) SetResource(resource ResourceType) *IAccessInfo {
	a.Resource = resource
	return a
}
func (a *IAccessInfo) SetAction(action ActionType) *IAccessInfo {
	a.Action = action
	return a
}
func (a *IAccessInfo) SetRule(rule RulesType) *IAccessInfo {
	a.Rules = rule
	return a
}

// Can  check related rule
//		execute authorize handler
//		get result
// use goroutine
func (ac *AccessControl) Can(info IQueryInfo) (resc bool) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	rules := ac.GetRules(info)
	return processRule(ctx, rules)
}

func processRule(ctx context.Context, rules RulesType) (pass bool) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	wg := new(sync.WaitGroup)
	wg.Add(len(rules))
	for _, rule := range rules {
		go func(rule RuleType, ctx context.Context) {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			default:
				if res, err := rule.JudgeRule(); err != nil {
					log.Println(err)
				} else if res {
					cancel()
					pass = res
					return
				}
			}
		}(rule, ctx)
	}
	wg.Wait()
	return
}

func testCtx(ctx context.Context) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	select {
	case <-ctx.Done():
		return false, ctx.Err()
	case <-time.After(1 * time.Minute):

	}
	print("here")
	time.Sleep(time.Minute * 1)
	return true, nil
}
