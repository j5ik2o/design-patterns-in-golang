package mediator

type Mediator interface {
	AddColleague(colleague Colleague)
	ColleagueChanged(colleagueUpdated Colleague, msg string)
}
