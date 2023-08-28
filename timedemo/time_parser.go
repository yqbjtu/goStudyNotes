package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time parsing")

	dateString := "2014-11-09T12:40:45.15Z"
	dateString2 := "2022-10-12T10:15:16.371+07:00"

	time1, err := time.Parse(time.RFC3339, dateString)
	if err != nil {
		fmt.Println("Error while parsing date :", err)
	}
	fmt.Println(time1)

	time2, err := time.Parse(time.RFC3339, dateString2)
	if err != nil {
		fmt.Println("Error while parsing date :", err)
	}
	fmt.Println(time2)

	dateString3 := "2007-10-09 22:50:02.023"
	time3, err := time.Parse("2006-01-02 15:04:05.000", dateString3)
	if err != nil {
		fmt.Println("Error while parsing date :", err)
	}
	fmt.Println(time3)

	timestampStr := "2023-07-24T21:58:11.480535414+08:00" // 对应Java ZonedDateTime，
	// yyyy-MM-dd'T'HH:mm:ss.SSSXXX
	time4, err := time.Parse("2006-01-02T15:04:05.000000000Z07:00", timestampStr)
	if err != nil {
		fmt.Println("Error while parsing date :", err)
	}
	fmt.Println(time4)

	timestampStr = "2023-02-26T04:10:52.931726137" // 对应Java LocalDateTime
	//yyyy-MM-dd'T'HH:mm:ss.SSSXXX
	time5, err := time.Parse("2006-01-02T15:04:05.000000000", timestampStr)
	if err != nil {
		fmt.Println("Error while parsing date :", err)
	}
	fmt.Println(time5)
}

/*运行结果
Time parsing
2014-11-09 12:40:45.15 +0000 UTC
2022-10-12 10:15:16.371 +0700 +0700
2007-10-09 22:50:02.023 +0000 UTC
2023-07-24 21:58:11.480535414 +0800 CST
2023-02-26 04:10:52.931726137 +0000 UTC
*/
