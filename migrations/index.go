package migrations 


var Migrations []string = []string{
	UsersUp,
	OrganizationUp,
	OrgUserUp,
	OrgIdsUp,
	ContactUp,
	GroupUp,
	GroupUserUp,
	RequirementUp,
	SubmissionUp,
	WorkflowUp,
	
}

var Rollback []string = []string{
	UsersDown,
	OrganisationDown,
	OrgUserDown,
	OrgIdsDown,
	ContactDown,
	GroupDown,
	GroupUserDown,
	RequirementDown,
	SubmissionDown,
	WorkflowDown,
}