package advance

import (
	"fmt"

	"github.com/bregydoc/gtranslate"
	"github.com/gen2brain/go-fitz"
	"github.com/jung-kurt/gofpdf"
)

func ConvertToEnglish() {
	// Open the PDF file
	doc, err := fitz.New("pkg/advance/german.pdf")
	if err != nil {
		fmt.Printf("Error opening PDF: %v\n", err)
		return
	}
	defer doc.Close()

	// Create a new PDF
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Iterate through pages
	for n := 0; n < doc.NumPage(); n++ {
		// Extract text
		text, err := doc.Text(n)
		if err != nil {
			fmt.Printf("Error extracting text from page %d: %v\n", n+1, err)
			continue
		}

		// Translate text
		translatedText, err := translateText(text, "de", "en")
		if err != nil {
			fmt.Printf("Error translating text from page %d: %v\n", n+1, err)
			continue
		}

		// Add a new page to the output PDF
		pdf.AddPage()

		// Add translated text to the page
		pdf.SetFont("Arial", "", 12)
		pdf.MultiCell(190, 10, translatedText, "0", "L", false)
	}

	// Save the new PDF
	err = pdf.OutputFileAndClose("output.pdf")
	if err != nil {
		fmt.Printf("Error saving PDF: %v\n", err)
		return
	}

	fmt.Println("PDF created successfully")
}

func translateText(text, from, to string) (string, error) {
	translated, err := gtranslate.TranslateWithParams(
		text,
		gtranslate.TranslationParams{
			From: from,
			To:   to,
		},
	)
	if err != nil {
		return "", err
	}
	return translated, nil
}
