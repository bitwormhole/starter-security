package permission

import (
	"errors"
	"strings"

	"github.com/bitwormhole/starter-security/keeper"
	"github.com/bitwormhole/starter-security/keeper/users"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

type permissionTemplateTableBuilder struct {
	context application.Context

	initialled bool
	table      map[string]*keeper.SimplePermissionRegistration
	loaders    map[string]keeper.PermissionLoader
	templates  map[string]keeper.PermissionTemplate
}

func (inst *permissionTemplateTableBuilder) init() {
	if inst.initialled {
		return
	}
	inst.table = make(map[string]*keeper.SimplePermissionRegistration)
	inst.loaders = make(map[string]keeper.PermissionLoader)
	inst.templates = make(map[string]keeper.PermissionTemplate)
	inst.initialled = true
}

func (inst *permissionTemplateTableBuilder) getTable() map[string]*keeper.SimplePermissionRegistration {
	inst.init()
	return inst.table
}

// func (inst *permissionTemplateTableBuilder) registerLoader(cpr *keeper.ComplexPermissionRegistration) error {
// 	loader := cpr.Loader
// 	name := cpr.LoaderName
// 	if loader == nil || name == "" {
// 		return nil
// 	}
// 	inst.init()
// 	table := inst.loaders
// 	old := table[name]
// 	if old != nil {
// 		return errors.New("the keeper.PermissionLoader is override, which named: " + name)
// 	}
// 	table[name] = loader
// 	return nil
// }

// func (inst *permissionTemplateTableBuilder) registerTemplate(cpr *keeper.ComplexPermissionRegistration) error {
// 	temp := cpr.Template
// 	name := cpr.TemplateName
// 	if temp == nil || name == "" {
// 		return nil
// 	}
// 	inst.init()
// 	table := inst.templates
// 	old := table[name]
// 	if old != nil {
// 		return errors.New("the keeper.PermissionTemplate is override, which named: " + name)
// 	}
// 	table[name] = temp
// 	return nil
// }

func (inst *permissionTemplateTableBuilder) LoadFromSimplePRs(src []*keeper.SimplePermissionRegistration) error {
	dst := inst.getTable()
	for _, item := range src {
		key := makeAccessKey(item.Method, item.PathPattern)
		old := dst[key]
		if old != nil {
			return errors.New("the permission is override: " + key)
		}
		dst[key] = item
	}
	return nil
}

func (inst *permissionTemplateTableBuilder) LoadFromComplexPRs(src []*keeper.ComplexPermissionRegistration) error {
	for _, item := range src {
		mid, err := inst.makeSimplePRs(item)
		if err != nil {
			return err
		}
		err = inst.LoadFromSimplePRs(mid)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *permissionTemplateTableBuilder) LoadFromRegistryList(src []keeper.PermissionRegistry) error {
	for _, item := range src {
		mid := item.GetRegistrationList()
		err := inst.LoadFromComplexPRs(mid)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *permissionTemplateTableBuilder) LoadFromResource(ctx application.Context, name string) error {

	text, err := ctx.GetResources().GetText(name)
	if err != nil {
		return err
	}

	props, err := collection.ParseProperties(text, nil)
	if err != nil {
		return err
	}

	const prefix = "permission."
	const suffix = ".methods"
	ids := inst.getIDsFromProperties(props, prefix, suffix)
	dst := make([]*keeper.ComplexPermissionRegistration, 0)

	for _, id := range ids {
		complex, err := inst.loadComplexPRFromProperties(props, prefix+id+".")
		if err != nil {
			return err
		}
		dst = append(dst, complex)
	}
	return inst.LoadFromComplexPRs(dst)
}

func (inst *permissionTemplateTableBuilder) loadComplexPRFromProperties(p collection.Properties, keyPrefix string) (*keeper.ComplexPermissionRegistration, error) {

	getter := p.Getter()
	loaderSel := getter.GetString(keyPrefix+"loader", "#keeper-default-permission-loader-factory")
	templateSel := getter.GetString(keyPrefix+"template", "#keeper-default-permission-template-factory")
	enabled := getter.GetBool(keyPrefix+"enabled", true)

	getter.CleanError()
	methods := getter.GetString(keyPrefix+"methods", "")
	paths := getter.GetString(keyPrefix+"paths", "")
	roles := getter.GetString(keyPrefix+"roles", "")

	err := getter.Error()
	if err != nil {
		return nil, err
	}

	cpr := &keeper.ComplexPermissionRegistration{}
	cpr.Methods = inst.parseStringList(methods)
	cpr.Paths = inst.parseStringList(paths)
	cpr.Roles = users.Roles(roles).List()
	cpr.LoaderFactorySelector = loaderSel
	cpr.TemplateFactorySelector = templateSel
	cpr.Enabled = enabled
	return cpr, nil
}

func (inst *permissionTemplateTableBuilder) getIDsFromProperties(p collection.Properties, prefix, suffix string) []string {
	dst := make([]string, 0)
	all := p.Export(nil)
	for k := range all {
		if strings.HasPrefix(k, prefix) && strings.HasSuffix(k, suffix) {
			id := k[len(prefix) : len(k)-len(suffix)]
			if id == "" {
				continue
			}
			dst = append(dst, id)
		}
	}
	return dst
}

func (inst *permissionTemplateTableBuilder) parseStringList(s string) []string {
	src := strings.Split(s, ",")
	dst := make([]string, 0)
	for _, item := range src {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		dst = append(dst, item)
	}
	return dst
}

func (inst *permissionTemplateTableBuilder) Create() (map[string]*keeper.SimplePermissionRegistration, error) {
	table := inst.getTable()
	for _, item := range table {
		err := inst.wire(item)
		if err != nil {
			return nil, err
		}
	}
	return table, nil
}

func (inst *permissionTemplateTableBuilder) wire(spr *keeper.SimplePermissionRegistration) error {

	loader := spr.Loader
	template := spr.Template

	if loader == nil {
		ldr, err := inst.makePermLoader(spr)
		if err != nil {
			return err
		}
		spr.Loader = ldr
	}

	if template == nil {
		temp, err := inst.makePermTemplate(spr)
		if err != nil {
			return err
		}
		spr.Template = temp
	}

	return nil
}

func (inst *permissionTemplateTableBuilder) makePermTemplate(spr *keeper.SimplePermissionRegistration) (keeper.PermissionTemplate, error) {
	selector := spr.TemplateFactorySelector
	o1, err := inst.context.GetComponent(selector)
	if err != nil {
		return nil, err
	}
	o2 := o1.(keeper.PermissionTemplateFactory)
	return o2.CreateTemplate(spr)

}

func (inst *permissionTemplateTableBuilder) makePermLoader(spr *keeper.SimplePermissionRegistration) (keeper.PermissionLoader, error) {
	selector := spr.LoaderFactorySelector
	o1, err := inst.context.GetComponent(selector)
	if err != nil {
		return nil, err
	}
	o2 := o1.(keeper.PermissionLoaderFactory)
	return o2.CreateLoader(spr)
}

func (inst *permissionTemplateTableBuilder) makeSimplePRs(cpr *keeper.ComplexPermissionRegistration) ([]*keeper.SimplePermissionRegistration, error) {
	methods := cpr.Methods
	paths := cpr.Paths
	dst := make([]*keeper.SimplePermissionRegistration, 0)
	for _, method := range methods {
		for _, path := range paths {
			spr := &keeper.SimplePermissionRegistration{}
			spr.Method = method
			spr.PathPattern = path
			spr.Roles = cpr.Roles
			spr.LoaderFactorySelector = cpr.LoaderFactorySelector
			spr.TemplateFactorySelector = cpr.TemplateFactorySelector
			spr.Enabled = cpr.Enabled
			spr.Loader = cpr.Loader
			spr.Template = cpr.Template
			dst = append(dst, spr)
		}
	}
	return dst, nil
}

func makeAccessKey(method, path string) string {
	method = strings.TrimSpace(method)
	method = strings.ToUpper(method)
	parts := strings.Split(path, "/")
	builder := strings.Builder{}
	builder.WriteString(method)
	builder.WriteString(":")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		builder.WriteString("/")
		builder.WriteString(part)
	}
	return builder.String()
}
