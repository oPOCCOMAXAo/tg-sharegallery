package internal

type HandlerType int

const (
	HandlerTypeUnknown HandlerType = iota
	HandlerTypeText
	HandlerTypeCallback
	HandlerTypeInline
)

func (s HandlerType) String() string {
	switch s {
	case HandlerTypeText:
		return "text"
	case HandlerTypeCallback:
		return "callback"
	case HandlerTypeInline:
		return "inline"
	default:
		return "unknown"
	}
}
