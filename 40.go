package main

import (
            "strings"
                "code.google.com/p/go-tour/wc"
       )

func WordCount(s string) map[string]int {
        fields := strings.Fields(s)
                        m := make(map[string]int)
                            
                            for str := range(fields) {
                                        m[fields[str]]++
                                                }
                    return m
}

func main() {
        wc.Test(WordCount)
}

