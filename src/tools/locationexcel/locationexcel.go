package main

import (
	"fmt"
	"github.com/micro/go-log"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	tool "tools/splitfile"
)

const (
	SOURCE_FILE_PATH = "F:\\软云\\06 AI\\28. 按省份除地理以外学科测评、错题次数统计\\20190509\\StatsResult"
	DEST_FILE_PATH   = "F:\\软云\\06 AI\\28. 按省份除地理以外学科测评、错题次数统计\\20190509\\StatsResult\\excel"
)

type course struct {
	courseShortName string
	courseName      string
}

func main() {
	courseMap := getCourseMap()
	files, _ := getAllFiles(SOURCE_FILE_PATH, ".txt", false)
	for _, file := range files {
		fileName := filepath.Base(file)
		fileSuffix := path.Ext(fileName)
		fileNameWithOutSuffix := strings.TrimSuffix(fileName, fileSuffix) //获取文件名
		fileIndex := strings.TrimLeft(fileNameWithOutSuffix, "course-")
		courseId, err := strconv.Atoi(fileIndex)
		if err != nil {
			log.Fatal("解析名称中的学科Id出错: ", file)
		}
		course, ok := courseMap[courseId]
		if !ok {
			continue
		}
		sheetName := course.courseName
		excelFileName := path.Join(DEST_FILE_PATH, fmt.Sprintf("%s.xlsx", course.courseShortName))
		if exist, _ := fileExists(excelFileName); !exist {
			xlFile := xlsx.NewFile()
			xlFile.AddSheet(sheetName)
			if err = xlFile.Save(excelFileName); err != nil {
				log.Fatal("save file error,", excelFileName, err.Error())
			}
		}
		xlFile, err := xlsx.OpenFile(excelFileName)
		if err != nil {
			log.Log("open file error, ", file)
			continue
		}
		s, exists := xlFile.Sheet[sheetName]
		if !exists {
			s, err = xlFile.AddSheet(sheetName)
		}
		lines, err := tool.ReadAllLine(file)
		writeSheetContent(lines, s)
		hiddenUnknownColumn(s)
		hiddenUnknownRow(s)
		xlFile.Save(excelFileName)
	}
}

const (
	DATA_ROW_BEGIN    = 2
	DATA_COLUMN_BEGIN = 2
)

func writeSheetContent(lines []string, sheet *xlsx.Sheet) error {
	for r_index, line := range lines {
		r := sheet.AddRow()
		sheet.SetColWidth(1, 1, 30)
		r.Height = 20
		items := strings.Split(line, "\t")
		for c_index, item := range items {
			c := r.AddCell()
			style := &xlsx.Style{
				Alignment: xlsx.Alignment{Horizontal: "center", Vertical: "center"},
				ApplyFill: true,
				Border:    xlsx.Border{Left: "thin", Right: "thin", Bottom: "thin", Top: "thin"},
			}
			if r_index < DATA_ROW_BEGIN {
				style.Fill = xlsx.Fill{PatternType: "solid", FgColor: "FFA6A6A6"}
			}
			if r_index >= DATA_ROW_BEGIN && c_index < DATA_ROW_BEGIN && items[0] == "知识" {
				style.Fill = xlsx.Fill{PatternType: "solid", FgColor: "FFC4D79B"}
			}
			if r_index >= DATA_ROW_BEGIN && c_index < DATA_ROW_BEGIN && items[0] == "技能" {
				style.Fill = xlsx.Fill{PatternType: "solid", FgColor: "FF8DB4E2"}
			}
			if r_index >= DATA_ROW_BEGIN && c_index < DATA_ROW_BEGIN && items[0] == "能力" {
				style.Fill = xlsx.Fill{PatternType: "solid", FgColor: "FFE26B0A"}
			}
			if r_index == 0 && c_index == 0 {
				c.Merge(1, 1)
			}
			if r_index == 0 && c_index >= DATA_COLUMN_BEGIN && c_index%2 == 0 {
				c.Merge(1, 0)
			}
			if r_index >= DATA_ROW_BEGIN && c_index >= DATA_COLUMN_BEGIN {
				i, err := strconv.Atoi(item)
				if err == nil {
					c.SetInt(i)
				} else {
					c.SetString(item)
				}
			} else {
				c.SetString(item)
			}
			c.SetStyle(style)
		}

	}
	return nil
}

func hiddenUnknownColumn(sheet *xlsx.Sheet) {
	if sheet.MaxCol <= DATA_COLUMN_BEGIN || sheet.MaxRow <= 1 {
		return
	}
	for i := 0; i < sheet.MaxCol; i++ {
		colVal := sheet.Rows[0].Cells[i].Value
		if strings.Contains(colVal, "未知(") || strings.Contains(colVal, "高新区") || strings.Contains(colVal, "青山湖区") {
			sheet.Col(i).Hidden = true
			sheet.Col(i + 1).Hidden = true
		}
	}
}

func hiddenUnknownRow(sheet *xlsx.Sheet) {
	if sheet.MaxCol < DATA_COLUMN_BEGIN || sheet.MaxRow <= DATA_ROW_BEGIN {
		return
	}
	for i := 0; i < sheet.MaxRow; i++ {
		colVal := sheet.Rows[i].Cells[1].Value
		if strings.Contains(colVal, "未知(") {
			sheet.Row(i).Hidden = true
		}
	}
}

func getCourseMap() map[int]course {
	courseMap := map[int]course{
		1:  {"语文", "初中语文"},
		2:  {"数学", "初中数学"},
		3:  {"英语", "初中英语"},
		4:  {"物理", "初中物理"},
		5:  {"化学", "初中化学"},
		6:  {"生物", "初中生物"},
		7:  {"历史", "初中历史"},
		8:  {"地理", "初中地理"},
		9:  {"语文", "高中语文"},
		10: {"数学", "高中数学"},
		11: {"英语", "高中英语"},
		12: {"物理", "高中物理"},
		13: {"化学", "高中化学"},
		14: {"生物", "高中生物"},
		15: {"历史", "高中历史"},
		16: {"地理", "高中地理"},
		17: {"政治", "高中政治"},
		19: {"政治", "初中政治"},
		20: {"奥数", "小学奥数"},
		21: {"语文", "小学语文"},
		22: {"数学", "小学数学"},
		23: {"英语", "小学英语"},
	}
	return courseMap
}

//获取指定目录下的所有文件,包含子目录下的文件
func getAllFiles(dirPth string, suffix string, containChild bool) (files []string, err error) {
	var dirs []string
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() && containChild { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			if childFiles, err := getAllFiles(dirPth+PthSep+fi.Name(), suffix, containChild); err == nil {
				files = append(files, childFiles...)
			}
		} else {
			// 过滤指定格式
			ok := strings.HasSuffix(fi.Name(), suffix)
			if ok {
				files = append(files, dirPth+PthSep+fi.Name())
			}
		}
	}
	//// 读取子目录下文件
	//for _, table := range dirs {
	//	temp, _ := getAllFiles(table)
	//	for _, temp1 := range temp {
	//		files = append(files, temp1)
	//	}
	//}
	return files, nil
}

func fileExists(path string) (bool, error) {
	s, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	if !s.IsDir() {
		return true, nil
	}
	return false, err
}
