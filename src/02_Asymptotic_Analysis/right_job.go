package main

type job struct {
	startDate int
	endDate int
}

func findRightJob(jobs []job) []job {
	return []job{}	
}

func main() {
	tests := []job{
		{
			startDate: 0,
			endDate: 3,
		},
		{
			startDate: 1,
			endDate: 5,
		},
	}

	_ = findRightJob(tests)

}