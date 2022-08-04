package util

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"weixin/common/consts"
)

// 内部错误码和外部错误码的分隔，内部错误码使用 10000000 ~ MAX_INT，外部错误码使用 0 ~ 9999999
const minInternalErr = 10000000

var debugFlag = "[DEBUG]"

// DEBUG 调试
func DEBUG(i ...interface{}) {
	now := time.Now().Format("2006-01-02 15:04:05")
	_, file, line, _ := runtime.Caller(1)
	fmt.Printf("\n%s %s [BEGIN] @%s:%d \n", now, debugFlag, file, line)
	fmt.Printf("%s %s [DATA]:%+v ", now, debugFlag, i)
	fmt.Printf("\n%s %s [END]   @%s:%d \n", now, debugFlag, file, line)
}

func IsNil(a interface{}) bool {
	defer func() { recover() }()
	return a == nil || reflect.ValueOf(a).IsNil()
}

// FormatMaxPage 最大页
func FormatMaxPage(total, limit int64) int64 {
	return int64(math.Ceil(float64(total) / float64(limit)))
}

// 格式化分页参数
func FormatLimitAndOffset(limitReq int64, offsetReq int64, maxLimit ...int64) (limit int64, offset int64) {
	offset = consts.DefaultOffset
	limit = consts.DefaultLimit
	if limitReq > 0 {
		limit = limitReq
	}

	if offsetReq > 0 {
		offset = offsetReq
	}

	if limit > consts.DefaultMaxLimit {
		limit = consts.DefaultMaxLimit

		if len(maxLimit) == 1 && maxLimit[0] > consts.DefaultMaxLimit {
			limit = maxLimit[0]
		}
	}

	return
}

// FormatLimitAndOffset2Int 格式化分页参数 int
func FormatLimitAndOffset2Int(limitReq int64, offsetReq int64, maxLimit ...int64) (limit int, offset int) {
	l, o := FormatLimitAndOffset(limitReq, offsetReq, maxLimit...)
	return int(l), int(o)
}

// 格式化时间戳
func FormatTimestamp(timestamp int64, returnFormatType int) string {
	if timestamp < 1 {
		return ""
	}

	date := ""
	tm := time.Unix(timestamp, 0)
	switch returnFormatType {
	case consts.FormatTimestampType1:
		date = tm.Format("2006-01-02 15:04:05")

	case consts.FormatTimestampType2:
		date = tm.Format("2006-01-02 15:04")

	case consts.FormatTimestampType3:
		date = tm.Format("2006-01-02")

	case consts.FormatTimestampType6:
		date = tm.Format("2006-01")

	default:
		date = tm.Format("2006-01-02 15:04:05 PM")

	}

	return date
}

// 格式化日期
func FormatDate(date string, dateType int) int64 {
	if date == "" {
		return 0
	}

	var timestamp int64
	switch dateType {
	case consts.FormatTimestampType1:
		loc, _ := time.LoadLocation("Local")
		tm, _ := time.ParseInLocation("2006-01-02 15:04:05", date, loc)
		timestamp = tm.Unix()

	case consts.FormatTimestampType2:
		loc, _ := time.LoadLocation("Local")
		tm, _ := time.ParseInLocation("2006-01-02 15:04", date, loc)
		timestamp = tm.Unix()

	case consts.FormatTimestampType3:
		loc, _ := time.LoadLocation("Local")
		tm, _ := time.ParseInLocation("2006-01-02", date, loc)
		timestamp = tm.Unix()

	case consts.FormatTimestampType4:
		loc, _ := time.LoadLocation("Local")
		tm, _ := time.ParseInLocation("01-02-2006", date, loc)
		timestamp = tm.Unix()

	case consts.FormatTimestampType5:
		loc, _ := time.LoadLocation("Local")
		tm, _ := time.ParseInLocation("200601", date, loc)
		timestamp = tm.Unix()
		break
	case consts.FormatTimestampType6:
		loc, _ := time.LoadLocation("Local")
		tm, _ := time.ParseInLocation("2006-01", date, loc)
		timestamp = tm.Unix()
		break
	default:
		timestamp = 0

	}

	return timestamp
}

//GetYearDiffer 根据出生日期，计算年龄
func GetYearDiffer(start_time, end_time string, dateType int) int64 {
	if start_time == "" || end_time == "" {
		return 0
	}
	var year int64
	switch dateType {
	case consts.FormatTimestampType1:
		b, _ := time.Parse("2006-01-02 15:04:05", start_time)
		a, _ := time.Parse("2006-01-02 15:04:05", end_time)
		d := a.Sub(b)
		year = int64(d.Hours() / 24 / 365)
	case consts.FormatTimestampType2:
		b, _ := time.Parse("2006-01-02 15:04", start_time)
		a, _ := time.Parse("2006-01-02 15:04", end_time)
		d := a.Sub(b)
		year = int64(d.Hours() / 24 / 365)

	case consts.FormatTimestampType3:
		b, _ := time.Parse("2006-01-02", start_time)
		a, _ := time.Parse("2006-01-02", end_time)
		d := a.Sub(b)
		year = int64(d.Hours() / 24 / 365)

	case consts.FormatTimestampType4:
		b, _ := time.Parse("2006-01", start_time)
		a, _ := time.Parse("2006-01", end_time)
		d := a.Sub(b)
		year = int64(d.Hours() / 24 / 365)
	default:
		year = 0

	}

	return year
}

type GenOuterErrInterface interface {
	GetErrno() int64
	GetErrmsg() string
}

// 根据年限计算起始时间
func GetStartYear(year int, returnFormatType int) string {
	if year < 0 {
		return ""
	}

	t := time.Now()
	tm := t.AddDate(-int(year), 0, 0)
	date := ""
	switch returnFormatType {
	case consts.FormatTimestampType1:
		date = tm.Format("2006-01-02 15:04:05")

	case consts.FormatTimestampType2:
		date = tm.Format("2006-01-02 15:04")

	case consts.FormatTimestampType3:
		date = tm.Format("2006-01-02")

	case consts.FormatTimestampType6:
		date = tm.Format("2006-01")

	default:
		date = tm.Format("2006-01-02 15:04:05 PM")

	}

	return date
}

// JoinInt64ToString splic 链接
func JoinInt64ToString(slice []int64, sep string) string {
	if slice == nil {
		return ""
	}

	sliceStr := make([]string, 0, len(slice))
	for _, i := range slice {
		sliceStr = append(sliceStr, strconv.FormatInt(i, 10))
	}

	return strings.Join(sliceStr, ",")
}

func GetResourcePath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)

	rst := filepath.Dir(path)

	resourcePath := fmt.Sprintf("%s/resource", rst)
	if IsDir(resourcePath) {
		return resourcePath
	}

	return fmt.Sprintf("%s/../resource", rst)
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()

}

// JoinInt64 []int64 -> string
func JoinInt64(i []int64, sep string) string {
	if len(i) == 0 {
		return ""
	}

	s := ""

	for _, v := range i {
		s = fmt.Sprintf("%s%s%s", s, strconv.FormatInt(v, 10), sep)
	}

	return strings.TrimRight(s, sep)
}

// 获取中文长度
func GetChineseLen(s string) int {
	return len([]rune(s))
}

// MaybeChinese 纯中文
func MaybeChinese(keywords string) bool {
	if GetChineseLen(keywords) > 30 {
		return false
	}
	reg := regexp.MustCompile("^[\u4e00-\u9fa5]+$")
	return reg.MatchString(keywords)
}

func MayEmail(keywords string) bool {
	reg := regexp.MustCompile(`\w[-\w.+]*@([A-Za-z0-9][-A-Za-z0-9]+\.)+[A-Za-z]{2,14}`)
	return reg.MatchString(keywords)
}

//纯英文（英文名）
func MayEnglish(keywords string) bool {
	reg := regexp.MustCompile(`^[A-Za-z]+$`)
	return reg.MatchString(keywords)
}

// 手机号规则验证
func MaybeMobile(keywords string) bool {
	reg := regexp.MustCompile(`(?:(?:\+|00)86)?1[3-9]\d{9}?$`)
	return reg.MatchString(keywords)
}

// 英文+数字
func MayEnglishNum(keywords string) bool {
	reg := regexp.MustCompile(`^[A-Za-z0-9]+$`)
	return reg.MatchString(keywords)
}

//CompareJson 比较两个json是否相等
func CompareJson(str1, str2 string) (compareRes bool, err error) {
	var (
		json1 map[string]interface{}
		json2 map[string]interface{}
	)

	if err = LoadJson(str1, &json1); err != nil {
		return false, err
	}
	if err = LoadJson(str2, &json2); err != nil {
		return false, err
	}
	compareRes = reflect.DeepEqual(json1, json2)
	return compareRes, nil
}

func LoadJson(str string, dist interface{}) (err error) {
	err = json.Unmarshal([]byte(str), dist)
	return err
}

//checkTime
func CheckTime(days string) error {
	year := time.Now().Year()
	var daysNew, _ = strconv.Atoi(days)
	if days != "" && (daysNew <= 0 || daysNew >= (year-1900)*366) {
		return errors.New("时间格式不合法")
	}

	return nil
}

// int64切片转json
func Slices2Json(slicesData []int64) (string, error) {
	if len(slicesData) == 0 {
		return "[]", nil
	}

	byto, err := json.Marshal(slicesData)
	if err != nil {
		return "[]", err
	}
	s := string(byto)
	return s, nil
}

// GetMonthStartEnd 获取某月第一天和最后一天的时间戳
// monthOffset 是偏移量，如果是获取本月则传0，上个月为-1，以此类推
func GetMonthStartEnd(monthOffset int64) (int64, int64) {
	now := time.Now()
	lastMonthFirstDay := now.AddDate(0, int(monthOffset), -now.Day()+1)
	lastMonthStart := time.Date(lastMonthFirstDay.Year(), lastMonthFirstDay.Month(), lastMonthFirstDay.Day(), 0, 0, 0, 0, now.Location()).Unix()
	lastMonthEndDay := lastMonthFirstDay.AddDate(0, 1, -1)
	lastMonthEnd := time.Date(lastMonthEndDay.Year(), lastMonthEndDay.Month(), lastMonthEndDay.Day(), 23, 59, 59, 0, now.Location()).Unix()
	return lastMonthStart, lastMonthEnd
}

// GetGivenYearMonthTimeStamp 获取指定年月的第一天和最后一天的时间戳
func GetGivenYearMonthTimeStamp(str string) (int64, int64, error) {
	_, err := time.Parse("2006-01", str)
	if err != nil {
		return 0, 0, errors.New("日期格式错误")
	}
	tt, _ := time.ParseInLocation("2006-01", str, time.Local)
	lastMonthFirstDay := tt.AddDate(0, 0, -tt.Day()+1)
	lastMonthStart := time.Date(lastMonthFirstDay.Year(), lastMonthFirstDay.Month(), lastMonthFirstDay.Day(), 0, 0, 0, 0, tt.Location()).Unix()
	lastMonthEndDay := lastMonthFirstDay.AddDate(0, 1, -1)
	lastMonthEnd := time.Date(lastMonthEndDay.Year(), lastMonthEndDay.Month(), lastMonthEndDay.Day(), 23, 59, 59, 0, tt.Location()).Unix()

	return lastMonthStart, lastMonthEnd, nil
}

func StrToTime(format, strTime string) (int64, error) {
	loc, err := time.LoadLocation("Local") //获取时区
	if err != nil {
		return 0, err
	}
	t, err := time.ParseInLocation(format, strTime, loc)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

//IsInSlice 判断一个值，是否在一个切片中
func IsInSlice(val int64, targetSlice []int64) bool {
	if len(targetSlice) == 0 {
		return false
	}
	for _, v := range targetSlice {
		if v == val {
			return true
		}
	}

	return false
}

// FormatFloat 保留n位小数，去尾数0
func FormatFloat(num float64, decimal int) string {
	// 默认乘1
	d := float64(1)
	if decimal > 0 {
		// 10的N次方
		d = math.Pow10(decimal)
	}
	// math.trunc作用就是返回浮点数的整数部分
	// 再除回去，小数点后无效的0也就不存在了
	return strconv.FormatFloat(math.Trunc(num*d)/d, 'f', -1, 64)
}

func Md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制

	return md5str
}
