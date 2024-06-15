package task

type Task struct {
	Name string			`json:"name"`
	Runner string		`json:"runner"`
	Command []string	`json:"command"`
	Cleanup bool		`json:"cleanup"`
	CleanupPath string	`json:"cleanup_path"`
	DependsOn []string	`json:"depends_on"`
}