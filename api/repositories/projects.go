package repositories

import (
	"context"

	"github.com/Christheoreo/project-manager/models"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Implements Interface IProjectsRepository.
type ProjectsRepositoryPostgres struct {
	Pool *pgxpool.Pool
}

// Helper method for sorting query rows in to projects
func sortRowsInToProjects(rows pgx.Rows) ([]models.Project, error) {
	projects := make([]models.Project, 0)

	projectsMap := make(map[int]models.Project)

	for rows.Next() {
		var project models.Project
		var tag string

		if err := rows.Scan(&project.ID, &project.Title, &project.Description, &project.Priority, &tag); err != nil {
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

		if tag != "" {
			project.Tags = []string{tag}
		} else {
			project.Tags = make([]string, 0)
		}
		projectsMap[project.ID] = project
	}
	for _, p := range projectsMap {
		projects = append(projects, p)
	}

	return projects, nil
}

func (r *ProjectsRepositoryPostgres) GetByID(ID int) (models.Project, error) {
	var project models.Project
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
	WHERE p.id = $1`
	rows, err := r.Pool.Query(context.Background(), query, ID)
	if err != nil {
		return project, err
	}

	projects, err := sortRowsInToProjects(rows)

	if err != nil {
		return project, err
	}

	project = projects[0]

	return project, nil
}

func (r *ProjectsRepositoryPostgres) GetByUser(user models.User) ([]models.Project, error) {
	projects := make([]models.Project, 0)

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
	return sortRowsInToProjects(rows)
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
func (r *ProjectsRepositoryPostgres) Insert(project models.NewProject, user models.User) (int, error) {
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
	query := `SELECT "user_id" FROM "projects" WHERE "id" = $1`
	err := r.Pool.QueryRow(context.Background(), query, ID).Scan(&userID)
	if err != nil {
		return -1, err
	}

	return userID, nil
}
