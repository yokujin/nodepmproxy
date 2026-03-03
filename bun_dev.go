package nodepmproxy

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/rs/zerolog/log"
)

func (s *NodePMProxy) runBunDev() {
	if s.SitePath == "" {
		log.Fatal().Msg("path to site is not known. don't know how to start 'bun dev'")
	}

	log.Debug().
		Any("path", s.SitePath).
		Any("port", s.Port).
		Msg("running bun dev")

	var cmd *exec.Cmd
	if s.Framework == SVELTE {
		cmd = exec.Command("bun", "dev", "--port", fmt.Sprintf("%d", s.Port))
	} else {
		cmd = exec.Command("bun", "dev")
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = s.SitePath
	cmd.Env = append(
		os.Environ(),
		fmt.Sprintf("PORT=%d", s.Port),
	)

	err := cmd.Run()
	if err != nil {
		log.Fatal().Err(err).Msg("bun dev")
	}
	// fmt.Printf("%s", out)
}
