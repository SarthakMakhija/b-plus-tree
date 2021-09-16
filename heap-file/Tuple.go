package heap_file

import "github.com/SarthakMakhija/b-plus-tree/heap-file/field"

type Tuple struct {
	fields []field.Field
}

func NewTuple() *Tuple {
	var fields []field.Field
	return &Tuple{
		fields: fields,
	}
}

func (tuple *Tuple) AddField(field field.Field) {
	tuple.fields = append(tuple.fields, field)
}

func (tuple Tuple) MarshalBinary() ([]byte, int) {
	var buffer []byte
	for _, aField := range tuple.fields {
		buffer = append(buffer, aField.MarshalBinary()...)
	}
	return buffer, len(buffer)
}

func (tuple *Tuple) UnMarshalBinary(buffer []byte, fieldTypes []field.FieldType) {
	for _, fieldType := range fieldTypes {
		aField := fieldType.UnMarshalBinary(buffer)
		tuple.fields = append(tuple.fields, aField)
	}
}
