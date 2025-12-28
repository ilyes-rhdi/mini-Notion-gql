package resolvers

var (
	User = NewUserResolver()

	Workspace = NewWorkspaceResolver()
	Page      = NewPageResolver()
	Block     = NewBlockResolver()
)
