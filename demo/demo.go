package demo

import (
	"fmt"

	"github.com/bitwormhole/starter-security/auths"
	"github.com/bitwormhole/starter-security/auths/common"
	starter_security_etc "github.com/bitwormhole/starter-security/demo/etc"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/lang"
)

func Config(cb application.ConfigBuilder) error {
	return starter_security_etc.Config(cb)
}

func Run(context application.Context) error {

	err := tryLogin(context, "user1", "123")
	if err != nil {
		return err
	}

	err = tryLogin(context, "user2", "789")
	if err != nil {
		return err
	}

	return nil
}

func tryLogin(context application.Context, user string, password string) error {

	in := context.Injector()
	var authMngr auths.AuthenticationManager = nil

	in.Inject("#auths").To(func(o lang.Object) bool {
		o2, ok := o.(auths.AuthenticationManager)
		if ok {
			authMngr = o2
		}
		return ok
	})

	err := in.Done()
	if err != nil {
		return err
	}

	token := &common.AuthToken{
		UserName: user,
		Secret:   password,
		AuthType: "password",
	}

	info, err := authMngr.Authenticate(token)
	if err != nil {
		return err
	}

	fmt.Println(info)
	return nil
}
