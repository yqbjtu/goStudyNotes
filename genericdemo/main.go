package main

import "fmt"

type Map [KEY int|string, VALUE string|float64] map[KEY]VALUE

func main() {
        // 必须带[xxx,xxx]
        res1 := Map[string, float64]{
                "key1": 0.1,
        }
        fmt.Println(res1)

        res2 := Map[int, string]{
                1: "v1",
        }
        fmt.Println(res2)
}

// 运行结果
// map[key1:0.1]
// map[1:v1]
