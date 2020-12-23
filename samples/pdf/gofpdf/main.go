package main

import (
	"log"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	err := generateSimplePDF()
	if err != nil {
		log.Fatal(err)
	}

	err = generateProtectedPDF()
	if err != nil {
		log.Fatal(err)
	}
}

func generateSimplePDF() error {
	// Quickstart
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world")

	if err := pdf.OutputFileAndClose("hello.pdf"); err != nil {
		return err
	}

	if err := pdf.Error(); err != nil {
		return err
	}

	return nil
}

func generateProtectedPDF() error {
	// Taken from https://github.com/jung-kurt/gofpdf/blob/a2a0e7f8a28b2eabe1a32097f0071a0f715a8102/fpdf_test.go#L1667-L1678
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetProtection(gofpdf.CnProtectPrint, "123", "abc")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.Write(10, "Password-protected.")

	if err := pdf.OutputFileAndClose("protected.pdf"); err != nil {
		return err
	}

	if err := pdf.Error(); err != nil {
		return err
	}

	return nil
}
