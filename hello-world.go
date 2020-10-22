package main

import (
    "crypto/sha1"
    "fmt"
)

func main() {

    m := make(map[string]int)

    m["k1"] = 7
    m["k2"] = 13

    fmt.Println("map:", m)

    v1 := m["k1"]
    fmt.Println("v1: ", v1)

    fmt.Println("len:", len(m))

    delete(m, "k2")
    fmt.Println("map:", m)

    _, prs := m["k2"]
    fmt.Println("prs:", prs)

    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)


    s := "sha1 this string"
    h := sha1.New()
    h.Write([]byte(s))
    bs := h.Sum(nil)
    fmt.Println(s)
    fmt.Printf("%x\n", bs)
}
