package resolvers

import (
	"github.com/graphql-go/graphql"
	"github.com/ilyes-rhdi/buildit-Gql/internal/services"
)

type UserResolver struct {
	profile *services.ProfileService
}

func NewUserResolver() *UserResolver {
	return &UserResolver{profile: services.NewProfileService()}
}

func (r *UserResolver) Me(p graphql.ResolveParams) (any, error) {
	uid, err := getUserID(p)
	if err != nil {
		return nil, err
	}
	u, err := r.profile.GetUser(uid)
	if err != nil {
		return nil, err
	}
	return userToMap(*u), nil
}
