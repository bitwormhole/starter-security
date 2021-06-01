package permissions

type Role struct {
	Name        string
	Permissions Permissions
}

type RoleList struct {
	text  string // like 'a,b,c'
	roles map[string]*Role
}

func (inst *RoleList) Init(text string) {
	/*
		// like 'a,b,c'
		text = strings.ToLower(text)
		srcList := strings.Split(text, ",")
		dstList := []string{}
		dstTable := map[string]string{}
		for i := range srcList {
			item := strings.TrimSpace(srcList[i])
			if item == "" {
				continue
			} else if dstTable[item] == item {
				continue
			}
			dstTable[item] = item
			dstList[i] = item
		}
		sort.Strings(dstList)
		builder := &strings.Builder{}
		sep := ""
		for i := range dstList {
			item := dstList[i]
			builder.WriteString(sep)
			builder.WriteString(item)
			sep = ","
		}
		inst.text = builder.String()
		inst.items = dstList
	*/
}
