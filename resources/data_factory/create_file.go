package data_factory

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"test_toolkit/config"
	"test_toolkit/models"
	"test_toolkit/pkg/convert"
	"test_toolkit/pkg/utils"
	"time"

	"github.com/op/go-logging"
)

var logger = logging.MustGetLogger("test")

// FileMeta define the local file informations
type FileMeta struct {
	FileName     string
	FileFullPath string
	FileMd5sum   string
	FileSize     int64
}

// ParseFileCreateInput ...
func ParseFileCreateInput(input models.FileCreateInput) models.FileCreateConfigMap {
	logger.Info("> Parse File Create Input args ...")
	cfgMap := models.FileCreateConfigMap{CreateInput: input}
	// Parse FileInputs to conf.fileCreateConfig
	cfgMap.CreateConfig = make([]models.FileCreateConfig, len(input.FileArgs))
	for i, v := range input.FileArgs {
		fArr := strings.Split(v, ":")
		// fmt.Println(fArr)
		cfgMap.CreateConfig[i].Type = fArr[0]
		cfgMap.CreateConfig[i].Num, _ = strconv.Atoi(fArr[1])
		nArr := strings.Split(fArr[2], "-")
		cfgMap.CreateConfig[i].SizeMin = convert.String2Byte(nArr[0])
		if len(nArr) > 1 {
			cfgMap.CreateConfig[i].SizeMax = convert.String2Byte(nArr[1])
		} else {
			cfgMap.CreateConfig[i].SizeMax = cfgMap.CreateConfig[i].SizeMin
		}
		cfgMap.CreateConfig[i].NamePrefix = fmt.Sprintf("test_%s", config.TimeStr)
		cfgMap.CreateConfig[i].ParentDir = path.Join(input.LocalDataDir, config.TimeStr)
	}
	logger.Debugf("Create File Input:%v", utils.Prettify(input))
	return cfgMap
}

// CreateFiles ...
func CreateFiles(input models.FileCreateInput) []FileMeta {
	cfgMap := ParseFileCreateInput(input)
	logger.Info("> Prepare data files ...")
	var fileMetaArr []FileMeta
	var randomSize int64
	var mode string = "a+"
	var fileMd5 string

	for _, fileConf := range cfgMap.CreateConfig {
		fileNamePrefix := fileConf.NamePrefix
		patternPrefix := fileConf.NamePrefix
		emptyIdx := fileConf.Num * input.EmptyPercent / 100
		randomIdx := fileConf.Num * input.RandomPercent / 100
		fileDir := fileConf.ParentDir
		_, err := os.Stat(fileDir)
		if os.IsNotExist(err) {
			err := os.MkdirAll(fileDir, os.ModePerm)
			if err != nil {
				logger.Panicf("mkdir failed![%v]", err)
			}
		}

		// Get the exist files list
		existFileInfoList, err := ioutil.ReadDir(fileDir)
		if err != nil {
			logger.Panicf("List local files fail: %s", err)
		}

		if input.RenameFile {
			timeStr := time.Now().Format("20060102150405")
			fileNamePrefix += "_" + timeStr
			patternPrefix += "_\\d{14}"
		}

		for i := 0; i < fileConf.Num; i++ {
			newFile := FileMeta{}
			fileName := fmt.Sprintf("%s_%d.%s", fileNamePrefix, i, fileConf.Type)
			filePath := path.Join(fileDir, fileName)
			// os.Rename the exist file with diff timeStr
			if input.RenameFile {
				pattern := fmt.Sprintf("%s_%d.%s", patternPrefix, i, fileConf.Type)
				for _, existFile := range existFileInfoList {
					existFileName := existFile.Name()
					matched, _ := regexp.MatchString(pattern, existFileName)
					if matched {
						existFilePath := path.Join(fileDir, existFileName)
						logger.Infof("os.Rename: %s -> %s", existFilePath, filePath)
						os.Rename(existFilePath, filePath)
						break
					}
				}
			}

			if i < emptyIdx {
				randomSize = 0
				mode = "w+"
			} else {
				randomSize = utils.GetRandomInt64(fileConf.SizeMin, fileConf.SizeMax)
			}

			fExist, _ := utils.PathExists(filePath)
			if (i < randomIdx) && fExist {
				fileMd5 = utils.GetFileMd5sumWithPath(filePath)
			} else {
				fileMd5 = utils.CreateFile(filePath, randomSize, 128, mode)
			}

			newFile.FileName = fileName
			newFile.FileFullPath = filePath
			newFile.FileSize = randomSize
			newFile.FileMd5sum = fileMd5
			logger.Infof("Local File(md5:%s):%s", fileMd5, filePath)
			fileMetaArr = append(fileMetaArr, newFile)
		}
	}
	return fileMetaArr
}
