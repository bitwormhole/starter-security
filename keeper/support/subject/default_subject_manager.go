package subject

import (
	"context"

	"github.com/bitwormhole/starter-security/keeper"
	"github.com/bitwormhole/starter/markup"
)

// DefaultSubjectManager ...
type DefaultSubjectManager struct {
	markup.Component `id:"keeper-subject-manager" class:"keeper-configurer" `

	ksCtx keeper.SecurityContext
}

func (inst *DefaultSubjectManager) _Impl() (keeper.SubjectManager, keeper.Configurer) {
	return inst, inst
}

// Configure ...
func (inst *DefaultSubjectManager) Configure(c *keeper.Context) error {
	inst.ksCtx = c
	c.Subjects = inst
	return nil
}

// GetSubject ...
func (inst *DefaultSubjectManager) GetSubject(ctx context.Context) (keeper.Subject, error) {

	holder, err := keeper.GetHolder(ctx)
	if err != nil {
		return nil, err
	}

	ac := holder.GetAccessContext()
	subject := &DefaultSubject{ac: ac}

	ac.Context = ctx
	ac.SecurityContext = inst.ksCtx
	ac.Subject = subject

	// lazy load
	// sc.Access = nil
	// sc.Adapter = nil
	// sc.Session = nil

	return subject, nil
}
