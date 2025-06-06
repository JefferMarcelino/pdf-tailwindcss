package application

import "pdf-tailwindcss/internal/domain"

type PDFService struct {
	pdfGenerator domain.PDFGenerator
}

func NewPDFService(g domain.PDFGenerator) *PDFService {
	return &PDFService{pdfGenerator: g}
}

func (s *PDFService) GeneratePDF(html string) ([]byte, error) {
	return s.pdfGenerator.GeneratePDF(html)
}
