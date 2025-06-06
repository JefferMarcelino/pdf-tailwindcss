package browserless

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pdf-tailwindcss/internal/domain"
)

type BrowserlessPDFGenerator struct {
	token              string
	host               string
	tailwindcssFileUrl string
}

func NewBrowserlessPDFGenerator(host, token, tailwindcssFileUrl string) domain.PDFGenerator {
	return &BrowserlessPDFGenerator{token: token, host: host, tailwindcssFileUrl: tailwindcssFileUrl}
}

func (b *BrowserlessPDFGenerator) GeneratePDF(html string) ([]byte, error) {
	fullHTML, err := buildFinalHTML(b.tailwindcssFileUrl, html)
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

func buildFinalHTML(tailwindcssFileUrl, htmlContent string) (string, error) {
	fullHTML := fmt.Sprintf(`
		<html>
			<head>
				<meta charset="utf-8">
				<link href="%s" rel="stylesheet">
			</head>
			<body>
				%s
			</body>
		</html>
	`, tailwindcssFileUrl, htmlContent)

	return fullHTML, nil
}
