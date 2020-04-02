package tool

import (
	"errors"
	"github.com/tealeg/xlsx"
	"path/filepath"
)

func ExportExcel(filePath, sheetName string, head []string, data [][]string) error {
	//验证文件夹是否存在
	err := CreateFileByNot(filepath.Dir(filePath) + "/")
	if err != nil {
		return err
	}
	//文件类型验证
	extMap := map[string]int{".xls": 1, ".csv": 1, ".xlsx": 1}
	if _, ok := extMap[filepath.Ext(filePath)]; !ok {
		return errors.New("文件扩展名错误：只限.csv .xls .xlsx")
	}

	xf := xlsx.NewFile()
	sheet, err := xf.AddSheet(sheetName)
	if err != nil {
		return err
	}

	lenHead := len(head)
	lenData := len(data)
	if lenHead == 0 || lenData == 0 {
		return errors.New("标题列或数据列不能为空")
	}
	if lenHead != len(data[0]) {
		return errors.New("标题列与数据列，数量不一致")
	}
	//添加列标题
	headRow := sheet.AddRow()
	for _, title := range head {
		headRow.AddCell().Value = title
	}
	//添加列数据
	for _, row := range data {
		dataRow := sheet.AddRow()
		for _, col := range row {
			dataRow.AddCell().Value = col
		}
	}

	err = xf.Save(filePath)
	if err != nil {
		return err
	}
	return nil
}
