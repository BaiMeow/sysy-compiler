package ast

import (
	"fmt"
	"reflect"
	"sysy/ast/types"
)

type DrawUnit struct {
	Label    string     `json:"label,omitempty"`
	Children []DrawUnit `json:"children,omitempty"`
}

// accept a point to a node, return a DrawUnit
func draw(n any) DrawUnit {
	u := DrawUnit{}
	v := reflect.ValueOf(n).Elem()
	t := reflect.TypeOf(n).Elem()

	if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		return draw(v.Elem().Interface())
	}

	u.Label = t.Name()
	for i := 0; i < t.NumField(); i++ {
		field := v.Field(i)

		// 跳过空值
		if (field.Kind() == reflect.Ptr || field.Kind() == reflect.Interface) && field.IsNil() {
			continue
		}

		// ignore 标签
		if t.Field(i).Tag.Get("draw") == "ignore" {
			continue
		}

		// 如果实现了Draw方法，调用Draw方法
		if field.MethodByName("Draw").Kind() != reflect.Invalid {
			unit := field.MethodByName("Draw").Call(nil)[0].Interface().(DrawUnit)
			unit.Label = fmt.Sprintf("%s: %s", t.Field(i).Name, unit.Label)
			u.Children = append(u.Children, unit)
			continue
		}

		switch field.Interface().(type) {
		case types.Type:
			u.Children = append(u.Children, DrawUnit{
				Label: fmt.Sprintf("%s: %s", t.Field(i).Name, field.Interface().(types.Type).String()),
			})
			continue
		case string:
			u.Children = append(u.Children, DrawUnit{
				Label: fmt.Sprintf("%s: %s", t.Field(i).Name, field.String()),
			})
			continue
		case Expression:
			return draw(field.Interface().(Expression))
		}

		if field.Kind() == reflect.Slice {
			if t.Field(i).Tag.Get("draw") == "unfold" {
				for i := 0; i < field.Len(); i++ {
					u.Children = append(u.Children, draw(field.Index(i).Addr().Interface()))
				}
			} else {
				var arrs []DrawUnit
				for i := 0; i < field.Len(); i++ {
					arrs = append(arrs, draw(field.Index(i).Addr().Interface()))
				}
				u.Children = append(u.Children, DrawUnit{
					Label:    t.Field(i).Name,
					Children: arrs,
				})
			}
			continue
		}

		if field.Kind() == reflect.Struct {
			u.Children = append(u.Children, draw(field.Addr().Interface()))
			continue
		}
	}
	return u
}
