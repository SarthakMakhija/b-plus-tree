package heap_file

import (
	"github.com/SarthakMakhija/b-plus-tree/heap-file/field"
	"github.com/SarthakMakhija/b-plus-tree/heap-file/tuple"
	"github.com/SarthakMakhija/b-plus-tree/index"
	"os"
	"strconv"
	"testing"
)

func TestPutsAndGets10000TuplesByTupleId(t *testing.T) {
	options := DbOptions{
		HeapFileOptions: HeapFileOptions{
			PageSize:                 os.Getpagesize(),
			FileName:                 "./heap.db",
			PreAllocatedPagePoolSize: 6,
			TupleDescriptor: tuple.TupleDescriptor{
				FieldTypes: []field.FieldType{field.StringFieldType{}, field.Uint16FieldType{}},
			},
		},
		IndexOptions: index.DefaultOptions(),
	}
	db, _ := Open(options)
	defer deleteFile(db.bufferPool.file)
	defer deleteFileByName(options.IndexOptions.FileName)

	tupleIds := make([]tuple.TupleId, 10000)
	for iterator := 0; iterator < 10000; iterator++ {
		aTuple := tuple.NewTuple()
		aTuple.AddField(field.NewStringField("Database Systems" + strconv.Itoa(iterator)))
		aTuple.AddField(field.NewUint16Field(uint16(iterator)))

		tupleId, _ := db.Put(aTuple)
		tupleIds[iterator] = tupleId
	}

	for iterator, tupleId := range tupleIds {
		readTuple := db.GetByTupleId(tupleId)

		stringFieldValue := readTuple.AllFields()[0].Value()
		expectedStringFieldValue := "Database Systems" + strconv.Itoa(iterator)

		if stringFieldValue != expectedStringFieldValue {
			t.Fatalf("Expected field value to be %v, received %v", expectedStringFieldValue, stringFieldValue)
		}

		uint16FieldValue := readTuple.AllFields()[1].Value()
		expectedUint16FieldValue := uint16(iterator)

		if uint16FieldValue != expectedUint16FieldValue {
			t.Fatalf("Expected field value to be %v, received %v", expectedUint16FieldValue, uint16FieldValue)
		}
	}
}

func TestPutsAndGets10000TuplesByKey(t *testing.T) {
	options := DbOptions{
		HeapFileOptions: HeapFileOptions{
			PageSize:                 os.Getpagesize(),
			FileName:                 "./heap.db",
			PreAllocatedPagePoolSize: 6,
			TupleDescriptor: tuple.TupleDescriptor{
				FieldTypes: []field.FieldType{field.StringFieldType{}, field.StringFieldType{}, field.Uint16FieldType{}},
			},
		},
		IndexOptions: index.DefaultOptions(),
	}
	db, _ := Open(options)
	defer deleteFile(db.bufferPool.file)
	defer deleteFileByName(options.IndexOptions.FileName)

	tuples := make([]*tuple.Tuple, 10000)
	for iterator := 0; iterator < 10000; iterator++ {
		aTuple := tuple.NewTuple()
		aTuple.AddField(field.NewStringField("Database Systems" + strconv.Itoa(iterator)))
		aTuple.AddField(field.NewStringField("ISBN-" + strconv.Itoa(iterator)))
		aTuple.AddField(field.NewUint16Field(uint16(iterator)))

		_, _ = db.Put(aTuple)
		tuples[iterator] = aTuple
	}

	for iterator, aTuple := range tuples {
		readTuple, _ := db.GetByKey(aTuple.KeyField())

		firstFieldValue := readTuple.AllFields()[0].Value()
		expectedFieldValue := "Database Systems" + strconv.Itoa(iterator)

		if firstFieldValue != expectedFieldValue {
			t.Fatalf("Expected field value to be %v, received %v", expectedFieldValue, firstFieldValue)
		}

		secondFieldValue := readTuple.AllFields()[1].Value()
		expectedFieldValue = "ISBN-" + strconv.Itoa(iterator)

		if secondFieldValue != expectedFieldValue {
			t.Fatalf("Expected field value to be %v, received %v", expectedFieldValue, secondFieldValue)
		}

		uint16FieldValue := readTuple.AllFields()[2].Value()
		expectedUint16FieldValue := uint16(iterator)

		if uint16FieldValue != expectedUint16FieldValue {
			t.Fatalf("Expected field value to be %v, received %v", expectedUint16FieldValue, uint16FieldValue)
		}
	}
}
