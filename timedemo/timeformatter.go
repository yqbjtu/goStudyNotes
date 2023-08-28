package main

import (
	"fmt"
	"time"
)

func displayDate(format string) {
	fmt.Println(time.Now().Format(format))
}

func main() {
	// The output time will vary for you
	// but you can check the time format in which we get the output
	displayDate("20060102150405")      // YYYYMMDDHHMMSS Format
	displayDate("2006/01/02 15/04/05") // YYYY/MM/DD HH/MM/SS Format
	// time.RFC3339 	RFC3339     = "2006-01-02T15:04:05Z07:00"
	//java     mapper.registerModule(new JodaModule()); 使用其中对ZonedDateTime 和LocalDateTime的序列化
	//"timestamp": "2023-07-24T21:58:11.480535414+08:00",, 对Java ZonedDateTime
	displayDate("2006-01-02T15:04:05.000000000Z07:00") // yyyy-MM-dd'T'HH:mm:ss.SSSXXX
	// "timestamp": "2023-02-26T04:10:52.931726137", 对Java LocalDateTime
	displayDate("2006-01-02T15:04:05.000000000") // yyyy-MM-dd'T'HH:mm:ss.SSSXXX

	displayDate("02/01/2006 15/04/05")               // DD/MM/YYYY HH/MM/SS Format
	displayDate("02/01/06 15/04/05")                 // DD/MM/YY HH/MM/SS Format
	displayDate("2006-01-02 15:04:05")               // YYYY-MM-DD HH:MM:SS Format
	displayDate("15:04:05, 2006-Jan-02 Mon")         // HH:MM:SS, YYYY-Mon-DD DayShort Format
	displayDate("15:04:05, 2006-Jan-02 Monday")      // HH:MM:SS, YYYY-Mon-DD DayLong Format
	displayDate("15:04:05, 2006-January-02 MST Mon") // HH:MM:SS, YYYY-Mon-DD TZ DayShort Format
	displayDate("3:4:05, 2006-1-02 MST Mon")         // H:M:SS, YYYY-M-DD TZ DayShort Format
	displayDate("3:4:05 pm, 2006-1-02 MST Mon")      // H:M:SS Part_of_day(small), YYYY-M-DD TZ DayShort Format
	displayDate("3:4:05 PM, 2006-1-02 MST Mon")      // H:M:SS Part_of_day(small), YYYY-M-DD TZ DayShort Format
	displayDate("03:04:05.000")                      // HH:MM:SS.MilliSeconds
	displayDate("03:04:05.000000")                   // HH:MM:SS.MicroSeconds
	displayDate("03:04:05.000000000")                // HH:MM:SS.NanoSeconds

}

/* 运行结果
20230827223553
2023/08/27 22/35/53
2023-08-27T22:35:53.709042000+08:00
2023-08-27T22:35:53.709044000
27/08/2023 22/35/53
27/08/23 22/35/53
2023-08-27 22:35:53
22:35:53, 2023-Aug-27 Sun
22:35:53, 2023-Aug-27 Sunday
22:35:53, 2023-August-27 CST Sun
10:35:53, 2023-8-27 CST Sun
10:35:53 pm, 2023-8-27 CST Sun
10:35:53 PM, 2023-8-27 CST Sun
10:35:53.709
10:35:53.709057
10:35:53.709058000
*/
