package funk

import "database/sql"

type Model interface {
	TableName() string
}

// Bar is
type Bar struct {
	Name string `tag_name:"BarName"`
	Bar  *Bar
	Bars []*Bar
}

func (b Bar) TableName() string {
	return "bar"
}

// Foo is
type Foo struct {
	ID         int
	FirstName  string `tag_name:"tag 1"`
	LastName   string `tag_name:"tag 2"`
	Age        int    `tag_name:"tag 3"`
	Bar        *Bar   `tag_name:"tag 4"`
	Bars       []*Bar
	EmptyValue sql.NullInt64

	BarInterface     interface{}
	BarPointer       interface{}
	GeneralInterface interface{}

	ZeroBoolValue bool
	ZeroIntValue  int
	ZeroIntPtrValue *int
}

func (f Foo) TableName() string {
	return "foo"
}

var bar = &Bar{
	Name: "Test",
	Bars: []*Bar{
		{
			Name: "Level1-1",
			Bar: &Bar{
				Name: "Level2-1",
			},
		},
		{
			Name: "Level1-2",
			Bar: &Bar{
				Name: "Level2-2",
			},
		},
	},
}

var foo = &Foo{
	ID:        1,
	FirstName: "Dark",
	LastName:  "Vador",
	Age:       30,
	Bar:       bar,
	EmptyValue: sql.NullInt64{
		Valid: true,
		Int64: 10,
	},
	Bars: []*Bar{
		bar,
		bar,
	},
	BarInterface: bar,
	BarPointer:   &bar,
	ZeroBoolValue: false,
	ZeroIntValue: 0,
	ZeroIntPtrValue: nil,
}

var foo2 = &Foo{
	ID:        1,
	FirstName: "Dark",
	LastName:  "Vador",
	Age:       30,
}

var m1 = map[string]interface{}{
	"id":        1,
	"firstname": "dark",
	"lastname":  "vador",
	"age":       30,
	"bar": map[string]interface{}{
		"name": "test",
		"bars": []map[string]interface{}{
			{
				"name": "level1-1",
				"bar": map[string]interface{}{
					"name": "level2-1",
				},
			},
			{
				"name": "level1-2",
				"bar": map[string]interface{}{
					"name": "level2-2",
				},
			},
		},
	},
}

var m2 = map[string]interface{}{
	"id":        1,
	"firstname": "dark",
	"lastname":  "vador",
	"age":       30,
}
