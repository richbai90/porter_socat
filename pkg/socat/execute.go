package socat

import (
	"context"

	"get.porter.sh/porter/pkg/exec/builder"
	yaml "gopkg.in/yaml.v2"
)

func (m *Mixin) loadAction() (*Action, error) {
	// if the action doesn't complete in 5 minutes, throw an error
	ctx := context.TODO()
	var action Action
	err := builder.LoadAction(ctx, m.Context, "", func(contents []byte) (interface{}, error) {
		err := yaml.Unmarshal(contents, &action)
		return &action, err
	})
	return &action, err
}

func (m *Mixin) Execute() error {
	action, err := m.loadAction()
	if err != nil {
		return err
	}

	_, err = builder.ExecuteSingleStepAction(m.Context, action)
	return err
}
