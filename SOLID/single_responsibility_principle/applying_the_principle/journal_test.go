package applying_the_principle

import "testing"

func Test_ApplyingThePrinciple(t *testing.T) {
	j := Journal{}
	j.AddEntry("I cried today.")
	j.AddEntry("I ate a bug.")
	t.Log(j.String())

	p := Persistence{}
	p.SaveToFile(&j, "journal.txt")
}
