package config

type Config struct {
	FilePickerConfig FilePickerConfig `json:"file_picker_config,omitempty"`
}

type FilePickerConfig struct {
	Cursor      string `json:"cursor,omitempty"`
	ShowSize    bool   `json:"show_size,omitempty"`
	ShowMode    bool   `json:"show_mode,omitempty"`
	ShowModTime bool   `json:"show_mod_time,omitempty"`
	ShowContent bool   `json:"show_content,omitempty"`
}

type FilePickerKeybinds struct {
	Up         []string `json:"up,omitempty"`
	Down       []string `json:"down,omitempty"`
	Back       []string `json:"back,omitempty"`
	SelectMode []string `json:"select_mode,omitempty"`
	SelectOne  []string `json:"select_one,omitempty"`
	Enter      []string `json:"enter,omitempty"`
}

type FilePickerStyles struct {
	CurrentFile  StyleConfig `json:"current_file,omitempty"`
	DefaultFile  StyleConfig `json:"default_file,omitempty"`
	Folder       StyleConfig `json:"folder,omitempty"`
	CurrentPath  StyleConfig `json:"current_path,omitempty"`
	ModeStyle    StyleConfig `json:"mode_style,omitempty"`
	ModTimeStyle StyleConfig `json:"mod_time_style,omitempty"`
	SizeStyle    StyleConfig `json:"size_style,omitempty"`
	Selected     StyleConfig `json:"selected,omitempty"`
	CursorStyle  StyleConfig `json:"cursor_style,omitempty"`
}

type StyleConfig struct {
	ForegroundColor string       `json:"foreground_color,omitempty"`
	BackgroundColor string       `json:"background_color,omitempty"`
	Border          BorderConfig `json:"border,omitempty"`
	Padding         []int        `json:"padding,omitempty"`
	Margin          []int        `json:"margin,omitempty"`
	Bold            bool
}

type BorderConfig struct {
	BorderType string `json:"border_type,omitempty"`
	Top        bool   `json:"top,omitempty"`
	Right      bool   `json:"right,omitempty"`
	Bottom     bool   `json:"bottom,omitempty"`
	Left       bool   `json:"left,omitempty"`
}
