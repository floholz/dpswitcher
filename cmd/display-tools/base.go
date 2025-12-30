package display_tools

type Display struct {
	ID        string
	Primary   bool
	Active    bool
	Connected bool
}

type DPConfigTool interface {
	Info() string
	ListDisplays() ([]Display, error)
	GetDisplay(id string) (Display, error)
	EnableDisplay(id string) error
	DisableDisplay(id string) error
	ToggleDisplay(id string) error
}
