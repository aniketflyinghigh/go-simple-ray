package main

import (
    "testing"
    . "github.com/franela/goblin"
    "math"
)

func Test(t *testing.T) {
    g := Goblin(t)
    g.Describe("Vector", func() {
        v1 := Vector{1, 2, 3}
        v2 := Vector{4, 5, 6}

        g.It("Should add two Vectors ", func() {    
            g.Assert(v1.add(v2)).Equal(Vector{5, 7, 9})
        })

        g.It("Should subtract two Vectors ", func() {    
            g.Assert(v1.subtract(v2)).Equal(Vector{-3, -3, -3})
        })

        g.It("Should multiply two Vectors ", func() {    
            g.Assert(v1.multiply(v2)).Equal(Vector{4, 10, 18})
        })

        g.It("Should divide two Vectors ", func() {    
            g.Assert(v1.divide(v2)).Equal(Vector{0.25, 0.4, 0.5})
        })

        g.It("Should dot product two Vectors ", func() {    
            g.Assert(v1.dot(v2)).Equal(32)
        })

        g.It("Should have a length", func() {    
            g.Assert(v1.length()).Equal(math.Sqrt(v1.lengthSquared()))
        })

        g.It("Should have a lengthsquared", func() {    
            g.Assert(v1.lengthSquared()).Equal(v1.length() * v1.length())
        })
    })
}