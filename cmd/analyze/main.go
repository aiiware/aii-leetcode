package main

import (
    "fmt"
)

func main() {
    // s1 := "aabb"
    // s2 := "ccdd"
    // s3 := "acbdacbd"
    
    fmt.Println("Analyzing: s1=\"aabb\", s2=\"ccdd\", s3=\"acbdacbd\"")
    fmt.Println("s3 characters: a c b d a c b d")
    fmt.Println("\nPossible interpretation:")
    fmt.Println("Take 'a' from s1[0]")
    fmt.Println("Take 'c' from s2[0]") 
    fmt.Println("Take 'b' from s1[2] ← PROBLEM: skipped s1[1]!")
    fmt.Println("Take 'd' from s2[2] ← PROBLEM: skipped s2[1]!")
    fmt.Println("Take 'a' from s1[1]")
    fmt.Println("Take 'c' from s2[1]")
    fmt.Println("Take 'b' from s1[3]")
    fmt.Println("Take 'd' from s2[3]")
    fmt.Println("\nThis violates the interleaving property because we must use characters in order within each string.")
    fmt.Println("We cannot skip s1[1] to use s1[2] first.")
}
