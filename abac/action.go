package abac

const (
	ActionCreate = "create"
	ActionRead   = "read"
	ActionUpdate = "update"
	ActionDelete = "delete"
)
const (
	ActionCreateAny = "create:any"
	ActionCreateOwn = "create:own"
	ActionReadAny   = "read:any"
	ActionReadOwn   = "read:own"
	ActionUpdateAny = "update:any"
	ActionUpdateOwn = "update:own"
	ActionDeleteAny = "delete:any"
	ActionDeleteOwn = "delete:own"
)

func (r ResourceGrantsType) EnsureResourceGrants(resource ResourceType) {
	if !resource.IsZero() && r[resource] == nil {
		r[resource] = make(ActionGrantsType)
	}
}

// CreateAny
func (r ResourceGrantsType) CreateAny(rule RulesType, resources ...ResourceType) ResourceGrantsType {
	for _, resource := range resources {
		r.EnsureResourceGrants(resource)
		r[resource][ActionCreateAny] = rule
	}
	return r
}

// Extend extend resourceGrants to ac
func (r ResourceGrantsType) Extend(ac AccessControl, subject SubjectType) ResourceGrantsType {
	if !subject.IsZero() && ac.Grants[subject] == nil {
		ac.Grants[subject] = make(ResourceGrantsType)
	}
	ac.Grants[subject] = r
	return r
}
