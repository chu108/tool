package tool

import "time"

const _Date = "2006-01-02"
const _Time = "15:04:05"
const _DateTime = _Date + " " + _Time

//返回年月日
func GetDate() string {
	return time.Now().Format(_Date)
}

//返回年月日时分秒
func GetTime() string {
	return time.Now().Format(_DateTime)
}

//时间戳转日期
func UnixToDateTime(t int64) string {
	return time.Unix(t, 0).Format(_DateTime)
}

//日期转时间戳
func DateTimeToUnix(t string) (int64, error) {
	tm, err := time.Parse(t, _DateTime)
	if err != nil {
		return 0, err
	}
	return tm.Unix(), nil
}

//毫秒转日期
func MilliSecondToDateTime(t int64) string {
	return time.Unix(0, t*int64(time.Millisecond)).Format(_DateTime)
}

/**
睡眠指定时间，如：5-10 之间的数
max: 10 最大数
min: 5 最小数
t: Second 秒 Millisecond 毫秒
*/
func SleepRand(max, min int, t string) {
	switch t {
	case "Second":
		time.Sleep(time.Second * time.Duration(GetRand(max, min)))
	case "Millisecond":
		time.Sleep(time.Millisecond * time.Duration(GetRand(max, min)))
	}
}
