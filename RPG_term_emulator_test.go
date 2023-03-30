package RPG_term_emulator

import (
	"testing"
)

func test_login(t testing.T) {
	want := 0
	if _,got := login_check("admin","admin","Jesus"); got != want {
		t.Errorf("login_check = %d, but wanted %d",got,want)
	}
}
