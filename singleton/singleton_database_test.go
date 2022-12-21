package singleton

import "testing"

type DummyDatabase struct {
	dummyData map[string]int
}

func (d *DummyDatabase) GetPopulation(name string) int {
	if len(d.dummyData) == 0 {
		d.dummyData = map[string]int{
			"alpha": 1,
			"beta":  2,
			"gamma": 3}
	}
	return d.dummyData[name]
}

func Test_GetSingletonDatabase(t *testing.T) {
	db := GetSingletonDatabase()
	pop := db.GetPopulation("Seoul")
	if pop != 17500000 {
		t.Errorf("Seoul's population is %d", pop)
	}
}

func Test_GetTotalPopulation(t *testing.T) {
	cities := []string{"Seoul", "Mexico City"}
	tp := GetTotalPopulation(&DummyDatabase{}, cities)
	if tp != 17500000+17400000 {
		t.Errorf("Total population is %d", tp)
	}
}
