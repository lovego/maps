package maps

import (
	"fmt"
	"log"
	"reflect"
	"sort"
	"strings"
)

// Println print a map with keys in order.
func Println(m interface{}) {
	mapV := reflect.ValueOf(m)
	keys := mapV.MapKeys()
	sort.Slice(keys, func(i, j int) bool {
		switch keys[i].Kind() {
		case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
			return keys[i].Int() < keys[j].Int()
		case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
			return keys[i].Uint() < keys[j].Uint()
		case reflect.String:
			return keys[i].String() < keys[j].String()
		default:
			log.Panicf("unsupported key type: %v", keys[i].Type())
		}
		return false
	})
	slice := []string{}
	for _, key := range keys {
		slice = append(slice, fmt.Sprintf("%v:%v", key, mapV.MapIndex(key)))
	}
	fmt.Printf("map[%s]\n", strings.Join(slice, " "))
}
