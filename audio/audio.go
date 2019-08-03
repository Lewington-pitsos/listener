package audio

import (
	"fmt"
	"os/exec"

	"github.com/lewington/listener/assist"
)

// PlaySound plays the given mp3 file.
func PlaySound(filename string) {
	cmd := exec.Command(
		"play",
		assist.PathToAsset()+filename,
	)
	stdout, err := cmd.Output()
	fmt.Println(string(stdout))
	assist.Check(err)

}
