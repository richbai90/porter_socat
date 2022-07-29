package socat

import (
	"fmt"
	"os"

	"get.porter.sh/porter/pkg/exec/builder"
)

var _ builder.ExecutableAction = Action{}
var _ builder.BuildableAction = Action{}

type Action struct {
	Name  string
	Steps []Step // using UnmarshalYAML so that we don't need a custom type per action
}

// MarshalYAML converts the action back to a YAML representation
// install:
//   socat:
//     ...
func (a Action) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{a.Name: a.Steps}, nil
}

// MakeSteps builds a slice of Step for data to be unmarshaled into.
func (a Action) MakeSteps() interface{} {
	return &[]Step{}
}

// UnmarshalYAML takes any yaml in this form
// ACTION:
// - socat: ...
// and puts the steps into the Action.Steps field
func (a *Action) UnmarshalYAML(unmarshal func(interface{}) error) error {
	results, err := builder.UnmarshalAction(unmarshal, a)
	if err != nil {
		return err
	}

	for actionName, action := range results {
		a.Name = actionName
		for _, result := range action {
			step := result.(*[]Step)
			a.Steps = append(a.Steps, *step...)
		}
		break // There is only 1 action
	}
	return nil
}

func (a Action) GetSteps() []builder.ExecutableStep {
	// Go doesn't have generics, nothing to see here...
	steps := make([]builder.ExecutableStep, len(a.Steps))
	for i := range a.Steps {
		steps[i] = a.Steps[i]
	}

	return steps
}

type Step struct {
	Instruction `yaml:"socat"`
}

// Actions is a set of actions, and the steps, passed from Porter.
type Actions []Action

// UnmarshalYAML takes chunks of a porter.yaml file associated with this mixin
// and populates it on the current action set.
// install:
//   socat:
//     ...
//   socat:
//     ...
// upgrade:
//   socat:
//     ...
func (a *Actions) UnmarshalYAML(unmarshal func(interface{}) error) error {
	results, err := builder.UnmarshalAction(unmarshal, Action{})
	if err != nil {
		return err
	}

	for actionName, action := range results {
		for _, result := range action {
			s := result.(*[]Step)
			*a = append(*a, Action{
				Name:  actionName,
				Steps: *s,
			})
		}
	}
	return nil
}

var _ builder.HasOrderedArguments = Instruction{}
var _ builder.ExecutableStep = Instruction{}
var _ builder.StepWithOutputs = Instruction{}

type Address struct {
	Address string   `yaml:"address,omitempty"`
	Options    []string `yaml:"options,omitempty"`
}

type Instruction struct {
	Name              string         `yaml:"name"`
	Description       string         `yaml:"description"`
	WorkingDir        string         `yaml:"dir,omitempty"`
	LeftAddress       Address        `yaml:"leftAddress"`
	RightAddress      Address        `yaml:"rightAddress"`
	SocatOptions      []string       `yaml:"socatOptions"`
}

func (s Instruction) GetCommand() string {
	return "porter_socat"
}

func (s Instruction) GetWorkingDir() string {
	return s.WorkingDir
}

func (s Instruction) GetArguments() []string {
	return []string{
		s.LeftAddress.String(),
		s.RightAddress.String(),
	}
}

func (s Instruction) GetSuffixArguments() []string {
	return s.SocatOptions
}

func (s Instruction) GetFlags() builder.Flags {
	return nil
}

func (s Instruction) SuppressesOutput() bool {
	return true
}

func (s Instruction) GetOutputs() []builder.Output {
	return nil
}

func (s Instruction) Async() bool {
	return true
}

func (a Address) String() string {
	str := a.Address
	if len(a.Options) > 0 {
		str += "," + a.Options[0]
	}

	for i := 1; i < len(a.Options); i++ {
		str += "," + a.Options[i]
	}
	fmt.Fprintf(os.Stderr, str)
	return str
}
