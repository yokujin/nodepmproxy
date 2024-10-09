package nodepmproxy

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/rs/zerolog/log"
)

func (s *NodePMProxy) runYarnDev() {
	if s.SitePath == "" {
		log.Fatal().Msg("path to site is not known. don't know how to start 'yarn dev'")
	}

	log.Debug().Any("path", s.SitePath).Any("port", s.Port).Msg("running yarn dev")

	err := os.Chdir(s.SitePath)
	if err != nil {
		log.Fatal().Err(err).Msg("chdir to site dir")
	}

	cmd := exec.Command("yarn", "dev")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	additionalEnv := fmt.Sprintf("PORT=%d", s.Port)
	newEnv := append(os.Environ(), additionalEnv)
	cmd.Env = newEnv

	err = cmd.Run()
	if err != nil {
		log.Fatal().Err(err).Msg("yarn dev")
	}
	// fmt.Printf("%s", out)
}
