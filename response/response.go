package response

// Error is structure holding erroneous response data
type Error struct {
	Msg       string
	Code      int
	Name      string
	Path      string
	Err       string
}
