package browserless

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"pdf-tailwindcss/internal/domain"
)

type BrowserlessPDFGenerator struct {
	token string
	host  string
}

func NewBrowserlessPDFGenerator(host, token string) domain.PDFGenerator {
	return &BrowserlessPDFGenerator{token: token, host: host}
}

func (b *BrowserlessPDFGenerator) GeneratePDF(html string) ([]byte, error) {
	fullHTML, err := buildFinalHTML(html)
	if err != nil {
		return nil, err
	}

	margin := "32px"
	payload := map[string]any{
		"html": fullHTML,
		"options": map[string]any{
			"printBackground": true,
			"format":          "A4",
			"margin": map[string]string{
				"top":    margin,
				"bottom": margin,
				"left":   margin,
				"right":  margin,
			},
		},
	}

	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/pdf?token=%s", b.host, b.token)

	resp, err := http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("browserless request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("browserless error: %s", string(body))
	}

	return io.ReadAll(resp.Body)
}

func buildFinalHTML(htmlContent string) (string, error) {
	tailwindCSS, err := os.ReadFile("./tailwind.css")
	if err != nil {
		return "", fmt.Errorf("failed to read tailwind css: %w", err)
	}

	fullHTML := fmt.Sprintf(`
		<html>
			<head>
				<meta charset="utf-8">
				<style>%s</style>
			</head>
			<body>
				%s
			</body>
		</html>
	`, string(tailwindCSS), htmlContent)

	return fullHTML, nil
}
