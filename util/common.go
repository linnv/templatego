package util

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/linnv/logx"
)

var WeekDaysCN = [...]string{
	"星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六",
}

var (
	allchars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	// allcharsLen = len(allchars)
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func StringArrayToStr(a []string) string {
	var res string = "["
	res += strings.Join(a, ",")
	res += "]"
	return res
}

func IntArrayToStr(a []int) string {
	var res string = "["
	for i, item := range a {
		var str = strconv.FormatInt(int64(item), 10)
		if i == 0 {
			res = res + str
		} else {
			res = res + "," + str
		}
	}
	return res + "]"
}

func InIntArray(i int, ints []int) bool {
	for _, v := range ints {
		if i == v {
			return true
		}
	}
	return false
}

func IntArraySubtract(a, b []int) (c []int) {
	c = []int{}

	for _, _a := range a {
		if !InIntArray(_a, b) {
			c = append(c, _a)
		}
	}

	return
}

func InStringArray(s string, strs []string) bool {
	for _, v := range strs {
		if s == v {
			return true
		}
	}
	return false
}

// 交集
func Intersection(a, b []int) []int {
	var c = make([]int, 0)
	for _, v := range a {
		if InIntArray(v, b) {
			c = append(c, v)
		}
	}

	return c
}

//数组去重
func UniqueIntArray(a []int) []int {
	al := len(a)
	if al == 0 {
		return a
	}

	ret := make([]int, al)
	index := 0

loopa:
	for i := 0; i < al; i++ {
		for j := 0; j < index; j++ {
			if a[i] == ret[j] {
				continue loopa
			}
		}
		ret[index] = a[i]
		index++
	}

	return ret[:index]
}

func GetDuplicateIntArray(a []int) []int {
	al := len(a)
	if al == 0 {
		return nil
	}

	ret := make([]int, al)
	dupRet := make([]int, 0, 1)
	index := 0

loopa:
	for i := 0; i < al; i++ {
		for j := 0; j < index; j++ {
			if a[i] == ret[j] {
				dupRet = append(dupRet, a[j])
				continue loopa
			}
		}
		ret[index] = a[i]
		index++
	}

	return dupRet
}

func GetMapIntKeys(m map[int]interface{}) []int {
	keys := make([]int, 0)
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

//(?i)\b((?:[a-z][\w-]+:(?:/{1,3}|[a-z0-9%])|www\d{0,3}[.]|[a-z0-9.\-]+[.][a-z]{2,4}/)(?:[^\s()<>]+|\(([^\s()<>]+|(\([^\s()<>]+\)))*\))+(?:\(([^\s()<>]+|(\([^\s()<>]+\)))*\)|[^\s`!()\[\]{};:\'".,<>?«»“”‘’]))
var urlPattern = regexp.MustCompile("^(http|https)://.*")

func IsValidURL(url string) bool {
	return urlPattern.MatchString(url)
}

var dateLayout = "2006-01-02"

func ParseDate(str string) int64 {
	t, err := time.ParseInLocation(dateLayout, str, time.Local)
	if err != nil {
		return 0
	}
	return t.Unix()
}

//将idsStr转换为[]int格式  2,3,4 -> []int{2,3,4}
func StrToIntArray(str string) []int {
	var strs = strings.Split(strings.TrimSpace(str), ",")
	var ids = make([]int, 0, len(strs))
	for _, idStr := range strs {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			continue
		}
		ids = append(ids, id)
	}
	return ids
}

//将strings转换为[]string格式  "a,b,c" -> []string{a,b,c}
func StrToStringArray(str string) []string {
	var strs = strings.Split(strings.TrimSpace(str), ",")
	var result = make([]string, 0, len(strs))
	for _, s := range strs {
		if len(s) > 0 {
			result = append(result, s)
		}
	}
	return result
}

func BytesEqual(a, b []byte) bool {
	al := len(a)
	if al != len(b) {
		return false
	}

	for i := 0; i < al; i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func SaveAsFile(b []byte, filepath string) (int, error) {
	f, err := os.Create(filepath)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	n, err := f.Write(b)
	if err != nil {
		os.Remove(filepath)
	}
	return n, err
}

func RandString(strlen int) string {
	return RandStr(strlen, allchars)
}

func RandStr(strlen int, chars string) string {
	b := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}

// func GetWidthAndHeight(size string) (width int, height int, ok bool) {
// 	var strs = strings.Split(size, constant.SIZE_JOIN_SYMBOL)
// 	if len(strs) < 2 {
// 		ok = false
// 		return
// 	}
// 	width, err := strconv.Atoi(strs[0])
// 	if err != nil {
// 		ok = false
// 		return
// 	}
//
// 	height, err = strconv.Atoi(strs[1])
// 	if err != nil {
// 		ok = false
// 		return
// 	}
// 	ok = true
// 	return
// }

func GetDatePath(sub string) string {
	var ps = strings.Split(time.Now().Format("2006,01"), ",")
	ps = append(ps, sub)
	return path.Join(ps...)
}

// func HashUserId(userId int) string {
// 	return Md5Encode(strconv.Itoa(userId))[:8]
// }

func CoerceInt(v interface{}) (int, error) {
	switch v := v.(type) {
	case string:
		i64, err := strconv.ParseInt(v, 10, 0)
		return int(i64), err
	case int, int16, int32, int64:
		return int(reflect.ValueOf(v).Int()), nil
	case uint, uint16, uint32, uint64:
		return int(reflect.ValueOf(v).Uint()), nil
	case float32:
		return int(v), nil
	case float64:
		return int(v), nil
	}
	return 0, errors.New("invalid value type")
}

func CoerceFloat(v interface{}) (float64, error) {
	switch v := v.(type) {
	case string:
		if len(v) > 0 {
			if iv, err := strconv.ParseFloat(v, 64); err == nil {
				return iv, nil
			}
		}
	case int, int16, int32, int64:
		return float64(reflect.ValueOf(v).Int()), nil
	case uint, uint16, uint32, uint64:
		return float64(reflect.ValueOf(v).Uint()), nil
	case float32:
		return float64(v), nil
	case float64:
		return float64(v), nil
	}
	return 0, errors.New("invalid value type")
}

func CoerceString(v interface{}) (string, error) {
	switch v := v.(type) {
	case string:
		return v, nil
	case int, int16, int32, int64, uint, uint16, uint32, uint64:
		return fmt.Sprintf("%d", v), nil
	case float32, float64:
		return fmt.Sprintf("%f", v), nil
	}
	return fmt.Sprintf("%s", v), nil
}

const DateFormatNoLine = "20060102"

func ParseDate2TimeEx(format, str string) (time.Time, error) {
	return time.ParseInLocation(format, str, time.Local)
}

func ParseDate2Time(str string) (time.Time, error) {
	return time.ParseInLocation(dateLayout, str, time.Local)
}

//e.g. 20170613->"星期二"
func GetWeekDayCN(date string) string {
	t, err := ParseDate2TimeEx(DateFormatNoLine, date)
	if err != nil {
		return "-"
	}
	return WeekDaysCN[t.Weekday()]
}

func MosaicStar(str string, beginKeep, endKeep int) string {
	stars := len(str) - beginKeep - endKeep
	if stars <= 0 {
		return str
	}

	return str[:beginKeep] + strings.Repeat("*", stars) + str[beginKeep+stars:]
}

func VerifyBankID(id string) bool {
	reg := `^[0-9]{%d}$`
	b, _ := regexp.MatchString(fmt.Sprintf(reg, len(id)), id)
	return b
}

func VerifyTaxID(id string) bool {
	reg := `^[\d\w]{%d}$`
	b, _ := regexp.MatchString(fmt.Sprintf(reg, len(id)), id)
	return b
}

var (
	mobilePattern = regexp.MustCompile(`^1[3578]\d{9}$`)
	// telephonePattern = regexp.MustCompile(`^(0[0-9]{2,3}\-)?([2-9][0-9]{6,7})+(\-[0-9]{1,4})?$`)

	CardIDPattern = regexp.MustCompile(`^(([京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼使领][A-Z](([0-9]{5}[DF])|([DF]([A-HJ-NP-Z0-9])[0-9]{4})))|([京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼使领][A-Z][A-HJ-NP-Z0-9]{4}[A-HJ-NP-Z0-9挂学警港澳使领]))$`)
)

func Number2Cn(number string) (ret string) {
	numberCn := map[byte]rune{
		'0': rune('零'), '1': rune('一'), '2': rune('二'), '3': rune('三'), '4': rune('四'), '5': rune('五'), '6': rune('六'), '7': rune('七'), '8': rune('八'), '9': rune('九'),
		// '1': rune('幺'),
	}
	runes := []rune(number)
	numberLen := len(runes)
	ll := make([]rune, 0, numberLen)
	for i := 0; i < numberLen; i++ {
		if o, ok := numberCn[number[i]]; ok {
			ll = append(ll, o)
		} else {
			ll = append(ll, runes[i])
		}
	}
	return string(ll)

}

func VerifyCardID(number string) (filterQuery string, ok bool) {
	return number, CardIDPattern.MatchString(number)
}

var NumberList = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

//ForceNumber implements ...
func ForceNumber(raw string) string {
	CnNumber := map[rune]byte{
		rune('零'): '0', rune('一'): '1', rune('二'): '2', rune('三'): '3', rune('四'): '4', rune('五'): '5', rune('六'): '6', rune('七'): '7', rune('八'): '8', rune('九'): '9', rune('幺'): '1',
		rune('〇'): '0', rune('壹'): '1', rune('贰'): '2', rune('叁'): '3', rune('肆'): '4', rune('伍'): '5', rune('陆'): '6', rune('柒'): '7', rune('捌'): '8', rune('玖'): '9', rune('貮'): '2', rune('两'): '2',
	}
	olds := []rune(raw)
	oldsLen := len(olds)
	news := make([]byte, 0, oldsLen)
	for i := 0; i < oldsLen; i++ {
		if olds[i] >= rune('0') && olds[i] <= rune('9') {
			news = append(news, byte(olds[i]))
			continue
		}
		if get, ok := CnNumber[olds[i]]; ok {
			news = append(news, get)
		}
	}
	return string(news)
}

func VerifyContactFilterFirst(number string) (filterQuery string, maybePhone, ok bool) {
	CnNumber := map[rune]byte{
		rune('零'): '0', rune('一'): '1', rune('二'): '2', rune('三'): '3', rune('四'): '4', rune('五'): '5', rune('六'): '6', rune('七'): '7', rune('八'): '8', rune('九'): '9', rune('幺'): '1',
		rune('〇'): '0', rune('壹'): '1', rune('贰'): '2', rune('叁'): '3', rune('肆'): '4', rune('伍'): '5', rune('陆'): '6', rune('柒'): '7', rune('捌'): '8', rune('玖'): '9', rune('貮'): '2', rune('两'): '2',
	}
	olds := []rune(number)
	oldsLen := len(olds)
	news := make([]byte, 0, oldsLen)
	for i := 0; i < oldsLen; i++ {
		if olds[i] >= rune('0') && olds[i] <= rune('9') {
			news = append(news, byte(olds[i]))
			continue
		}
		if get, ok := CnNumber[olds[i]]; ok {
			news = append(news, get)
		}
	}
	for i := 0; i < len(NumberList); i++ {
		if strings.Contains(number, NumberList[i]) {
			maybePhone = true
		}
	}

	filterQuery = string(news)
	ok = VerifyContact(filterQuery)
	logx.Warnf("filterQuery: %s maybe:%v,ok:%v \n", filterQuery, maybePhone, ok)
	return
}

func VerifyContact(number string) bool {
	return mobilePattern.MatchString(number)
	// r :=
	// if !r {
	// 	return telephonePattern.MatchString(number)
	// }
	// return true
}

func ChineseLength(str string) int {
	return len([]rune(str))
}

func IsChineseStr(str string) bool {
	for _, r := range str {
		if !unicode.Is(unicode.Scripts["Han"], r) {
			return false
		}
	}
	return true
}

func VerifyPersonalID(id string) bool {
	reg := `^[\d\w]{%d}$`
	b, _ := regexp.MatchString(fmt.Sprintf(reg, len(id)), id)
	return b
}

func StringArrayEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for _, _a := range a {
		finda := false
		for _, _b := range b {
			if _b == _a {
				finda = true
				break
			}
		}
		if !finda {
			return false
		}
	}

	return true
}

func IntArrayEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for _, _a := range a {
		finda := false
		for _, _b := range b {
			if _b == _a {
				finda = true
				break
			}
		}
		if !finda {
			return false
		}
	}

	return true
}

var (
	emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

func ValidateFormat(email string) bool {
	return emailRegexp.MatchString(email)
}

func PanicOnError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
