package element

import (
	"errors"

	"github.com/bitwormhole/starter-security/auths/common"
)

type DemoPasswordAuthDAO struct {
}

func (inst *DemoPasswordAuthDAO) _impl_() common.AuthDAO {
	return inst
}

func (inst *DemoPasswordAuthDAO) Find(user string, mechanism string) (common.Auth, error) {

	if user == "user1" {
		item := &DemoPasswordAuthData{
			accountId:   "a1",
			accountUUID: "a1111",
			salt:        "qwerty",
			mech:        "password",
			secret:      "037927422e66f1d7ddeaac0441c7e6cba4784e6748bb689a25f79686bca7aaa1",
		}
		return item, nil
	} else if user == "user2" {
		item := &DemoPasswordAuthData{
			accountId:   "a2",
			accountUUID: "a2222",
			salt:        "qwerty",
			mech:        "password",
			secret:      "8097c3ac988f90a1471ed5928b7aaee45e33f1f443fcccf5c7dffe5efce18c9f",
		}
		return item, nil
	}
	return nil, errors.New("no user:" + user)
}

////////////////////////////////////////////////////////////////////////////////

type DemoPasswordAuthData struct {
	accountId   string
	accountUUID string
	salt        string
	secret      string
	mech        string
}

func (inst *DemoPasswordAuthData) AccountID() string {
	return inst.accountId
}

func (inst *DemoPasswordAuthData) AccountUUID() string {
	return inst.accountUUID
}

func (inst *DemoPasswordAuthData) Salt() string {
	return inst.salt
}

func (inst *DemoPasswordAuthData) Secret() string {
	return inst.secret
}

func (inst *DemoPasswordAuthData) Mechanism() string {
	return inst.mech
}
