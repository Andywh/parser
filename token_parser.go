package main

func parse(ts []token) interface{}  {
	t := ts[0]
	ts = append(ts[:0],ts[1:]...)
	if t.token_type == braceLeft {
		obj := make(map[token]interface{})
		for ; ts[0].token_type != braceRight; {
			k := ts[0]
			//_colon := ts[1]
			// 确保 k.token_type 必须是 string
			// 确保 _colon 必须是 colon
			ts = append(ts[:0],ts[1:]...)
			ts = append(ts[:0],ts[1:]...)
			v := parse(ts)
			obj[k] = v
			_comma := ts[0]
			if _comma.token_type == comma {
				ts = append(ts[:0],ts[1:]...)
			}
		}
		ts = append(ts[:0],ts[1:]...)
		return obj
	} else if t.token_type == bracketLeft {
		l := make([]interface{}, 0)
		for ;ts[0].token_type != bracketRight; {
			v := parse(ts)
			_comma := ts[0]
			if _comma.token_type == comma {
				ts = append(ts[:0],ts[1:]...)
			}
			l = append(l, v)
		}
		ts = append(ts[:0],ts[1:]...)
		return l
	} else {
		return t
	}
	return nil
}