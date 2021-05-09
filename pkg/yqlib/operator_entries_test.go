package yqlib

import (
	"testing"
)

var entriesOperatorScenarios = []expressionScenario{
	{
		description: "to_entries Map",
		document:   `{a: 1, b: 2}`,
		expression: `to_entries`,
		expected: []string{
			"D0, P[], (!!seq)::- key: a\n  value: 1\n- key: b\n  value: 2\n",
		},
	},
	{
		description: "to_entries Array",
		document:   `[a, b]`,
		expression: `to_entries`,
		expected: []string{
			"D0, P[], (!!seq)::- key: 0\n  value: a\n- key: 1\n  value: b\n",
		},
	},
	{
		description: "from_entries map",
		document:   `{a: 1, b: 2}`,
		expression: `to_entries | from_entries`,
		expected: []string{
			"D0, P[], (!!map)::a: 1\nb: 2\n",
		},
	},
	{
		description: "from_entries with numeric key indexes",
		subdescription: "from_entries always creates a map, even for numeric keys",
		document:   `[a,b]`,
		expression: `to_entries | from_entries`,
		expected: []string{
			"D0, P[], (!!map)::0: a\n1: b\n",
		},
	},
}

func TestEntriesOperatorScenarios(t *testing.T) {
	for _, tt := range entriesOperatorScenarios {
		testScenario(t, &tt)
	}
	documentScenarios(t, "Entries", entriesOperatorScenarios)
}