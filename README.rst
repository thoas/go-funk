go-fn
=====

``go-fn`` is a modern Go library based on reflect_.

This project has been started as an experiment to learn reflect_, feel free to use it :)

Installation
------------

.. code-block:: bash

    go get github.com/thoas/go-fn

Usage
-----

.. code-block:: go

    import "github.com/thoas/go-fn"

These examples will be based on following data model:

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

    # fixtures
	f := &Foo{
		ID:        1,
		FirstName: "Drew",
		LastName:  "Olson",
		Age:       30,
		Bar: &Bar{
			Name: "Test",
		},
	}

You can import ``go-fn`` using a basic statement:

.. code-block:: go

    import "github.com/thoas/go-fn"

fn.SliceOf
..........

``SliceOf`` will return a slice based on an element.

.. code-block:: go

	result := fn.SliceOf(f) # will return a []*Foo

fn.Contains
...........

``Contains`` will return if an element is present in a iterable container (slice, map, string).

It's one frustrating thing in Go to implement ``contains`` methods for each types, for example:

.. code-block:: go

    func ContainsInt(s []int, e int) bool {
        for _, a := range s {
            if a == e {
                return true
            }
        }
        return false
    }

this can be replaced by calling:

.. code-block:: go

    # slice of string
    fn.Contains([]string{"foo", "bar"}, "bar") # true

    # slice of *Foo
    fn.Contains([]*Foo{f}, f) # true

    # string
    fn.Contains("florent", "rent") # true

    # even map
    fn.Contains(map[int]string{1: "Florent"}, 1) # true

.. _reflect: https://golang.org/pkg/reflect/
