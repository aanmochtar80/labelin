package services

import (
	"fmt"
	"mime/multipart"
	"strings"

	"labelin/internal/models"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

type ExcelService struct {
	DB *gorm.DB
}

func NewExcelService(db *gorm.DB) *ExcelService {
	return &ExcelService{DB: db}
}

type ImportResult struct {
	Success int
	Failed  int
	Skipped int
	Errors  []string
}

func (s *ExcelService) ImportGuests(file *multipart.FileHeader) (ImportResult, error) {
	result := ImportResult{}

	src, err := file.Open()
	if err != nil {
		return result, err
	}
	defer src.Close()

	f, err := excelize.OpenReader(src)
	if err != nil {
		return result, fmt.Errorf("failed to open excel: %v", err)
	}
	defer f.Close()

	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		return result, fmt.Errorf("no sheets found in excel file")
	}

	rows, err := f.GetRows(sheets[0])
	if err != nil {
		return result, fmt.Errorf("failed to get rows: %v", err)
	}

	if len(rows) < 2 {
		return result, fmt.Errorf("file is empty or missing data rows")
	}

	// Basic Header Validation (Optional: make it robust based on your template)
	// Expected: Nama | Instansi | Alamat | Kota | Provinsi | Kode Pos | HP

	for i, row := range rows {
		if i == 0 { // Skip header
			continue
		}

		// Pad row if it has fewer columns than expected
		for len(row) < 7 {
			row = append(row, "")
		}

		fullName := strings.TrimSpace(row[0])
		if fullName == "" {
			result.Skipped++
			result.Errors = append(result.Errors, fmt.Sprintf("Row %d: Nama Kosong", i+1))
			continue
		}

		guest := models.Guest{
			FullName:    fullName,
			Institution: strings.TrimSpace(row[1]),
			Address:     strings.TrimSpace(row[2]),
			City:        strings.TrimSpace(row[3]),
			Province:    strings.TrimSpace(row[4]),
			PostalCode:  strings.TrimSpace(row[5]),
			Phone:       strings.TrimSpace(row[6]),
		}

		// Check duplicate by name and phone (or just name for simplicity)
		var existing models.Guest
		err := s.DB.Where("full_name = ? AND (phone = ? OR phone = '')", guest.FullName, guest.Phone).First(&existing).Error
		
		if err == gorm.ErrRecordNotFound {
			if err := s.DB.Create(&guest).Error; err != nil {
				result.Failed++
				result.Errors = append(result.Errors, fmt.Sprintf("Row %d: Failed to save %s", i+1, fullName))
			} else {
				result.Success++
			}
		} else if err == nil {
			// Update if exists
			guest.ID = existing.ID
			guest.CreatedAt = existing.CreatedAt
			if err := s.DB.Save(&guest).Error; err != nil {
				result.Failed++
				result.Errors = append(result.Errors, fmt.Sprintf("Row %d: Failed to update %s", i+1, fullName))
			} else {
				result.Success++
			}
		} else {
			result.Failed++
			result.Errors = append(result.Errors, fmt.Sprintf("Row %d: DB Error %v", i+1, err))
		}
	}

	return result, nil
}
