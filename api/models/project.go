package models

import (
	"context"
	"fmt"
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

// @todo add priority rating
func (p *Project) GetById(tagId int) (project dtos.ProjectDto, err error) {
	query := "SELECT p.id, p.title, p.description, pr.name FROM projects p inner join priorities pr on pr.id = p.priority_id where p.id = $1"
	err = p.Pool.QueryRow(context.Background(), query, tagId).Scan(&project.ID, &project.Title, &project.Description)
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
		err = rows.Scan(&project.ID, &project.Title, &project.Description)
		if err != nil {
			return
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
	sb.WriteString("($1,$2) RETURNING id")
	//
	query := sb.String()
	err = p.Pool.QueryRow(context.Background(), query, user.ID, project.Title, project.Description, project.PriorityID).Scan(&id)

	if err != nil {
		return
	}

	var sbTwo strings.Builder

	sbTwo.WriteString("INSERT INTO \"project_tags\" (\"project_id\", \"tag_id\")")
	sbTwo.WriteString("VALUES")
	values := make([]interface{}, 0)

	for i := 0; i < len(project.TagIDs); i++ {
		var sb strings.Builder

		dollarVar := i + 1

		sb.WriteString(fmt.Sprintf("($%d,$%d)", dollarVar, dollarVar))
		if i == len(project.TagIDs)-1 {
			sb.WriteString(";")
		} else {
			sb.WriteString(",")
		}
		sbTwo.WriteString(sb.String())
		values = append(values, id, project.TagIDs[i])
	}

	query = sb.String()
	err = p.Pool.QueryRow(context.Background(), query, values...).Scan(&id)
	// insert the tag
	return
}
