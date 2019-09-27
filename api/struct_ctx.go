package api

type Ctx struct {
	LogPrefix string `json:"-"`
	ReqBody   []byte `json:"ReqBody"`
}

func (c *Ctx) GetCurLogCtx() string {
	if c == nil {
		return ""
	}
	return "todo"
}
