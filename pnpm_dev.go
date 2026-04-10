package nodepmproxy

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/rs/zerolog/log"
)

func (s *NodePMProxy) runPnpmDev() {
	if s.SitePath == "" {
		log.Fatal().Msg("path to site is not known. don't know how to start 'pnpm dev'")
	}

	log.Debug().
		Any("path", s.SitePath).
		Any("port", s.Port).
		Msg("running pnpm dev")

	var cmd *exec.Cmd
	if s.Framework == SVELTE {
		cmd = exec.Command("pnpm", "dev", "--port", fmt.Sprintf("%d", s.Port))
	} else {
		cmd = exec.Command("pnpm", "dev")
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
		log.Fatal().Err(err).Msg("pnpm dev")
	}
	// fmt.Printf("%s", out)
}
