package abac

import (
	"context"
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
	ProcessContext(ctx ContextType)
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
	Context  ContextType  `json:"context,omitempty" example:"blog"`
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

// SetGrant set one info for ac
func (ac *AccessControl) SetGrant(info IAccessInfo) *GrantsType {
	ac.Grants[info.Subject] = ResourceGrantsType{info.Resource: ActionGrantsType{info.Action: info.Rules}}
	return &ac.Grants
}

// SetGrants
// set multi infos for ac
func (ac *AccessControl) SetGrants(infos ...IAccessInfo) *GrantsType {
	for _, info := range infos {
		ac.SetGrant(info)
	}
	return &ac.Grants
}

//GetGrants get all grants within a controller
func (ac *AccessControl) GetGrants() GrantsType {
	grants := make(GrantsType)
	for key, value := range ac.Grants {
		grants[key] = value
	}
	return grants
}

//GetSubject get grants of a certain subject
func (g GrantsType) GetSubject(subject SubjectType) ResourceGrantsType {
	resGrants := make(ResourceGrantsType)
	for key, value := range g[subject] {
		resGrants[key] = value
	}
	return resGrants
}

//Subject get grants of a certain subject
func (g GrantsType) Subject(subject SubjectType) ResourceGrantsType {
	if !subject.IsZero() && g[subject] == nil {
		g[subject] = make(ResourceGrantsType)
	}
	return g[subject]
}

//GetResource get grants of a certain resource
func (r ResourceGrantsType) GetResource(resource ResourceType) ActionGrantsType {
	actGrants := make(ActionGrantsType)
	for key, value := range r[resource] {
		actGrants[key] = value
	}
	return actGrants
}

//Resource get grants of a certain resource
func (r ResourceGrantsType) Resource(resource ResourceType) ActionGrantsType {
	return r[resource]
}

//GetAction get all rules of a certain action
func (a ActionGrantsType) GetAction(action ActionType) RulesType {
	ruleGrants := make(RulesType, len(a[action]))
	for i, rule := range a[action] {
		ruleGrants[i] = rule
	}
	return ruleGrants
}

//Action get all rules of a certain action
func (a ActionGrantsType) Action(action ActionType) RulesType {
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

// CanAnd  check related rule
//		execute authorize handler
//		get result
// logic: and(if any rule failed, can = false)
func (ac *AccessControl) CanAnd(info IQueryInfo) (can bool, err error) {
	return andProcessRule(info.Context, ac.GetRules(info))
}

// CanOr  check related rule
//		execute authorize handler
//		get result
// logic: or(if any rule passed, can = true)
func (ac *AccessControl) CanOr(info IQueryInfo) (can bool, err error) {
	return orProcessRule(info.Context, ac.GetRules(info))
}

// CanHandler  check related rule
//		execute authorize handler
//		get result
// logic: or(if any rule passed, can = true)
func (ac *AccessControl) CanHandler(info interface{}, handler func(info interface{}) (bool, error)) (can bool, err error) {
	return handler(info)
}
