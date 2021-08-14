package models

import (
	"context"
	"strings"

	"github.com/Christheoreo/project-manager/dtos"
	"github.com/jackc/pgx/v4/pgxpool"
)

type (
	Project struct {
		Pool *pgxpool.Pool
	}

	ProjectComponentToInsert struct {
		ID          int
		Title       string      `json:"title"`
		Description string      `json:"description"`
		Data        interface{} `json:"data"`
	}

	ProjectToInsert struct {
		ID          int
		UserID      int                        `json:"userId"`
		Title       string                     `json:"title"`
		Description string                     `json:"description"`
		Tags        []string                   `json:"tags"`
		Priority    string                     `json:"priority"`
		Components  []ProjectComponentToInsert `json:"components"`
	}
)

func (p *Project) GetById(tagId int) (project dtos.ProjectDto, err error) {
	query := "SELECT p.id, p.title, p.description, pr.name FROM projects p inner join priorities pr on pr.id = p.priority_id where p.id = $1"
	err = p.Pool.QueryRow(context.Background(), query, tagId).Scan(&project.ID, &project.Title, &project.Description, &project.Priority)
	if err != nil {
		return
	}

	// Get the tags for the project.
	queryTags := "SELECT t.name FROM tags t inner join project_tags pt on pt.tag_id = t.id where pt.project_id = $1"
	rowTags, errTags := p.Pool.Query(context.Background(), queryTags, project.ID)
	if errTags != nil {
		return
	}

	project.Tags = make([]string, 0)
	for rowTags.Next() {
		var tag string
		errTags = rowTags.Scan(&tag)

		if errTags != nil {
			return
		}

		project.Tags = append(project.Tags, tag)
	}
	return
}

func (p *Project) GetByUser(user dtos.UserDto) (projects []dtos.ProjectDto, err error) {
	query := "SELECT p.id, p.title, p.description, pr.name FROM projects p inner join priorities pr on pr.id = p.priority_id where p.user_id = $1"
	rows, err := p.Pool.Query(context.Background(), query, user.ID)
	if err != nil {
		return
	}

	projects = make([]dtos.ProjectDto, 0)

	for rows.Next() {
		var project dtos.ProjectDto
		err = rows.Scan(&project.ID, &project.Title, &project.Description, &project.Priority)
		if err != nil {
			return
		}

		// Get the tags for the project.
		queryTags := "SELECT t.name FROM tags t inner join project_tags pt on pt.tag_id = t.id where pt.project_id = $1"
		rowTags, errTags := p.Pool.Query(context.Background(), queryTags, project.ID)
		if errTags != nil {
			return
		}

		project.Tags = make([]string, 0)
		for rowTags.Next() {
			var tag string
			errTags = rowTags.Scan(&tag)

			if errTags != nil {
				return
			}

			project.Tags = append(project.Tags, tag)
		}

		projects = append(projects, project)
	}
	return
}

func (p *Project) HasProjectBeenTakenByUser(title string, userID int) (exists bool, err error) {
	var count int
	query := "SELECT COUNT(*) FROM \"projects\" where \"title\" = $1 AND \"user_id\" = $2"
	err = p.Pool.QueryRow(context.Background(), query, title, userID).Scan(&count)
	if err != nil {
		return
	}

	exists = count > 0
	return
}

func (p *Project) Insert(project dtos.NewProjectDto, user dtos.UserDto) (id int, err error) {
	var sb strings.Builder

	sb.WriteString("INSERT INTO \"projects\" (\"user_id\", \"title\", \"description\", \"priority_id\")")
	sb.WriteString(" VALUES ")
	sb.WriteString("($1,$2,$3,$4) RETURNING id")
	//
	query := sb.String()
	err = p.Pool.QueryRow(context.Background(), query, user.ID, project.Title, project.Description, project.PriorityID).Scan(&id)

	if err != nil {
		return
	}

	for _, tagID := range project.TagIDs {
		var tempID int
		q := "INSERT INTO \"project_tags\" (\"project_id\", \"tag_id\") VALUES ($1, $2) RETURNING id"
		err = p.Pool.QueryRow(context.Background(), q, id, tagID).Scan(&tempID)
	}

	return
}
