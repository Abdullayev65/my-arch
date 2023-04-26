package time_tool

import "time"

const layout = ""

func Parse(str *string, t **time.Time) (done bool, err error) {
	if str == nil {
		return false, nil
	}

	var t1 time.Time
	t1, err = time.Parse(layout, *str)

	if err != nil {
		return false, err
	}

	*t = &t1

	return true, nil
}

func Format(t *time.Time, str **string) (done bool) {
	if t == nil {
		return false
	}

	s := t.Format(layout)
	*str = &s

	return true
}
