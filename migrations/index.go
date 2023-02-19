package migrations 


var Migrations []string = []string{
	UsersUp,
	OrganizationUp,
	OrgUserUp,
	OrgIdsUp,
	
}

var Rollback []string = []string{
	UsersDown,
	OrganisationDown,
	OrgUserDown,
	OrgIdsDown,
}