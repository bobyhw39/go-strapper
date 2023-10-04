package fsmutils

func AsEnterState(stateName string) string {
	return "enter_" + stateName
}
