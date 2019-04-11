package main

import (
	"fmt"
	"time"
)

// time包提供了时间的显示和测量用的函数。日历的计算采用的是公历
func main() {

	// 返回当前本地时间
	time.Now()

	// 创建一个本地时间，对应sec和nsec表示的Unix时间（从January 1, 1970 UTC至该时间的秒数和纳秒数）
	time.Unix(1e9, 0)

	// 返回使用给定的名字创建的*time.Location
	// 如果name是""或"UTC"，返回UTC；如果name是"Local"，返回Local；否则name应该是IANA时区数据库里有记录的地点名（该数据库记录了地点和对应的时区），如"America/New_York"
	// 需要的时区数据库可能不是所有系统都提供，特别是非Unix系统。此时LoadLocation会查找环境变量ZONEINFO指定目录或解压该变量指定的zip文件（如果有该环境变量）；然后查找Unix系统的惯例时区数据安装位置，最后查找$GOROOT/lib/time/zoneinfo.zip
	loc, _ := time.LoadLocation("Europe/Berlin")

	// 使用给定的地点名name和时间偏移量offset（单位秒）创建并返回一个Location
	time.FixedZone("Beijing Time", int((8 * time.Hour).Seconds()))

	// 返回指定时间
	t := time.Date(2019, time.February, 7, 0, 0, 0, 0, time.UTC)

	// 返回t的地点和时区信息
	t.Location()

	// 返回采用本地和本地时区，但指向同一时间点的Time
	t.Local()

	// 返回unix时间，即从时间点January 1, 1970 UTC到时间点t所经过的时间（单位秒）
	t.Unix()

	// 返回采用UTC和零时区，但指向同一时间点的Time
	t.UTC()

	// 计算t所在的时区，返回该时区的规范名（如"CET"）和该时区相对于UTC的时间偏移量（单位秒）
	t.Zone()

	// 报告t是否代表Time零值的时间点，January 1, year 1, 00:00:00 UTC
	t.IsZero()

	// 返回采用loc指定的地点和时区，但指向同一时间点的Time。如果loc为nil会panic
	t.In(loc)

	// 将t表示为Unix时间，即从时间点January 1, 1970 UTC到时间点t所经过的时间（单位纳秒）
	// 如果纳秒为单位的unix时间超出了int64能表示的范围，结果是未定义的。注意这就意味着Time零值调用UnixNano方法的话，结果是未定义的
	t.UnixNano()

	// 判断两个时间是否相同，会考虑时区的影响，因此不同时区标准的时间也可以正确比较。本方法和用t==u不同，这种方法还会比较地点和时区信息
	t.Equal(time.Now())

	// 判断t的时间点是否在u之前
	t.Before(time.Now())

	// 判断t的时间点是否在u之后
	t.After(time.Now())

	// 返回时间点t对应的年、月、日
	t.Date()

	// 返回t对应的那一天的时、分、秒
	t.Clock()

	// 返回时间点t对应的年份
	t.Year()

	// 返回时间点t对应的那一年的第几天，平年的返回值范围[1,365]，闰年[1,366]
	t.YearDay()

	// 返回时间点t对应那一年的第几月
	t.Month()

	// 返回时间点t对应那一月的第几日
	t.Day()

	// 返回时间点t对应的那一周的周几
	t.Weekday()

	// 返回时间点t对应的ISO 9601标准下的年份和星期编号
	// 星期编号范围[1,53]，1月1号到1月3号可能属于上一年的最后一周，12月29号到12月31号可能属于下一年的第一周
	t.ISOWeek()

	// 返回t对应的那一天的第几小时，范围[0, 23]
	t.Hour()

	// 返回t对应的那一小时的第几分种，范围[0, 59]
	t.Minute()

	// 返回t对应的那一分钟的第几秒，范围[0, 59]
	t.Second()

	// 返回t对应的那一秒内的纳秒偏移量，范围[0, 999999999]
	t.Nanosecond()

	// 增加2小时, 减少使用负号
	t.Add(time.Hour * 2)

	// 返回增加了给出的年份、月份和天数的时间点Time。例如，时间点January 1, 2011调用AddDate(-1, 2, 3)会返回March 4, 2010
	t.AddDate(1, 2, 3)

	// 获取当前时间减去指定时间的时间
	// 如果结果超出了Duration可以表示的最大值/最小值，将返回最大值/最小值。要获取时间点t-d（d为Duration），可以使用t.Add(-d)
	t.Sub(time.Now())

	// 返回距离t最近的时间点，该时间点应该满足从Time零值到该时间点的时间段能整除d；如果有两个满足要求的时间点，距离t相同，会向上舍入；如果d <= 0，会返回t的拷贝
	t.Round(time.Hour)

	// 类似Round，但是返回的是最接近但早于t的时间点；如果d <= 0，会返回t的拷贝
	t.Truncate(time.Hour)

	// 根据layout指定的格式返回t代表的时间点的格式化文本表示
	t.Format("2006-01-02 15:04:05")

	// 返回采用如下格式字符串的格式化时间
	// "2006-01-02 15:04:05.999999999 -0700 MST"
	t.String()

	// 实现了gob.GobEncoder接口
	gobBytes, _ := t.GobEncode()

	// 实现了gob.GobDecoder接口
	_ = t.GobDecode(gobBytes)

	// 实现了encoding.BinaryMarshaler接口
	binaryBytes, _ := t.MarshalBinary()

	// 实现了encoding.BinaryUnmarshaler接口
	_ = t.UnmarshalBinary(binaryBytes)

	// 实现了json.Marshaler接口
	jsonBytes, _ := t.MarshalJSON()

	// 实现了json.Unmarshaler接口
	_ = t.UnmarshalJSON(jsonBytes)

	// 实现了encoding.TextMarshaler接口
	textBytes, _ := t.MarshalText()

	// 实现了encoding.TextUnmarshaler接口
	_ = t.UnmarshalText(textBytes)

	// 解析一个格式化的时间字符串并返回它代表的时间
	// layout规定了参考格式
	time.Parse("2006 Jan 02 15:04:05", "2019 Feb 07 12:15:30.918273645")

	// 类似Parse但有两个重要的不同之处
	// 第一，当缺少时区信息时，Parse将时间解释为UTC时间，而ParseInLocation将返回值的Location设置为loc
	// 第二，当时间字符串提供了时区偏移量信息时，Parse会尝试去匹配本地时区，而ParseInLocation会去匹配loc
	time.ParseInLocation("2006 Jan 02 15:04:05", "2019 Feb 07 12:15:30.918273645", loc)

	// 解析一个时间段字符串
	// 一个时间段字符串是一个序列，每个片段包含可选的正负号、十进制数、可选的小数部分和单位后缀，如"300ms"、"-1.5h"、"2h45m"
	// 合法的单位有"ns"、"us" /"µs"、"ms"、"s"、"m"、"h"
	d, _ := time.ParseDuration("1h15m30.918273645s")
	fmt.Println(d.Hours(), d.Minutes(), d.Nanoseconds())

	// 返回从t到现在经过的时间，等价于time.Now().Sub(t)
	time.Since(t)

	// 阻塞当前go程至少d代表的时间段。d<=0时，Sleep会立刻返回
	time.Sleep(time.Second * 2)

	// 创建一个Timer，它会在最少过去时间段d后到期，向其自身的C字段发送当时的时间
	time.NewTimer(time.Minute * 2)

	// 会在另一线程经过时间段d后向返回值发送当时的时间。等价于NewTimer(d).C
	time.After(time.Minute * 2)

	// 另起一个go程等待时间段d过去，然后调用f。它返回一个Timer，可以通过调用其Stop方法来取消等待和对f的调用
	time.AfterFunc(time.Millisecond*1000, func() {
		fmt.Println("time after 1 second test")
	})

	// 返回一个新的Ticker，该Ticker包含一个通道字段，并会每隔时间段d就向该通道发送当时的时间
	// 它会调整时间间隔或者丢弃tick信息以适应反应慢的接收者
	// 如果d<=0会panic。关闭该Ticker可以释放相关资源
	time.NewTicker(time.Second * 10)

	// 是NewTicker的封装，只提供对Ticker的通道的访问。如果不需要关闭Ticker，本函数就很方便
	time.Tick(time.Second * 10)
}
