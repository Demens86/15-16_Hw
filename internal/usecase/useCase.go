package usecase

import (
	"prac/internal/repo"
)

type Case struct {
	Repo *repo.Repo
}

func NewCase(repo *repo.Repo) *Case {
	return &Case{Repo: repo}
}
