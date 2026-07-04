package model

type Symbol struct {
	Mask string
}

func (s *Symbol) GetMask () string {
	return s.Mask
}