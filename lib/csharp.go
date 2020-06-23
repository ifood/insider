package lib

import (
	"fmt"
	"log"
	"strconv"
	"io/ioutil"
	"path/filepath"
	"github.com/insidersec/insider/analyzers"
	"github.com/insidersec/insider/lexer"
	"github.com/insidersec/insider/models"
)

func AnalyzeCSharpSourceCode(dirname string, report *models.SASTReport) error {
	files, rules, err := LoadsFilesAndRules(dirname, "csharp")

	if err != nil {
		return err
	}

	appSize, err := analyzers.GetUnpackedAppSize(dirname)

	if err != nil {
		return err
	}

	report.Info.Size = fmt.Sprintf("%s MB", strconv.Itoa(appSize))

	for _, file := range files {
		fileContent, err := ioutil.ReadFile(filepath.Clean(file))

		if err != nil {
			return err
		}

		fileForAnalyze := lexer.NewInputFile(dirname, file, fileContent)

		report.Info.NumberOfLines = report.Info.NumberOfLines + len(fileForAnalyze.NewlineIndexes)

		fileSummary := analyzers.AnalyzeFile(fileForAnalyze, rules)

		for _, finding := range fileSummary.Findings {
			vulnerability := ConvertFindingToReport(
				fileForAnalyze.Name,
				fileForAnalyze.DisplayName,
				finding,
			)

			report.Vulnerabilities = append(report.Vulnerabilities, vulnerability)
		}
	}

	log.Printf("Scanned %d lines", report.Info.NumberOfLines)

	return nil
}
