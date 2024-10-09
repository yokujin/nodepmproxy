//go:build !prod

package nodepmproxy

import (
	"fmt"
	"net/http"
	"time"

	"git.0x21.ru/yokujin/nodepmproxy/wsp"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func (s *NodePMProxy) SetupEcho(e *echo.Echo) {

	go s.runYarnDev()

	wp, err := wsp.NewProxy(fmt.Sprintf("ws://localhost:%d/_nuxt/", s.Port), func(r *http.Request) error {
		// // Permission to verify
		// r.Header.Set("Cookie", "----")
		// // Source of disguise
		// r.Header.Set("Origin", fmt.Sprintf("http://localhost:%d", 8080))
		return nil
	})
	if err != nil {
		log.Fatal().Err(err).Msg("error creating websocket proxy")
	}

	e.Any("/_nuxt", echo.WrapHandler(wp))
	e.Any("/_nuxt/", echo.WrapHandler(wp))
	e.Any("/*", s.GetOr404())

}

func (s *NodePMProxy) GetOr404() echo.HandlerFunc {
	return func(c echo.Context) error {
		// log.Debug().Msg("======================= New request ===================================")
		// log.Debug().Msgf("%#v\n", c)
		// log.Debug().Msgf("---------\nREQUEST: %#v\n", c.Request())

		rewrittenUrl := fmt.Sprintf("http://localhost:%d%s", s.Port, c.Request().RequestURI)
		// log.Debug().Msgf("rewritten url: %s", rewrittenUrl)

		client := &http.Client{
			// CheckRedirect: redirectPolicyFunc,
		}

		// create new request
		req, err := http.NewRequest(c.Request().Method, rewrittenUrl, c.Request().Body)
		if err != nil {
			log.Error().Err(err).Msg("creating request")
		}

		// Move all incoming headers into proxied request
		for k, vals := range c.Request().Header {
			for _, val := range vals {
				req.Header.Add(k, val)
			}
		}

		// make request
		var resp *http.Response

		for {
			resp, err = client.Do(req)
			if err != nil {
				log.Error().Err(err).Msg("client.Do")
				time.Sleep(time.Second)
				continue
			}
			break
		}
		defer func() {
			if resp != nil {
				resp.Body.Close()
			}
		}()
		// log.Debug().Msgf("---------\nAPI RESPONSE: %#v\n", resp)

		return c.Stream(resp.StatusCode, resp.Header.Get("Content-Type"), resp.Body)
	}
}