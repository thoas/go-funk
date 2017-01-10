go-funk
=======

.. image:: https://secure.travis-ci.org/thoas/go-funk.svg?branch=master
    :alt: Build Status
    :target: http://travis-ci.org/thoas/go-funk

.. image:: https://godoc.org/github.com/thoas/go-funk?status.svg
    :alt: GoDoc
    :target: https://godoc.org/github.com/thoas/go-funk

.. image:: https://goreportcard.com/badge/github.com/thoas/go-funk
    :alt: Go report
    :target: https://goreportcard.com/report/github.com/thoas/go-funk

``go-funk`` is a modern Go library based on reflect_.

As it relies a lot on reflect_, be careful this code runs exclusively on runtime so you must have a good test suite.

This project has started as an experiment to learn reflect_. It may looks like lodash_ in some aspects but
it will have its own ROADMAP, lodash_ is an awesome library with a lot of works behind it, all features included in
``go-funk`` come from internal use cases.

Why this name?
--------------

Long story, short answer because ``func`` is a reserved word in Go, I wanted something similar.

Initially this project was named ``fn`` I don't need to explain why that was a bad idea for french speakers :)

Let's ``funk``!

.. image:: https://media.giphy.com/media/3oEjHQKtDXpeGN9rW0/giphy.gif

<3

Installation
------------

.. code-block:: bash

    go get github.com/thoas/go-funk

Usage
-----

.. code-block:: go

    import "github.com/thoas/go-funk"

These examples will be based on the following data model:

.. code-block:: go

    type Foo struct {
        ID        int
        FirstName string `tag_name:"tag 1"`
        LastName  string `tag_name:"tag 2"`
        Age       int    `tag_name:"tag 3"`
    }

    func (f Foo) TableName() string {
        return "foo"
    }

With fixtures:

.. code-block:: go

    f := &Foo{
        ID:        1,
        FirstName: "Foo",
        LastName:  "Bar",
        Age:       30,
    }

You can import ``go-funk`` using a basic statement:

.. code-block:: go

    import "github.com/thoas/go-funk"

funk.Contains
.............

``funk.Contains`` returns true if an element is present in a iteratee (slice, map, string).

One frustrating thing in Go is to implement ``contains`` methods for each types, for example:

.. code-block:: go

    func ContainsInt(s []int, e int) bool {
        for _, a := range s {
            if a == e {
                return true
            }
        }
        return false
    }

this can be replaced by ``funk.Contains``:

.. code-block:: go

    // slice of string
    funk.Contains([]string{"foo", "bar"}, "bar") // true

    // slice of Foo ptr
    funk.Contains([]*Foo{f}, f) // true
    funk.Contains([]*Foo{f}, nil) // false

    b := &Foo{
        ID:        2,
        FirstName: "Florent",
        LastName:  "Messa",
        Age:       28,
    }

    funk.Contains([]*Foo{f}, b) // false

    // string
    funk.Contains("florent", "rent") // true
    funk.Contains("florent", "foo") // false

    // even map
    funk.Contains(map[int]string{1: "Florent"}, 1) // true

funk.IndexOf
............

``funk.IndexOf`` gets the index at which the first occurrence of value is found in array or return -1
if the value cannot be found.

.. code-block:: go

    // slice of string
    funk.IndexOf([]string{"foo", "bar"}, "bar") // 1
    funk.IndexOf([]string{"foo", "bar"}, "gilles") // -1

funk.LastIndexOf
................

``funk.LastIndexOf`` gets the index at which the last occurrence of value is found in array or return -1
if the value cannot be found.

.. code-block:: go

    // slice of string
    funk.LastIndexOf([]string{"foo", "bar", "bar"}, "bar") // 2
    funk.LastIndexOf([]string{"foo", "bar"}, "gilles") // -1

funk.ToMap
..........

``funk.ToMap`` transforms a slice of structs to a map based on a ``pivot`` field.

.. code-block:: go

    f := &Foo{
        ID:        1,
        FirstName: "Gilles",
        LastName:  "Fabio",
        Age:       70,
    }

    b := &Foo{
        ID:        2,
        FirstName: "Florent",
        LastName:  "Messa",
        Age:       80,
    }

    results := []*Foo{f, b}

    mapping := funk.ToMap(results, "ID") // map[int]*Foo{1: f, 2: b}

funk.Filter
...........

``funk.Filter`` filters a slice based on a predicate.

.. code-block:: go

    r := funk.Filter([]int{1, 2, 3, 4}, func(x int) bool {
        return x%2 == 0
    }) // []int{2, 4}

funk.Find
.........

``funk.Find`` finds an element in a slice based on a predicate.

.. code-block:: go

    r := funk.Find([]int{1, 2, 3, 4}, func(x int) bool {
        return x%2 == 0
    }) // 2

funk.Map
........

``funk.Map`` allows you to manipulate an iteratee (map, slice) and to transform it to another type:

* map -> slice
* map -> map
* slice -> map
* slice -> slice

.. code-block:: go

    r := funk.Map([]int{1, 2, 3, 4}, func(x int) int {
        return "Hello"
    }) // []int{2, 4, 6, 8}

    r := funk.Map([]int{1, 2, 3, 4}, func(x int) string {
        return "Hello"
    }) // []string{"Hello", "Hello", "Hello", "Hello"}

    r = funk.Map([]int{1, 2, 3, 4}, func(x int) (int, int) {
        return x, x
    }) // map[int]int{1: 1, 2: 2, 3: 3, 4: 4}

    mapping := map[int]string{
        1: "Florent",
        2: "Gilles",
    }

    r = funk.Map(mapping, func(k int, v string) int {
        return k
    }) // []int{1, 2}

    r = funk.Map(mapping, func(k int, v string) (string, string) {
        return fmt.Sprintf("%d", k), v
    }) // map[string]string{"1": "Florent", "2": "Gilles"}

funk.Get
........

``funk.Get`` retrieves the value at path of struct(s).

.. code-block:: go

    var bar *Bar = &Bar{
        Name: "Test",
        Bars: []*Bar{
            &Bar{
                Name: "Level1-1",
                Bar: &Bar{
                    Name: "Level2-1",
                },
            },
            &Bar{
                Name: "Level1-2",
                Bar: &Bar{
                    Name: "Level2-2",
                },
            },
        },
    }

    var foo *Foo = &Foo{
        ID:        1,
        FirstName: "Dark",
        LastName:  "Vador",
        Age:       30,
        Bar:       bar,
        Bars: []*Bar{
            bar,
            bar,
        },
    }

    funk.Get([]*Foo{foo}, "Bar.Bars.Bar.Name") // []string{"Level2-1", "Level2-2"}
    funk.Get(foo, "Bar.Bars.Bar.Name") // []string{"Level2-1", "Level2-2"}
    funk.Get(foo, "Bar.Name") // Test

``funk.Get`` also handles ``nil`` values:

.. code-block:: go

    bar := &Bar{
        Name: "Test",
    }

    foo1 := &Foo{
        ID:        1,
        FirstName: "Dark",
        LastName:  "Vador",
        Age:       30,
        Bar:       bar,
    }

    foo2 := &Foo{
        ID:        1,
        FirstName: "Dark",
        LastName:  "Vador",
        Age:       30,
    } // foo2.Bar is nil

    funk.Get([]*Foo{foo1, foo2}, "Bar.Name") // []string{"Test"}
    funk.Get(foo2, "Bar.Name") // nil

funk.Keys
.........

``funk.Keys`` creates an array of the own enumerable map keys or struct field names.

.. code-block:: go

    funk.Keys(map[string]int{"one": 1, "two": 2}) // []string{"one", "two"} (iteration order is not guaranteed)

    foo := &Foo{
        ID:        1,
        FirstName: "Dark",
        LastName:  "Vador",
        Age:       30,
    }

    funk.Keys(foo) // []string{"ID", "FirstName", "LastName", "Age"} (iteration order is not guaranteed)

funk.Values
...........

``funk.Values`` creates an array of the own enumerable map values or struct field values.

.. code-block:: go

    funk.Values(map[string]int{"one": 1, "two": 2}) // []string{1, 2} (iteration order is not guaranteed)

    foo := &Foo{
        ID:        1,
        FirstName: "Dark",
        LastName:  "Vador",
        Age:       30,
    }

    funk.Values(foo) // []interface{}{1, "Dark", "Vador", 30} (iteration order is not guaranteed)

funk.ForEach
............

``funk.ForEach`` allows you to range over an iteratee (map, slice)

.. code-block:: go

    funk.ForEach([]int{1, 2, 3, 4}, func(x int) {
        fmt.Println(x)
    })

funk.ForEachRight
............

``funk.ForEachRight`` allows you to range over an iteratee (map, slice) from the right

.. code-block:: go

    results := []int{}

    funk.ForEachRight([]int{1, 2, 3, 4}, func(x int) {
        results = append(results, x)
    })

    fmt.Println(results) // []int{4, 3, 2, 1}

funk.Chunk
..........

``funk.Chunk`` creates an array of elements split into groups with the length
of the size. If array can't be split evenly, the final chunk will be the remaining element.

.. code-block:: go

    funk.Chunk([]int{1, 2, 3, 4, 5}, 2) // [][]int{[]int{1, 2}, []int{3, 4}, []int{5}}

funk.FlattenDeep
................

``funk.FlattenDeep`` recursively flattens array.

.. code-block:: go

    funk.FlattenDeep([][]int{[]int{1, 2}, []int{3, 4}}) // []int{1, 2, 3, 4}

funk.Uniq
.........

``funk.Uniq`` creates an array with unique values.

.. code-block:: go

    funk.Uniq([]int{0, 1, 1, 2, 3, 0, 0, 12}) // []int{0, 1, 2, 3, 12}

funk.Shuffle
............

``funk.Shuffle`` creates an array of shuffled values

.. code-block:: go

    funk.Shuffle([]int{0, 1, 2, 3, 4}) // []int{2, 1, 3, 4, 0}

funk.Reverse
............

``funk.Reverse`` transforms an array the first element will become the last,
the second element will become the second to last, etc.

.. code-block:: go

    funk.Reverse([]int{0, 1, 2, 3, 4}) // []int{4, 3, 2, 1, 0}

funk.SliceOf
............

``funk.SliceOf`` will return a slice based on an element.

.. code-block:: go

    funk.SliceOf(f) // will return a []*Foo{f}

funk.RandomInt
..............

``funk.RandomInt`` generates a random int, based on a min and max values

.. code-block:: go

    funk.RandomInt(0, 100) // will be between 0 and 100

funk.RandomString
.................

``funk.RandomString`` generates a random string with a fixed length

.. code-block:: go

    funk.RandomString(4) // will be a string of 4 random characters

funk.Shard
..........

``funk.Shard`` generates a sharded string with a fixed length and depth

.. code-block:: go

    funk.Shard("e89d66bdfdd4dd26b682cc77e23a86eb", 1, 2, false) // []string{"e", "8", "e89d66bdfdd4dd26b682cc77e23a86eb"}

    funk.Shard("e89d66bdfdd4dd26b682cc77e23a86eb", 2, 2, false) // []string{"e8", "9d", "e89d66bdfdd4dd26b682cc77e23a86eb"}

    funk.Shard("e89d66bdfdd4dd26b682cc77e23a86eb", 2, 2, true) // []string{"e8", "9d", "66", "bdfdd4dd26b682cc77e23a86eb"}


Contributing
------------

* Ping me on twitter `@thoas <https://twitter.com/thoas>`_
* Fork the `project <https://github.com/thoas/go-funk>`_
* Fix `open issues <https://github.com/thoas/go-funk/issues>`_ or request new features

Don't hesitate ;)

.. _reflect: https://golang.org/pkg/reflect/
.. _lodash: https://lodash.com/
