package services

import (
	"encoding/json"

	"labelin/internal/models"
	"github.com/jung-kurt/gofpdf"
)

type Element struct {
	Type     string  `json:"type"`      // "text"
	Field    string  `json:"field"`     // "FullName", "Address", etc
	X        float64 `json:"x"`         // in mm relative to label top-left
	Y        float64 `json:"y"`         // in mm relative to label top-left
	FontSize float64 `json:"font_size"`
	Align    string  `json:"align"`     // "L", "C", "R"
}

func GeneratePDFLabels(guests []models.Guest, tpl models.LabelTemplate) (*gofpdf.Fpdf, error) {
	// Parse elements
	var elements []Element
	if err := json.Unmarshal([]byte(tpl.ElementsJSON), &elements); err != nil {
		// Use default elements if empty or invalid
		elements = []Element{
			{Type: "text", Field: "PrefixTitle", X: 5, Y: 5, FontSize: 10, Align: "L"},
			{Type: "text", Field: "FullName", X: 5, Y: 10, FontSize: 12, Align: "L"},
			{Type: "text", Field: "Institution", X: 5, Y: 15, FontSize: 10, Align: "L"},
			{Type: "text", Field: "Address", X: 5, Y: 20, FontSize: 10, Align: "L"},
			{Type: "text", Field: "City", X: 5, Y: 25, FontSize: 10, Align: "L"},
		}
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(tpl.MarginLeft, tpl.MarginTop, tpl.MarginLeft)
	pdf.SetAutoPageBreak(false, 0)
	pdf.AddPage()

	col := 0
	row := 0

	startX := tpl.MarginLeft
	startY := tpl.MarginTop

	for i, guest := range guests {
		x := startX + float64(col)*(tpl.Width+tpl.SpacingX)
		y := startY + float64(row)*(tpl.Height+tpl.SpacingY)

		// Draw Label Border (Optional, for debugging or actual cutting lines)
		pdf.Rect(x, y, tpl.Width, tpl.Height, "D")

		// Print elements inside label
		for _, el := range elements {
			pdf.SetFont("Arial", "", el.FontSize)
			
			var text string
			switch el.Field {
			case "FullName":
				text = guest.FullName
			case "Institution":
				text = guest.Institution
			case "Address":
				text = guest.Address
			case "City":
				text = guest.City
			case "PrefixTitle":
				text = "Kepada Yth."
			case "Notes":
				text = guest.Notes
			}

			// GoFpdf positioning is tricky. SetXY moves cursor. Text prints from cursor down/right.
			pdf.SetXY(x+el.X, y+el.Y)
			
			alignStr := "L"
			if el.Align == "C" { alignStr = "C" }
			if el.Align == "R" { alignStr = "R" }

			// CellFormat(w, h, txtStr, borderStr, ln, alignStr, fill, link, linkStr)
			// Using w=0 to not restrict width, h=0
			pdf.CellFormat(tpl.Width-el.X, 0, text, "", 0, alignStr, false, 0, "")
		}

		col++
		if col >= tpl.Columns {
			col = 0
			row++
			if row >= tpl.Rows {
				row = 0
				// Move to next page if there are more guests
				if i < len(guests)-1 {
					pdf.AddPage()
				}
			}
		}
	}

	return pdf, nil
}
