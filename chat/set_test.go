package chat

import (
	"testing"

	"github.com/antonilol/ssh-chat/chat/message"
	"github.com/antonilol/ssh-chat/set"
)

func TestSet(t *testing.T) {
	var err error
	s := set.New()
	u := message.NewUser(message.SimpleID("foo"))

	if s.In(u.ID()) {
		t.Errorf("Set should be empty.")
	}

	err = s.Add(set.Itemize(u.ID(), u))
	if err != nil {
		t.Error(err)
	}

	if !s.In(u.ID()) {
		t.Errorf("Set should contain user.")
	}

	u2 := message.NewUser(message.SimpleID("bar"))
	err = s.Add(set.Itemize(u2.ID(), u2))
	if err != nil {
		t.Error(err)
	}

	err = s.Add(set.Itemize(u2.ID(), u2))
	if err != set.ErrCollision {
		t.Error(err)
	}

	size := s.Len()
	if size != 2 {
		t.Errorf("Set wrong size: %d (expected %d)", size, 2)
	}
}
