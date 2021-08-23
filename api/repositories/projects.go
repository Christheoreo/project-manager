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

// @todo get the components
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
func (r *ProjectsRepositoryPostgres) GetByUserBackup(user dtos.UserDto) ([]dtos.ProjectDto, error) {
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

func (r *ProjectsRepositoryPostgres) GetByUser(user dtos.UserDto) ([]dtos.ProjectDto, error) {
	projects := make([]dtos.ProjectDto, 0)

	query := `SELECT 
		p.id,
		p.title,
		p.description,
		pr.name,
		COALESCE(X.name,'')
	FROM projects p 
		INNER JOIN priorities pr ON p.priority_id = pr.id
		LEFT OUTER JOIN 
	(select * from project_tags pt 
		INNER JOIN tags t ON pt.tag_id = t.id) as X ON p.id = X.project_id
	WHERE p.user_id = $1`
	rows, err := r.Pool.Query(context.Background(), query, user.ID)
	if err != nil {
		return projects, err
	}

	projectsMap := make(map[int]dtos.ProjectDto)

	for rows.Next() {
		var project dtos.ProjectDto
		var tag string

		if err = rows.Scan(&project.ID, &project.Title, &project.Description, &project.Priority, &tag); err != nil {
			return projects, err
		}

		p, ok := projectsMap[project.ID]

		if ok {
			if tag != "" {
				p.Tags = append(p.Tags, tag)
				projectsMap[project.ID] = p
			}
			continue
		}

		project.Tags = make([]string, 0)
		projectsMap[project.ID] = project

		if tag != "" {
			project.Tags = append(project.Tags, tag)
			projectsMap[project.ID] = project
		}
	}
	for _, p := range projectsMap {
		projects = append(projects, p)
	}

	return projects, nil
}

func (r *ProjectsRepositoryPostgres) TitleTaken(title string, userID int) (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM "projects" where "title" = $1 AND user_id = $2`
	err := r.Pool.QueryRow(context.Background(), query, title, userID).Scan(&count)
	if err != nil {
		return true, err
	}
	return count > 0, nil

}
func (r *ProjectsRepositoryPostgres) Insert(project dtos.NewProjectDto, user dtos.UserDto) (int, error) {
	var id int
	query := `INSERT INTO "projects" (user_id, "title", "description", "priority_id") VALUES ($1,$2,$3,$4) RETURNING id`
	err := r.Pool.QueryRow(context.Background(), query, user.ID, project.Title, project.Description, project.PriorityID).Scan(&id)

	if err != nil {
		return -1, err
	}

	for _, tagID := range project.TagIDs {
		var tempID int
		tagsQuery := `INSERT INTO "project_tags" ("project_id", "tag_id") VALUES ($1, $2) RETURNING id`
		err = r.Pool.QueryRow(context.Background(), tagsQuery, id, tagID).Scan(&tempID)

		if err != nil {
			return -1, err
		}
	}

	return id, nil
}

func (r *ProjectsRepositoryPostgres) GetOwnerID(ID int) (int, error) {
	var userID int
	query := "SELECT user_id FROM projects WHERE p.id = $1"
	err := r.Pool.QueryRow(context.Background(), query, ID).Scan(&userID)
	if err != nil {
		return -1, err
	}

	return userID, nil
}
