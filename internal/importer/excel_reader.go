package importer

import (
	"monitoring-edc/internal/monitoring"
	"time"

	"github.com/xuri/excelize/v2"
)

func ReadExcel(path string) ([][]string, error) {

	file, err := excelize.OpenFile(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	sheetName := file.GetSheetName(0)
	rows, err := file.GetRows(sheetName)

	if err != nil {
		return nil, err
	}
	return rows, nil
}

func ParseMonitoring(path string) ([]monitoring.Monitoring, error) {
	file, err := excelize.OpenFile(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	sheet := file.GetSheetName(0)

	rows, err := file.GetRows(sheet)
	if err != nil {
		return nil, err
	}

	var result []monitoring.Monitoring

	for i, row := range rows {
		// skip header
		if i == 0 {
			continue
		}

		if len(row) < 12 {
			continue
		}

		tgl, _ := time.Parse("2006-01-02", row[0])

		data := monitoring.Monitoring{
			Tgl:        tgl,
			Kdcab:      row[1],
			Cabang:     row[2],
			Kdtk:       row[3],
			Nama:       row[4],
			Station:    row[5],
			Cek:        row[6],
			IP:         row[7],
			EDCBCA:     row[8],
			EDCMandiri: row[9],
			EDCMTI:     row[10],
			EDCMDRMTI:  row[11],
		}

		result = append(result, data)
	}

	return result, nil
}
