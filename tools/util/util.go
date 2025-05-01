package util

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func RandNum(min, max int64) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	code := r.Int63n(max-min) + min
	return Int2Str(int(code))
}

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func IntToBool(i int) bool {
	if i == 0 {
		return false
	}
	return true
}

func StringToInt(s string) int {
	if s == "" {
		return 0
	}
	value, _ := strconv.Atoi(s)
	return value
}

func StringToInt32(s string) int32 {
	if s == "" {
		return 0
	}
	value, _ := strconv.Atoi(s)
	return int32(value)
}

func Float64ToStr(float float64) string {
	return strconv.FormatFloat(float, 'f', -1, 64)
}

func StringToInt64(s string) int64 {
	value, _ := strconv.ParseInt(s, 10, 64)
	return value
}

func StringToFloat64(s string) float64 {
	value, err := strconv.ParseFloat(s, 10)
	if err != nil {
		return 0.0
	}
	return value
}

//字符串转int,失败返回默认值,没有默认值返回）
func Int2Str(i int) string {
	return strconv.Itoa(i)
}

func Str2Int(s string) int {
	if s == "" || len(s) == 0 {
		return 0
	}
	res, err := strconv.Atoi(s)
	if err != nil {
		res = 0
	}
	return res
}

func Str2Int64(s string) int64 {
	if s == "" || len(s) == 0 {
		return 0
	}
	res, err := strconv.Atoi(s)
	if err != nil {
		res = 0
	}
	return int64(res)
}

//查看字符串是否包含子字符串
func HasString(param ...string) bool {
	if len(param) < 2 {
		return false
	}

	str := param[0]

	for i := 1; i < len(param); i++ {
		if strings.Contains(str, param[i]) {
			return true
		}
	}
	return false
}

func SliceIntMap(str, sub string, mapData map[string]int) []int {
	params := strings.Split(str, sub)
	data := make([]int, 0)

	for _, param := range params {
		if param != "" {
			num, ok := mapData[param]
			if ok {
				data = append(data, num)
			}
		}

	}
	return data
}

func SliceString2Int(str, sub string) []int {
	params := strings.Split(str, sub)
	data := make([]int, 0)

	for _, param := range params {
		if param != "" {
			num, err := strconv.Atoi(param)
			if err == nil {
				data = append(data, num)
			}
		}

	}
	return data
}

func SliceStringMap(str, sub string, mapData map[string]string) []string {
	data := make([]string, 0)
	if mapData == nil {
		return data
	}

	params := strings.Split(str, sub)
	for _, param := range params {
		if param != "" {
			value, ok := mapData[param]
			if ok {
				data = append(data, value)
			}
		}

	}
	return data
}

func SliceString(str, sub string) []string {
	params := strings.Split(str, sub)
	data := make([]string, 0)

	for _, param := range params {
		if param != "" {
			data = append(data, param)
		}
	}
	return data
}

func SliceInt(str, sub string) []int {
	params := strings.Split(str, sub)
	data := make([]int, 0)
	for _, param := range params {
		if param != "" {
			data = append(data, Str2Int(param))
		}
	}
	return data
}

func SliceToString(data []string, sub string) string {
	if len(data) == 0 {
		return ""
	}
	str := ""
	for k, v := range data {
		if k != len(data)-1 {
			str += v + sub
		} else {
			str += v
		}

	}
	return str
}

//校验时间戳单位是否为秒级别：true.秒级别时间戳；false.非秒级别时间戳
func TimestampUnit(times int) bool {
	var bool = true
	if times >= 10000000000 {
		bool = false
	}
	return bool
}

//切片 变成字符串 默认,号连接
func SliceIntToString(s []string, old, new string) string {
	if new == "" {
		new = ","
	}
	if old == "" {
		old = " "
	}
	return strings.Replace(strings.Trim(fmt.Sprint(s), "[]"), old, new, -1)
}

//切片变执行sqls时的in条件
func SliceStringToString(s []string, old, new string) string {
	if new == "" {
		new = "','"
	}
	if old == "" {
		old = " "
	}
	return "'" + strings.Replace(strings.Trim(fmt.Sprint(s), "[]"), old, new, -1) + "'"
}

//整数转浮点型百分比
func IntAccuracy(i, j int) float64 {
	if i == 0 || j == 0 {
		return 0
	}
	return float64(i) / float64(j)
}

func FloatString(num float64, base int) string {
	if base > 4 {
		base = 4
	}
	if base == 0 {
		base = 1
	}

	str := ""
	switch base {
	case 0:
		{
			str = fmt.Sprintf("%0.0f", num)
		}
	case 1:
		{
			str = fmt.Sprintf("%0.1f", num)
		}
	case 2:
		{
			str = fmt.Sprintf("%0.2f", num)
		}
	case 3:
		{
			str = fmt.Sprintf("%0.3f", num)
		}
	case 4:
		{
			str = fmt.Sprintf("%0.4f", num)
		}
	}
	return str
}

//整数转浮点型百分比字符串 最大支持小数点后四位
func IntAccuracyString(i, j, base int) string {
	if base > 4 {
		base = 4
	}
	if i == 0 || j == 0 {
		return "0"
	}

	num := float64(i) / float64(j)
	str := ""
	switch base {
	case 0:
		{
			str = fmt.Sprintf("%0.0f", num)
		}
	case 1:
		{
			str = fmt.Sprintf("%0.1f", num)
		}
	case 2:
		{
			str = fmt.Sprintf("%0.2f", num)
		}
	case 3:
		{
			str = fmt.Sprintf("%0.3f", num)
		}
	case 4:
		{
			str = fmt.Sprintf("%0.4f", num)
		}
	}
	return str
}

//从计数的map中取出排在前面的count个
func GetMapSortN(mapData map[string]int, count int) []string {
	if mapData == nil {
		return nil
	}

	mapTemp := make(map[int][]string)
	sortTemp := make([]int, 0)
	for k, v := range mapData {
		temp, ok := mapTemp[v]
		if ok {
			temp = append(temp, k)
			mapTemp[v] = temp
		} else {
			sortTemp = append(sortTemp, v)
			mapTemp[v] = []string{k}
		}
	}
	sort.Ints(sortTemp)
	num := count
	data := make([]string, 0)
	for i := len(sortTemp) - 1; i >= 0; i-- {
		sli := mapTemp[sortTemp[i]]
		for _, v := range sli {
			num--
			data = append(data, v)
			if num == 0 {
				return data
			}
		}
	}
	return data
}

//去除字符串中最后一个字符
func TrimLastChar(s string) string {
	r, size := utf8.DecodeLastRuneInString(s)
	if r == utf8.RuneError && (size == 0 || size == 1) {
		size = 0
	}
	return s[:len(s)-size]
}

//去除字符串的空格  换行 制表符
func StringFmt(s string) string {
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "\t", "")
	return s
}

func MapFmt(data map[string]interface{}) {
	if data == nil {
		return
	}

	for k, v := range data {
		value, ok := v.(string)
		if ok {
			if value == "" {
				delete(data, k)
			}
		}
	}
}

func CreateUUID() string {
	return uuid.NewV4().String()
}

func VerifyEmail(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

//并发安全对象
var snowFlakeID *snowflake.Node

func CreateSerialNumber() string {
	if snowFlakeID == nil {
		snowFlakeID, _ = snowflake.NewNode(1)
	}

	id := snowFlakeID.Generate()
	return strconv.FormatInt(int64(id), 10)
}