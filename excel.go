package tool

import (
	"errors"
	"github.com/tealeg/xlsx"
	"path/filepath"
)

func ExportExcel(filePath, sheetName string, head []string, data [][]string) error {
	if !IsExist(filePath) {
		return errors.New("文件不存在")
	}
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

/**
读取excel
filePath 文件路径
sheetNum 工作区从0开始
*/
func ReadExcel(filePath string, sheetNum int) ([][]string, error) {
	if !IsExist(filePath) {
		return nil, errors.New("文件不存在")
	}
	xf, err := xlsx.OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	if len(xf.Sheets) < sheetNum {
		return nil, errors.New("sheetNum 错误！")
	}
	data := make([][]string, 0, 20)
	for _, row := range xf.Sheets[sheetNum].Rows {
		cells := make([]string, 0, 10)
		for _, cell := range row.Cells {
			cells = append(cells, cell.String())
		}
		data = append(data, cells)
	}
	return data, nil
}

/**
获取下载excel文件头
*/
func GetDownExcelHeader(fileName string) map[string]string {
	headerMap := map[string]string{
		"Content-Disposition": "attachment; filename=" + fileName,
		"Accept-Ranges":       "bytes",
		"Cache-Control":       "must-revalidate, post-check=0, pre-check=0",
		"Pragma":              "no-cache",
		"Expires":             "0",
	}
	return headerMap
}
