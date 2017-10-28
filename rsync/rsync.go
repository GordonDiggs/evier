package rsync

type Options struct {
	User   string `yaml:",omitempty"`
	Host   string `yaml:",omitempty"`
	Path   string `yaml:",omitempty"`
	Rsh    string `yaml:",omitempty"`
	Delete bool
}
