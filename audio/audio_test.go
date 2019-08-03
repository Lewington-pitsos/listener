package audio

import (
	"testing"

	"github.com/lewington/listener/testhelp"
)

func TestAudio(t *testing.T) {
	testhelp.SkipIfShort(t)

	PlaySound("whoosh.mp3")
}
