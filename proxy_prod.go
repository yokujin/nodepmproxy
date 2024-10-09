//go:build prod

package nodepmproxy

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func (s *NodePMProxy) SetupEcho(e *echo.Echo) {
	if s.Embedded {
		e.Any("/*", s.GetFromEmbeddedOr404())
	} else {
		e.Any("/*", s.GetOr404())
	}
}

func (s *NodePMProxy) GetOr404() echo.HandlerFunc {
	// site_path := viper.GetString("site_path")
	// assetHandler := echo.WrapHandler(http.FileServer(http.Dir(site_path)))
	// return assetHandler

	return func(c echo.Context) error {
		path := c.Param("*")
		// log.Debug().Any("path", path).Msg("nuxt get or 404")

		fullPath := filepath.Join(s.SitePath, path)

		finfo, err := os.Stat(fullPath)
		if os.IsNotExist(err) {
			return c.NoContent(http.StatusNotFound)
		}

		if finfo.IsDir() {
			fullPath = filepath.Join(fullPath, "index.html")
		}

		return c.File(fullPath)
	}
}

func (s *NodePMProxy) GetFromEmbeddedOr404() echo.HandlerFunc {
	// fsys, err := fs.Sub(ui.UI, ".output/public")
	// if err != nil {
	// 	log.Error().Err(err).Msg("error creating embedded fs")
	// }

	// assetHandler := echo.WrapHandler(http.StripPrefix("/ui", http.FileServer(http.FS(fsys))))
	assetHandler := echo.WrapHandler(http.FileServer(http.FS(s.EmbeddedFS)))

	return assetHandler
}
