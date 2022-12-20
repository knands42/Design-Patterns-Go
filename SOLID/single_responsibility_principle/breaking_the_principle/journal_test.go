package breaking_the_principle

import "testing"

func Test_BreakingThePrinciple(t *testing.T) {
	j := Journal{}
	j.AddEntry("I cried today.")
	j.AddEntry("I ate a bug.")
	t.Log(j.String())

	j.Save("journal.txt")
}
