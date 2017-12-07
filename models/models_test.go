package models

import (
	"testing"
)

func TestAddListShouldAddItemToList(t *testing.T) {
	l := List{Name: "l"}
	l.Add("first").Add("second")

	if len(l.Items) != 2 {
		t.Errorf("expected %d items but found %d", 2, len(l.Items))
		t.FailNow()
	}
	if l.Items[0].Description != "first" {
		t.Errorf("expected description %s - found %s\n", "first", l.Items[0].Description)
		t.Fail()
	}

}

func TestAddListToUser(t *testing.T) {
	u := User{Name: "Kalpesh"}
	u.Add(List{Name: "work"}).Add(List{Name: "personal"})

	if len(u.Lists) != 2 {
		t.Errorf("expected %d number of items found %d", 2, len(u.Lists))
		t.FailNow()
	}
	lName := u.Lists[0].Name
	if lName != "work" {
		t.Errorf("expected first item to be %s  found %s", "work", lName)
	}
}
