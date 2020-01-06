package person

type Service struct {
}

func (s *Service) PaymentRatio() map[string]float64 {
	return map[string]float64{
		"person-0": 0.66,
		"person-1": 0.34,
	}
}
