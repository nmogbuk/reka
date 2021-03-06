// Code generated by "enumer -type=Action ./rules/rule.go"; DO NOT EDIT.

//
package rules

import (
	"fmt"
)

const _ActionName = "DoNothingStopResumeDestroy"

var _ActionIndex = [...]uint8{0, 9, 13, 19, 26}

func (i Action) String() string {
	if i < 0 || i >= Action(len(_ActionIndex)-1) {
		return fmt.Sprintf("Action(%d)", i)
	}
	return _ActionName[_ActionIndex[i]:_ActionIndex[i+1]]
}

var _ActionValues = []Action{0, 1, 2, 3}

var _ActionNameToValueMap = map[string]Action{
	_ActionName[0:9]:   0,
	_ActionName[9:13]:  1,
	_ActionName[13:19]: 2,
	_ActionName[19:26]: 3,
}

// ActionString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ActionString(s string) (Action, error) {
	if val, ok := _ActionNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Action values", s)
}

// ActionValues returns all values of the enum
func ActionValues() []Action {
	return _ActionValues
}

// IsAAction returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Action) IsAAction() bool {
	for _, v := range _ActionValues {
		if i == v {
			return true
		}
	}
	return false
}
