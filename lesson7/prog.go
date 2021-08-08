package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

func inquery(q string, args ...interface{}) (string, []interface{}) {
	n := 0
	a := []interface{}{}
	re := regexp.MustCompile(`\?`)
	str := re.ReplaceAllStringFunc(q, func(string2 string) string {
		if reflect.TypeOf(args[n]).Kind() == reflect.Slice {
			r := []string{}
			v := reflect.ValueOf(args[n])
			fmt.Println(v)
			for i := 0; i < v.Len(); i++ {
				a = append(a, v.Index(i).Interface())
				r = append(r, "?")
			}
			n++
			return strings.Join(r, ",")
		} else {
			a = append(a, args[n])
			n++
			return string2
		}
	})
	return str, a
}
func main() {
	fmt.Println(inquery("SELECT * FROM table WHERE deleted = ? AND id IN(?) AND count < ?",
		false, []int{1, 6, 234}, 555))

}
