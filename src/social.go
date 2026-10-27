package src

import (
	"fmt"
	"image/png"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// GenerateSocialImage rasterizes a social preview SVG for link preview crawlers.
func GenerateSocialImage(svgPath, pngPath string) error {
	svgAbs, err := filepath.Abs(svgPath)
	if err != nil {
		return err
	}
	pngAbs, err := filepath.Abs(pngPath)
	if err != nil {
		return err
	}
	browser, err := findBrowser()
	if err != nil {
		return err
	}
	filePath := filepath.ToSlash(svgAbs)
	if !strings.HasPrefix(filePath, "/") {
		filePath = "/" + filePath
	}
	fileURL := (&url.URL{Scheme: "file", Path: filePath}).String()
	args := []string{
		"--headless",
		"--disable-gpu",
		"--hide-scrollbars",
		"--force-device-scale-factor=1",
		"--window-size=1200,630",
		"--screenshot=" + pngAbs,
		fileURL,
	}
	if runtime.GOOS == "linux" && os.Geteuid() == 0 {
		args = append([]string{"--no-sandbox"}, args...)
	}
	if output, err := exec.Command(browser, args...).CombinedOutput(); err != nil {
		return fmt.Errorf("render preview: %w\n%s", err, output)
	}
	f, err := os.Open(pngAbs)
	if err != nil {
		return err
	}
	defer f.Close()
	config, err := png.DecodeConfig(f)
	if err != nil {
		return err
	}
	if config.Width != 1200 || config.Height != 630 {
		return fmt.Errorf("unexpected preview dimensions: %dx%d", config.Width, config.Height)
	}
	return nil
}
