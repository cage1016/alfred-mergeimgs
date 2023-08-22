package alfred

import (
	aw "github.com/deanishe/awgo"
)

const (
	Debug           = "DEBUG"
	QueryExtension  = "QUERY_EXTENSION"
	MaxQueryResults = "MAX_QUERY_RESULTS"
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
