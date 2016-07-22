package models

type Control struct {
	Action string `json:"action,omitempty" form:"action"`
	Value int32 `json:"value,omitempty" form:"value"`
}

