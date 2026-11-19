package application

type Exporter interface {
	Export(path string, rows []ExportRow) error
}
