package models

type AttributeParam struct {
	Attribute string
	Param     string
}

func NewAP(a string, p string) AttributeParam {
	return AttributeParam{
		Attribute: a,
		Param:     p,
	}
}
