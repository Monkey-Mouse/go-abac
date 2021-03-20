package abac

type AccessControl struct {
}

type AcArgs struct {
	Subject  string `json:"subject,omitempty" example:"user"`
	Action   string `json:"action,omitempty" example:"delete"`
	Resource string `json:"resource,omitempty" example:"blog"`
}

func (ac *AccessControl) Can(args AcArgs) (res bool) {
	//todo check related rule
	//execute authorize handler
	//get result
	return false
}
