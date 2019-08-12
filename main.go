package main

import (
	"encoding/csv"
	"os"

	"github.com/trongcuong1710/facebook_tools/facebook"

	"github.com/sirupsen/logrus"
	"github.com/trongcuong1710/facebook_tools/config"
	"github.com/trongcuong1710/facebook_tools/constant"
)

func main() {
	fileName := "result.csv"
	file, err := os.Create(config.GetExportDirectory() + fileName)
	if err != nil {
		logrus.Warn(constant.ErrorLogTag, err, "os.Create")
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	comments, err := facebook.DefaultClient.GetComments("2391535717726871")
	if err != nil {
		logrus.Warn(constant.ErrorLogTag, err, "facebook.DefaultClient.GetComments")
		return
	}

	for _, comment := range comments {
		row := []string{
			comment.ID,
			comment.Message,
			comment.CreatedAt,
		}

		if writeErr := writer.Write(row); writeErr != nil {
			logrus.Warn(constant.ErrorLogTag, err, "row", row, "writer.Write")
		}
	}
}
