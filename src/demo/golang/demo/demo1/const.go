package demo1

import "github.com/bitwormhole/starter-security/keeper"

const (
	theUserName = "demo"
	theUserPass = "demo"
)

func configAccess() keeper.Access {
	ab := keeper.AccessBuilder{
		Method:      "GET",
		Path:        "/demo/xyz",
		PathPattern: "/demo/:id",
		Params:      map[string]string{"id": "this-is-a-demo"},
	}
	return ab.Create()
}
