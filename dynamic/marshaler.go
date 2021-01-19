package dynamic

import "encoding/json"

func (c Configuration) MarshalJSON() ([]byte, error) {
	return json.Marshal(c)
}
