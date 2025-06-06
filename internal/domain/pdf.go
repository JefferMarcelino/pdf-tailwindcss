package domain

type PDFGenerator interface {
	GeneratePDF(html string) ([]byte, error)
}
