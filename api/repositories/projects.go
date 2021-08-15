package repositories

import (
	"context"

	"github.com/Christheoreo/project-manager/dtos"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Implements Interface IProjectsRepository.
type ProjectsRepositoryPostgres struct {
	Pool *pgxpool.Pool
}

func (r *ProjectsRepositoryPostgres) GetByID(ID int) (project dtos.ProjectDto, err error) {
	query := "SELECT p.id, p.title, p.description, pr.name FROM projects p inner join priorities pr on pr.id = p.priority_id where p.id = $1"
	err = r.Pool.QueryRow(context.Background(), query, ID).Scan(&project.ID, &project.Title, &project.Description, &project.Priority)
	if err != nil {
		return
	}

	// Get the tags for the project.
	queryTags := "SELECT t.name FROM tags t inner join project_tags pt on pt.tag_id = t.id where pt.project_id = $1"
	rowTags, errTags := r.Pool.Query(context.Background(), queryTags, project.ID)
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
func (r *ProjectsRepositoryPostgres) GetByUser(user dtos.UserDto) ([]dtos.ProjectDto, error) {
	projects := make([]dtos.ProjectDto, 0)

	query := "SELECT p.id, p.title, p.description, pr.name FROM projects p inner join priorities pr on pr.id = p.priority_id where p.user_id = $1"
	rows, err := r.Pool.Query(context.Background(), query, user.ID)
	if err != nil {
		return projects, err
	}

	for rows.Next() {
		var project dtos.ProjectDto
		err = rows.Scan(&project.ID, &project.Title, &project.Description, &project.Priority)
		if err != nil {
			return projects, err
		}

		// Get the tags for the project.
		queryTags := "SELECT t.name FROM tags t inner join project_tags pt on pt.tag_id = t.id where pt.project_id = $1"
		rowTags, errTags := r.Pool.Query(context.Background(), queryTags, project.ID)
		if errTags != nil {
			return projects, errTags
		}

		project.Tags = make([]string, 0)
		for rowTags.Next() {
			var tag string
			errTags = rowTags.Scan(&tag)

			if errTags != nil {
				return projects, errTags
			}

			project.Tags = append(project.Tags, tag)
		}

		projects = append(projects, project)
	}
	return projects, nil
}

/**
@todo finish adding these.
**/
func (r *ProjectsRepositoryPostgres) TitleTaken(title string, userID int) (bool, error) {
	return
}
func (r *ProjectsRepositoryPostgres) Insert(project dtos.NewProjectDto, user dtos.UserDto) (int, error) {
	return
}
