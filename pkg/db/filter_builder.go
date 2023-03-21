package db

import (
	"fmt"
	"reflect"
)

type Filter struct {
	Key      string      `json:"key"`
	Value    interface{} `json:"value"`
	Operator string      `json:"operator"`
}

type Search struct {
	Value    interface{} `json:"search"`
	Target   []string    `json:"search_target"`
	Operator string      `json:"operator"`
}

type FilterData struct {
	Filter *[]Filter `json:"filter"`
	Search *[]Search `json:"search"`
}

func (f *FilterData) FilterQueryBuilder(isStandalone bool) string {

	var filter string
	var search string
	var counter1 int
	var counter2 int
	var result string

	where := "WHERE"

	if f.Filter != nil {
		loop := len(*f.Filter)
		if loop != 0 {
			isFirst := false
			isNull := true
			for _, ft := range *f.Filter {

				if ft.Value == nil || fmt.Sprintf("%v", ft.Value) == "<nil>" {
					continue
				} else {
					isFirst = true
					isNull = false
				}

				value := valueFormater(ft.Value)
				if ft.Operator == "ILIKE" || ft.Operator == "LIKE" {
					first := value[0:1]
					last := value[len(value)-1:]
					if first == "'" && last == "'" {
						value = fmt.Sprintf("'%s%s%s'", "%", value[1:len(value)-1], "%")
					} else {
						value = ("%" + fmt.Sprintf("%v", value) + "%")
					}
				}

				if isFirst {
					filter += fmt.Sprintf(` %s %s %s`, ft.Key, ft.Operator, value)
				} else {
					filter += fmt.Sprintf(` AND %s %s %s`, ft.Key, ft.Operator, value)
				}
			}
			if !isNull {
				counter1 += 1
			}
		}
	}

	if f.Search != nil {
		loop := len(*f.Search)
		if loop != 0 {
			var isNull = true
			for _, s := range *f.Search {
				if s.Value == nil || fmt.Sprintf("%v", s.Value) == "<nil>" {
					continue
				} else {
					isNull = false
				}

				value := valueFormater(s.Value)
				if s.Operator == "ILIKE" || s.Operator == "LIKE" {
					first := value[0:1]
					last := value[len(value)-1:]
					if first == "'" && last == "'" {
						value = fmt.Sprintf("'%s%s%s'", "%", value[1:len(value)-1], "%")
					} else {
						value = ("%" + fmt.Sprintf("%v", value) + "%")
					}
				}

				loop2 := len(s.Target)
				if loop2 != 0 && s.Operator != "" {
					for i := 0; i < loop2; i++ {

						if i == 0 {
							search += fmt.Sprintf(` %s %s %s`, s.Target[i], s.Operator, value)
						} else {
							search += fmt.Sprintf(` OR %s %s %s`, s.Target[i], s.Operator, value)
						}
					}
				}
			}
			if !isNull {
				counter2 += 1
			}
		}
	}

	if isStandalone {
		result += where
	} else {
		result += "AND"
	}
	if counter1+counter2 == 0 {
		return ""
	}
	if counter1 == 1 {
		result += filter
	}
	if counter1+counter2 == 2 {
		result += " AND"
	}
	if counter2 == 1 {
		result += fmt.Sprintf(" (%s)", search)
	}

	return result
}

func valueFormater(value interface{}) string {

	if value == true {
		return "TRUE"
	}
	if value == false {
		return "FALSE"
	}
	if reflect.TypeOf(value).Kind() == reflect.Ptr {
		if val, ok := value.(*int64); ok {
			return fmt.Sprintf("%v", *val)
		}
		if val, ok := value.(*string); ok {
			return fmt.Sprintf(`'%s'`, *val)
		}
	}
	if reflect.TypeOf(value).Kind() == reflect.String {
		return fmt.Sprintf(`'%v'`, value)
	}
	return fmt.Sprintf(`%v`, value)
}
