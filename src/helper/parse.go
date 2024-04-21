package helper

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Print to Json
func JsonString(o interface{}) string {
	b, e := json.Marshal(o)
	if e != nil {
		return string([]byte("{}"))
	}

	return string(b)
}

func PrintJson(o interface{}) {
	fmt.Println(JsonString(o))
}

func StructToMap(from interface{}, to interface{}) error {
	inrec, _ := json.Marshal(from)
	if e := json.Unmarshal(inrec, to); e != nil {
		return e
	}
	return nil
}

func MapToStruct(from interface{}, to interface{}) error {
	err := json.Unmarshal([]byte(JsonString(from)), to)
	if err != nil {
		return err
	}
	return nil
}

func ToInt(o interface{}) int {
	switch dataType := Value(o); dataType.Kind() {
	case reflect.String:
		s, ok := o.(string)
		if !ok {
			return 0
		}

		i, e := strconv.Atoi(s)
		if e != nil {
			return 0
		}
		return i
	case reflect.Float64:
		f := o.(float64)
		return int(f)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return int(dataType.Int())
	}
	return 0
}

func ToString(o interface{}) string {
	switch x := o.(type) {
	case primitive.ObjectID:
		return x.Hex()
	case float64:
		newvalue := decimal.NewFromFloat(o.(float64))
		return newvalue.String()
	}
	return fmt.Sprintf("%v", o)
}

func ToFloat64WithoutRounding(o interface{}) float64 {
	//fmt.Printf("\ndec: %v\n", decimalPoint)
	if IsPointer(o) {
		return float64(0)
	}

	var f float64
	var e error

	t := strings.ToLower(TypeName(o))
	v := Value(o)

	if t != "interface{}" && strings.HasPrefix(t, "int") {
		f = float64(v.Int())
	} else if strings.HasPrefix(t, "uint") {
		f = float64(v.Uint())
	} else if strings.HasPrefix(t, "float") {
		f = float64(v.Float())
	} else {
		f, e = strconv.ParseFloat(v.String(), 64)
		if e != nil {
			return 0
		}
	}

	return f
}

func ToFloat64(o interface{}, decimalPoint int, rounding string) float64 {
	//fmt.Printf("\ndec: %v\n", decimalPoint)
	if IsPointer(o) {
		return float64(0)
	}

	var f float64
	var e error

	t := strings.ToLower(TypeName(o))
	v := Value(o)

	if t != "interface{}" && strings.HasPrefix(t, "int") {
		f = float64(v.Int())
	} else if strings.HasPrefix(t, "uint") {
		f = float64(v.Uint())
	} else if strings.HasPrefix(t, "float") {
		f = float64(v.Float())
	} else {
		f, e = strconv.ParseFloat(v.String(), 64)
		if e != nil {
			return 0
		}
	}

	//fmt.Printf("\ndec: %v\n", decimalPoint)
	switch rounding {
	case "RoundAuto":
		return RoundingAuto64(f, decimalPoint)
	case "RoundDown":
		return RoundingDown64(f, decimalPoint)
	case "RoundUp":
		return RoundingUp64(f, decimalPoint)
	}

	if math.IsNaN(f) || math.IsInf(f, 0) {
		f = 0
	}

	return f
}

func RoundingAuto64(f float64, decimalPoint int) (retValue float64) {

	tempPow := math.Pow(10, float64(decimalPoint))
	f = f * tempPow

	if f < 0 {
		f = math.Ceil(f - 0.5)
	} else {
		f = math.Floor(f + 0.5)
	}

	retValue = f / tempPow
	return
}

func RoundingDown64(f float64, decimalPoint int) (retValue float64) {
	tempPow := math.Pow(10, float64(decimalPoint))
	f = f * tempPow
	f = math.Floor(f)
	retValue = f / tempPow
	return
}

func RoundingUp64(f float64, decimalPoint int) (retValue float64) {
	tempPow := math.Pow(10, float64(decimalPoint))
	f = f * tempPow
	f = math.Ceil(f)
	retValue = f / tempPow
	return
}

func IsPointer(o interface{}) bool {
	v := reflect.ValueOf(o)
	return v.Kind() == reflect.Ptr
}

func Value(o interface{}) reflect.Value {
	return reflect.ValueOf(o)
}

func TypeName(o interface{}) string {
	typeName := fmt.Sprintf("%T", o)

	switch o.(type) {
	case string:
		typeName = "string"
	case int:
		typeName = "int"
	case int32:
		typeName = "int32"
	case float64:
		typeName = "float64"
	case bool:
		typeName = "bool"
	}

	return typeName
}

func ToTime(o interface{}) time.Time {
	switch x := o.(type) {
	case primitive.DateTime:
		return x.Time()
	case time.Time:
		return x
	case string:
		StoT, _ := time.Parse("2006-01-02", x)
		return StoT
	}
	return time.Time{}
}

// escaping value to avoid sql injection
func SqlEscape(s string) string {
	return strings.Replace(s, "'", "''", -1)
}

func SplitLines(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

// cast slice interface to slice integer
func ToArrayInt(i []interface{}) []int {
	ints := make([]int, len(i))
	for x := range i {
		ints[x] = ToInt(i[x])
	}
	return ints
}

func Contains(d []string, s string) bool {
	isContain := false
	if len(d) > 0 {
		for _, x := range d {
			if x == s {
				isContain = true
				break
			}
		}
	}

	return isContain
}

func Div(f1, f2 float64) float64 {
	if f2 == 0 {
		return 0
	}

	return f1 / f2
}

func RemoveDuplicateString(strSlice []string) []string {
	result := []string{}
	keys := make(map[string]bool)

	for _, each := range strSlice {
		if _, value := keys[each]; !value {
			keys[each] = true
			result = append(result, each)
		}
	}

	return result
}

// FloatToStringMDP will format the decimal as the minimum decimal precision if does not have decimal,
// or return the original (or up to maximal decimal precision if maxPrec > 0) if decimal higher than the minimum precision given. e.g.
//
// with minPrec 2 & maxPrec 0, float 10 will be formated as 10.00.
//
// with minPrec 2 & maxPrec 0, float 10.12345678 will be formated as 10.12345678.
//
// with minPrec 2 & maxPrec 4,  float 10.12345678 will be formated as 10.1235.
//
// with minPrec 0 & maxPrec 4,  float 10.12345678 will be formated as 10.1235.
//
// minPrec <= 0 and maxPrec <= 0 will return original decimal precision.
func FloatToStringMDP(f float64, minPrec int, maxPrec int) string {
	if maxPrec > 0 {
		f, _ = strconv.ParseFloat(strconv.FormatFloat(f, 'f', maxPrec, 64), 64)
	}
	minPrecMultiplier := math.Pow10(minPrec)
	// Avoid golang's imprecision operation. e.g. 10.20 * 100 = 1019.9999999999999. should be 1020.
	fMultipled, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", f*minPrecMultiplier), 64)
	if fMultipled == math.Floor(fMultipled) && minPrec > 0 {
		return strconv.FormatFloat(f, 'f', minPrec, 64)
	}
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func InterfaceToFloat64(intf interface{}) float64 {
	res := 0.0
	if intf != nil {
		res = intf.(float64)
	}
	return res
}

func TrimLeftRightSpace(str string) string {
	str = strings.TrimLeft(str, " ")
	str = strings.TrimRight(str, " ")
	return str
}

func TrimLeftRightChar(str string, substring string) string {
	str = strings.TrimLeft(str, substring)
	str = strings.TrimRight(str, substring)
	return str
}

func ArrayInterfacesToString(data []interface{}) string {
	result := ""
	for _, each := range data {
		if each != nil && reflect.TypeOf(each).String() == "string" {
			if result == "" {
				result = result + each.(string)
			} else {
				result = result + "," + each.(string)
			}
		}
	}
	return result
}

func ContainsObjectId(d []primitive.ObjectID, s primitive.ObjectID) bool {
	isContain := false
	if len(d) > 0 {
		for _, x := range d {
			if x == s {
				isContain = true
				break
			}
		}
	}

	return isContain
}

func Ordinal(x int) string {
	suffix := "th"
	switch x % 10 {
	case 1:
		if x%100 != 11 {
			suffix = "st"
		}
	case 2:
		if x%100 != 12 {
			suffix = "nd"
		}
	case 3:
		if x%100 != 13 {
			suffix = "rd"
		}
	}
	return strconv.Itoa(x) + suffix
}

func RemoveStringItem(items []string, item string) []string {
	newitems := []string{}

	for _, i := range items {
		if i != item {
			newitems = append(newitems, i)
		}
	}

	return newitems
}

func StringToBool(data string) bool {
	res := false

	res, err := strconv.ParseBool(data)
	if err != nil {
		return res
	}

	return res
}

func FieldFilterBoardpack(fieldfilter string) string {
	fieldnamefilter := "clientgroup"
	if fieldfilter == "AccountID" {
		fieldnamefilter = "clientnumber"
	}
	return fieldnamefilter
}
