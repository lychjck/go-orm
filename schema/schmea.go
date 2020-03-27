package schema

import (
	"go-orm/dialect"
	"go/ast"
	"reflect"
)

type Field struct {
	Name string
	Type string
	Tag  string
}

type Schema struct {
	Model      interface{}
	Name       string
	Fields     []*Field
	FieldNames []string
	FieldMap   map[string]*Field
}

func (schema *Schema) GetField(name string) *Field  {
	return schema.FieldMap[name]
}

func Parse(dest interface{},d dialect.Dialect) *Schema  {
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	//fmt.Println(modelType)
	//fmt.Println(modelType.Kind())
	//fmt.Println(modelType.Name())
	schema := &Schema{
		Model: dest,
		Name: modelType.Name(),
		FieldMap: make(map[string]*Field),
	}

	for i:=0;i<modelType.NumField();i++{
		p:=modelType.Field(i)
		if !p.Anonymous && ast.IsExported(p.Name){
			field := &Field{
				Name: p.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}
			if v,ok := p.Tag.Lookup("goorm");ok{
				field.Tag=v
			}
			schema.Fields = append(schema.Fields,field)
			schema.FieldNames = append(schema.FieldNames,p.Name)
			schema.FieldMap[p.Name] = field
		}
	}
	return schema
}