package integ

import "github.com/indeedeng/iwf-golang-sdk/iwf"

type basicWorkflow struct{}

func (b basicWorkflow) GetStates() []iwf.StateDef {
	return []iwf.StateDef{
		iwf.StartingStateDef(&basicWorkflowState1{}),
		iwf.NonStartingStateDef(&basicWorkflowState2{}),
	}
}

func (b basicWorkflow) GetPersistenceSchema() []iwf.PersistenceFieldDef {
	return nil
}

func (b basicWorkflow) GetCommunicationSchema() []iwf.CommunicationMethodDef {
	return nil
}

func (b basicWorkflow) GetWorkflowType() string {
	return ""
}
