package models

type Command struct {
	Flower string `json:"id,omitempty" form:"id"`
	Action string `json:"action,omitempty" form:"action"`
	Value int32 `json:"value,omitempty" form:"value"`
}
