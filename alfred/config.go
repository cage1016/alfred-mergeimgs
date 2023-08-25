package alfred

import (
	aw "github.com/deanishe/awgo"
)

const (
	Debug           = "DEBUG"
	QueryExtension  = "QUERY_EXTENSION"
	MaxQueryResults = "MAX_QUERY_RESULTS"
	Gravity         = "GRAVITY"
	Background      = "BACKGROUND"
	Offset          = "OFFSET"
)

func GetDebug(wf *aw.Workflow) bool {
	return wf.Config.GetBool(Debug)
}

func GetQueryExtension(wf *aw.Workflow) string {
	return wf.Config.Get(QueryExtension)
}

func GetMaxQueryResults(wf *aw.Workflow) int {
	return wf.Config.GetInt(MaxQueryResults)
}

func GetGravity(wf *aw.Workflow) string {
	return wf.Config.Get(Gravity)
}

func GetBackground(wf *aw.Workflow) string {
	return wf.Config.Get(Background)
}

func GetOffset(wf *aw.Workflow) string {
	return wf.Config.Get(Offset)
}
