package main

type JobSite struct {
	subscibers []Observer
	vacancies  []string
}

func (j *JobSite) addVacancies(vacancy string) {
	j.vacancies = append(j.vacancies, vacancy)
	j.SendAll()
}
func (j *JobSite) removeVacancies(vacancy string) {
	j.vacancies = remove(j.vacancies, vacancy)
	j.SendAll()
}

// -------
func (j *JobSite) subscribe(subscriber Observer) {
	j.subscibers = append(j.subscibers, subscriber)
}

func (j *JobSite) unsubscribe(subscriber Observer) {
	j.subscibers = removeO(j.subscibers, subscriber)
}
func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
func removeO(s []Observer, r Observer) []Observer {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func (j *JobSite) SendAll() {
	for _, subscriber := range j.subscibers {
		subscriber.HandleEvent(j.vacancies)
	}
}
