package iwf

type StateDecision struct {
	NextStates []StateMovement
}

func SingleNextState(state WorkflowState, input interface{}) *StateDecision {
	return &StateDecision{
		NextStates: []StateMovement{
			{
				NextStateId:    GetFinalWorkflowStateId(state),
				NextStateInput: input,
			},
		},
	}
}

func MultiNextStates(states ...WorkflowState) *StateDecision {
	var movements []StateMovement
	for _, st := range states {
		movements = append(movements, StateMovement{
			NextStateId: GetFinalWorkflowStateId(st),
		})
	}
	return &StateDecision{
		NextStates: movements,
	}
}

func MultiNextStatesWithInput(movements ...StateMovement) *StateDecision {
	return &StateDecision{
		NextStates: movements,
	}
}

func MultiNextStatesByStateIds(nextStateIds ...string) *StateDecision {
	var movements []StateMovement
	for _, id := range nextStateIds {
		movements = append(movements, StateMovement{
			NextStateId: id,
		})
	}
	return &StateDecision{
		NextStates: movements,
	}
}

// DeadEnd means no next step for this thread. It is essentially graceful completion without a workflow result
var DeadEnd = &StateDecision{}

var ForceFailingWorkflow = &StateDecision{
	NextStates: []StateMovement{
		{
			NextStateId: ForceFailingWorkflowStateId,
		},
	},
}

var GracefulCompletingWorkflow = GracefulCompleteWorkflow(nil)

func GracefulCompleteWorkflow(output interface{}) *StateDecision {
	return &StateDecision{
		NextStates: []StateMovement{
			{
				NextStateId:    GracefulCompletingWorkflowStateId,
				NextStateInput: output,
			},
		},
	}
}

var ForceCompletingWorkflow = ForceCompleteWorkflow(nil)

func ForceCompleteWorkflow(output interface{}) *StateDecision {
	return &StateDecision{
		NextStates: []StateMovement{
			{
				NextStateId:    ForceCompletingWorkflowStateId,
				NextStateInput: output,
			},
		},
	}
}
