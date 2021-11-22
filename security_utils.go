package security

import (
	"context"
	"errors"

	"github.com/bitwormhole/starter/contexts"
)

////////////////////////////////////////////////////////////////////////////////
// public

// GetSubject 从给定的上下文取 Subject 对象
func GetSubject(ctx context.Context) (Subject, error) {
	holder, err := getSubjectHolder(ctx)
	if err != nil {
		return nil, err
	}
	return holder.getSubject()
}

// GetSecurityManager 取安全管理器
func GetSecurityManager(ctx context.Context) (SecurityManager, error) {
	holder, err := getSubjectHolder(ctx)
	if err != nil {
		return nil, err
	}
	return holder.getSecurityManager()
}

// SetSecurityManager 设置安全管理器
func SetSecurityManager(ctx context.Context, sm SecurityManager) error {

	holder, err := getSubjectHolder(ctx)
	if err != nil {
		return err
	}

	sc, err := holder.getSubjectContext()
	if err != nil {
		return err
	}

	sc.SetSecurityManager(sm)
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// private

type subjectHolder struct {
	subject Subject
	context SubjectContext
}

func (inst *subjectHolder) getSubject() (Subject, error) {

	subject := inst.subject
	if subject != nil {
		return subject, nil
	}

	ctx, err := inst.getSubjectContext()
	if err != nil {
		return nil, err
	}

	sm, err := inst.getSecurityManager()
	if err != nil {
		return nil, err
	}

	subject = sm.CreateSubject(ctx)
	if subject == nil {
		return nil, errors.New("SecurityManager.CreateSubject(ctx)==nil")
	}

	inst.subject = subject
	return subject, nil
}

func (inst *subjectHolder) getSecurityManager() (SecurityManager, error) {
	sc, err := inst.getSubjectContext()
	if err != nil {
		return nil, err
	}
	man := sc.GetSecurityManager()
	if man == nil {
		return nil, errors.New("no SecurityManager setup, need for security.SetSecurityManager()")
	}
	return man, nil
}

func (inst *subjectHolder) getSubjectContext() (SubjectContext, error) {
	ctx := inst.context
	if ctx == nil {
		return nil, errors.New("no SecurityManager setup, need for security.SetSecurityManager()")
	}
	return ctx, nil
}

func getSubjectHolder(ctx context.Context) (*subjectHolder, error) {

	const key = "github.com/bitwormhole/starter/security/subject-holder#binding"

	o1 := ctx.Value(key)
	holder, ok := o1.(*subjectHolder)
	if ok {
		return holder, nil
	}

	ctx2, err := contexts.GetContextSetter(ctx)
	if err != nil {
		return nil, err
	}

	holder = &subjectHolder{}
	ctx2.SetValue(key, holder)
	return holder, nil
}
