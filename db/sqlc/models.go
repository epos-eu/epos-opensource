// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package sqlc

type Docker struct {
	Name           string
	Directory      string
	ApiUrl         string
	GuiUrl         string
	BackofficeUrl  string
	ApiPort        int64
	GuiPort        int64
	BackofficePort int64
}

type Kubernetes struct {
	Name          string
	Directory     string
	Context       string
	ApiUrl        string
	GuiUrl        string
	BackofficeUrl string
	Protocol      string
}
