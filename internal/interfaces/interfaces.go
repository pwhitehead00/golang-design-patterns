package interfaces

import "fmt"

type project interface {
	GetRepo(r string) string
}

type gitlab struct {
	client string
}

func (gl *gitlab) GetRepo(r string) string {
	fmt.Println("client:", gl.client)
	return fmt.Sprintf("repo: %s", r)
}

type github struct {
	client string
}

func (gh *github) GetRepo(r string) string {
	fmt.Println("client:", gh.client)
	return fmt.Sprintf("repo: %s", r)
}

func NewProject(t string) project {
	switch t {
	case "gitlab":
		return &gitlab{
			client: "gitlab client",
		}
	case "github":
		return &github{
			client: "github client",
		}
	default:
		return nil
	}
}
