package tg

func (s *Service) CreateStartLink(
	params string,
) (string, error) {
	return "https://t.me/" + s.me.Username + "?start=" + params, nil
}
