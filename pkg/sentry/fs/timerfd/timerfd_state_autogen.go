// automatically generated by stateify.

package timerfd

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (t *TimerOperations) StateTypeName() string {
	return "pkg/sentry/fs/timerfd.TimerOperations"
}

func (t *TimerOperations) StateFields() []string {
	return []string{
		"timer",
		"val",
	}
}

func (t *TimerOperations) beforeSave() {}

func (t *TimerOperations) StateSave(stateSinkObject state.Sink) {
	t.beforeSave()
	if !state.IsZeroValue(&t.events) {
		state.Failf("events is %#v, expected zero", &t.events)
	}
	stateSinkObject.Save(0, &t.timer)
	stateSinkObject.Save(1, &t.val)
}

func (t *TimerOperations) afterLoad() {}

func (t *TimerOperations) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &t.timer)
	stateSourceObject.Load(1, &t.val)
}

func init() {
	state.Register((*TimerOperations)(nil))
}