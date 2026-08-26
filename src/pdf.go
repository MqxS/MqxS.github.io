package src

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// GeneratePDF prints the generated portfolio HTML to a PDF using
// an installed Chromium-based browser (Chrome, Edge, or Chromium).
func GeneratePDF(htmlPath, pdfPath string) error {
	htmlAbs, err := filepath.Abs(htmlPath)
	if err != nil {
		return err
	}

	pdfAbs, err := filepath.Abs(pdfPath)
	if err != nil {
		return err
	}

	browser, err := findBrowser()
	if err != nil {
		return err
	}

	// file:/// URL required by Chromium.
	fileURL := "file:///" + filepath.ToSlash(htmlAbs)

	args := []string{
		"--headless",
		"--disable-gpu",
		"--allow-file-access-from-files",
		"--no-pdf-header-footer",
		"--print-to-pdf=" + pdfAbs,
		fileURL,
	}

	// Chromium refuses to run as root on Linux unless the sandbox is disabled.
	// This mainly helps containers/CI and does not affect normal desktop use.
	if runtime.GOOS == "linux" && os.Geteuid() == 0 {
		args = append([]string{"--no-sandbox"}, args...)
	}

	cmd := exec.Command(browser, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to generate PDF: %w", err)
	}

	fmt.Println("Generated PDF:", pdfAbs)
	return nil
}

func findBrowser() (string, error) {
	var candidates []string

	switch runtime.GOOS {
	case "windows":
		candidates = []string{
			`C:\Program Files\Google\Chrome\Application\chrome.exe`,
			`C:\Program Files (x86)\Google\Chrome\Application\chrome.exe`,
			`C:\Program Files\Microsoft\Edge\Application\msedge.exe`,
			`C:\Program Files (x86)\Microsoft\Edge\Application\msedge.exe`,
		}

	case "darwin":
		candidates = []string{
			"/Applications/Google Chrome.app/Contents/MacOS/Google Chrome",
			"/Applications/Microsoft Edge.app/Contents/MacOS/Microsoft Edge",
			"/Applications/Chromium.app/Contents/MacOS/Chromium",
		}

	default:
		candidates = []string{
			"google-chrome",
			"google-chrome-stable",
			"chromium",
			"chromium-browser",
			"microsoft-edge",
		}
	}

	for _, candidate := range candidates {
		if filepath.IsAbs(candidate) {
			if _, err := os.Stat(candidate); err == nil {
				return candidate, nil
			}
			continue
		}

		if path, err := exec.LookPath(candidate); err == nil {
			return path, nil
		}
	}

	return "", fmt.Errorf("Chrome, Edge, or Chromium was not found")
}
