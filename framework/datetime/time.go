package datetime

import (
	"database/sql/driver"
	"errors"
	"time"
)

// 时间
type Time struct {
	time.Time
}

// 编码为自定义的Json格式
func (t Time) MarshalJSON() ([]byte, error) {

	// 时间为零返回null
	if t.IsZero() {
		return []byte("null"), nil
	}

	return []byte("\"" + t.Format("15:04:05") + "\""), nil
}

// 将Json格式解码
func (t *Time) UnmarshalJSON(data []byte) error {

	var err error

	if len(data) == 2 || string(data) == "null" {
		return err
	}

	var now time.Time

	// 自定义格式解析
	if now, err = time.ParseInLocation("15:04:05", string(data), time.Local); err == nil {
		*t = Time{now}
		return err
	}

	// 带引号的自定义格式解析
	if now, err = time.ParseInLocation("\"15:04:05\"", string(data), time.Local); err == nil {
		*t = Time{now}
		return err
	}

	return err
}

// 转换为数据库值
func (t Time) Value() (driver.Value, error) {

	if t.IsZero() {
		return nil, nil
	}

	return t.Time, nil
}

// 数据库值转换为Time
func (t *Time) Scan(value interface{}) error {

	if val, ok := value.(time.Time); ok {
		*t = Time{Time: val}
		return nil
	}

	return errors.New("无法将值转换为时间戳")
}
