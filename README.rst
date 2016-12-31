go-fn
=====

``go-fn`` is a modern Go library based on reflect_.

This project has started as an experiment to learn reflect_, feel free to use it :)

.. image:: https://secure.travis-ci.org/thoas/go-fn.svg?branch=master
    :alt: Build Status
    :target: http://travis-ci.org/thoas/go-fn

.. image:: https://godoc.org/github.com/thoas/go-fn?status.svg
    :alt: GoDoc
    :target: https://godoc.org/github.com/thoas/go-fn

.. image:: https://goreportcard.com/badge/github.com/thoas/go-fn
    :alt: Go report
    :target: https://goreportcard.com/report/github.com/thoas/go-fn


Installation
------------

.. code-block:: bash

    go get github.com/thoas/go-fn

Usage
-----

.. code-block:: go

    import "github.com/thoas/go-fn"

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

You can import ``go-fn`` using a basic statement:

.. code-block:: go

    import "github.com/thoas/go-fn"

fn.SliceOf
..........

``fn.SliceOf`` will return a slice based on an element.

.. code-block:: go

	result := fn.SliceOf(f) // will return a []*Foo{f}

fn.Contains
...........

``fn.Contains`` returns true if an element is present in a iteratee (slice, map, string).

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

this can be replaced by ``fn.Contains``:

.. code-block:: go

    // slice of string
    fn.Contains([]string{"foo", "bar"}, "bar") // true

    // slice of *Foo
    fn.Contains([]*Foo{f}, f) // true
    fn.Contains([]*Foo{f}, nil) // false

	b := &Foo{
		ID:        2,
		FirstName: "Florent",
		LastName:  "Messa",
		Age:       28,
	}

    fn.Contains([]*Foo{f}, b) // false

    // string
    fn.Contains("florent", "rent") // true
    fn.Contains("florent", "foo") // false

    // even map
    fn.Contains(map[int]string{1: "Florent"}, 1) // true

fn.ToMap
........

``fn.ToMap`` transforms a slice of structs to a map based on a ``pivot`` field.

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

	mapping := fn.ToMap(results, "ID") // map[int]*Foo{1: f, 2: b}

fn.Filter
.........

``fn.Filter`` filters a slice based on a predicate.

.. code-block:: go

	r := fn.Filter([]int{1, 2, 3, 4}, func(x int) bool {
		return x%2 == 0
	}) // []int{2, 4}

fn.Find
.........

``fn.Find`` finds an element in a slice based on a predicate.

.. code-block:: go

	r := fn.Find([]int{1, 2, 3, 4}, func(x int) bool {
		return x%2 == 0
	}) // 2

fn.Map
......

``fn.Map`` allows you to manipulate an iteratee (map, slice) and to transform it to another type:

* map -> slice
* map -> map
* slice -> map
* slice -> slice

.. code-block:: go

	r := fn.Map([]int{1, 2, 3, 4}, func(x int) int {
		return "Hello"
	}) // []int{2, 4, 6, 8}

	r := fn.Map([]int{1, 2, 3, 4}, func(x int) string {
		return "Hello"
	}) // []string{"Hello", "Hello", "Hello", "Hello"}

	r = fn.Map([]int{1, 2, 3, 4}, func(x int) (int, int) {
		return x, x
	}) // map[int]int{1: 1, 2: 2, 3: 3, 4: 4}

	mapping := map[int]string{
		1: "Florent",
		2: "Gilles",
	}

	r = fn.Map(mapping, func(k int, v string) int {
		return k
	}) // []int{1, 2}

	r = fn.Map(mapping, func(k int, v string) (string, string) {
		return fmt.Sprintf("%d", k), v
	}) // map[string]string{"1": "Florent", "2": "Gilles"}

fn.ForEach
..........

``fn.ForEach`` allows you to range over an iteratee (map, slice)

.. code-block:: go

	fn.ForEach([]int{1, 2, 3, 4}, func(x int) {
		fmt.Println(x)
	})

.. _reflect: https://golang.org/pkg/reflect/
