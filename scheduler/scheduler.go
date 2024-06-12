package scheduler

type Schelduler interface {
	SelectCandidateNodes()
	Score()
	Pick()
}
