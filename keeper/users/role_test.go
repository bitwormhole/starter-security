package users

import "testing"

func TestRoles(t *testing.T) {

	str1 := "this,is,a,test"
	t.Log("str1", str1)

	roles1 := Roles(str1)
	t.Log("r1=", roles1)

	r2 := roles1.Add("xyz")
	t.Log("xyz=", r2)

	list := roles1.List()
	t.Log("list=", list)

	r3 := roles1.Normalize()
	t.Log("r3=", r3)

	r4 := r2.Remove("xyz")
	t.Log("r4=", r4)

	str2 := roles1.String()
	t.Log("str2=", str2)

}
