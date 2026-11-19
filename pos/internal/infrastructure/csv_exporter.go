package infrastructure

import (
	"encoding/csv"
	"fmt"
	"os"
	"pos/internal/application"
	"strconv"
)

type CSVExporter struct{}

func NewCSVExporter() *CSVExporter {
	return &CSVExporter{}
}

func (e *CSVExporter) Export(path string, rows []application.ExportRow) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	header := []string{
		"Student ID",
		"Student Name",
		"Criterion",
		"Score",
		"Final Score",
	}

	err = writer.Write(header)
	if err != nil {
		return fmt.Errorf("failed to write CSV header: %w", err)
	}

	for _, row := range rows {
		record := []string{
			row.StudentID,
			row.StudentName,
			row.Criterion,
			strconv.FormatFloat(row.Score, 'f', 2, 64),
			strconv.FormatFloat(row.FinalScore, 'f', 2, 64),
		}

		err = writer.Write(record)
		if err != nil {
			return fmt.Errorf("failed to write CSV row: %w", err)
		}
	}

	writer.Flush()

	err = writer.Error()
	if err != nil {
		return fmt.Errorf("failed to flush CSV file: %w", err)
	}

	return nil
}
