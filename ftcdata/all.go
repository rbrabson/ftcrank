package ftcdata

func LoadAll() {
	LoadAdvancements()
	LoadAwards()
	LoadEvents()
	LoadMatches()
	LoadRankings()
	LoadSchedules()
	LoadHybridSchedules()
	LoadScores()
	LoadTeams()
}

func RetrieveAll() {
	RetrieveAdvancements()
	RetrieveAwards()
	RetrieveEvents()
	RetrieveMatches()
	RetrieveRankings()
	RetrieveSchedules()
	RetrieveHybridSchedules()
	RetrieveScores()
	RetrieveTeams()
}

func UpdateAll(eventCode string) {
	UpdateEvents(eventCode)
	UpdateAdvancements(eventCode)
	UpdateAwards(eventCode)
	UpdateMatches(eventCode)
	UpdateRankings(eventCode)
	UpdateSchedules(eventCode)
	UpdateHybridSchedules(eventCode)
	UpdateScores(eventCode)
}
