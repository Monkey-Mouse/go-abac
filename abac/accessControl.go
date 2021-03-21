package abac

type roleType string
type SubjectType string
type ResourceType string
type ActionType string
type RulesType []string
type GrantsType map[SubjectType]ResourceGrantsType
type ResourceGrantsType map[ResourceType]ActionGrantsType
type ActionGrantsType map[ActionType]RulesType
type AccessControl struct {
	role   roleType `json:"role"`
	Grants GrantsType
}
type IAccessInfo struct {
	Subject  SubjectType  `json:"subject,omitempty" example:"user"`
	Action   ActionType   `json:"action,omitempty" example:"delete"`
	Resource ResourceType `json:"resource,omitempty" example:"blog"`
	Rules    RulesType    `json:"rules"`
}

type AcArgs struct {
	Subject  string `json:"subject,omitempty" example:"user"`
	Action   string `json:"action,omitempty" example:"delete"`
	Resource string `json:"resource,omitempty" example:"blog"`
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

func (ac *AccessControl) AddRules(info IAccessInfo) *GrantsType {
	//d:=ac.Grants[info.Subject][info.Resource][info.Action]
	ac.Grants[info.Subject][info.Resource][info.Action] = append(ac.Grants[info.Subject][info.Resource][info.Action], info.Rules...)
	return &ac.Grants
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

func (ac *AccessControl) Role(role roleType) *AccessControl {
	ac.role = role
	return ac
}

func (ac *AccessControl) Can(args AcArgs) (res bool) {
	//todo check related rule
	//execute authorize handler
	//get result
	return false
}
