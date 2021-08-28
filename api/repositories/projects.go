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

// @todo get the components!
func (r *ProjectsRepositoryPostgres) GetByID(ID int) (models.Project, error) {
	var project models.Project

	//this gets everything in one go, but a lot of data filtering has to happen
	// 	query := `SELECT
	// 	p.id,
	// 	p.title,
	// 	p.description,
	// 	pr.name,
	// 	COALESCE(X.name,''),
	//         COALESCE(c.id,-1)  as componentID,
	//         COALESCE(c.title, '')  as componentTitle,
	//         COALESCE(c.description, '')  as componentDescription,
	//         COALESCE(cd.id, -1)  as cdID,
	//         COALESCE(cd.key, '')  as cdKey,
	//         COALESCE(cd.value, '')  as cdValue

	// FROM projects p
	// 	INNER JOIN priorities pr ON p.priority_id = pr.id

	// 	LEFT OUTER JOIN
	// (select * from project_tags pt
	// 	INNER JOIN tags t ON pt.tag_id = t.id) as X ON p.id = X.project_id
	// 			LEFT OUTER JOIN
	// 				components c  ON c.project_id = p.id
	// LEFT OUTER JOIN component_data cd ON cd.component_id = c.id
	// WHERE p.id = $1`

	query := `SELECT
			p.id,
			p.title,
			p.description,
			pr.name
		FROM projects p
			INNER JOIN priorities pr ON p.priority_id = pr.id
			WHERE p.id = $1`

	err := r.Pool.QueryRow(context.Background(), query, ID).Scan(&project.ID, &project.Title, &project.Description, &project.Priority)
	if err != nil {
		return project, err
	}

	// now get the tags

	queryTags := `SELECT
		t.name
	FROM tags t
	INNER JOIN "project_tags" pt
		on pt.tag_id = t.id
	WHERE pt.project_id = $1`
	rowTags, errTags := r.Pool.Query(context.Background(), queryTags, project.ID)
	if errTags != nil {
		return project, err
	}

	project.Tags = make([]string, 0)
	for rowTags.Next() {
		var tag string
		errTags = rowTags.Scan(&tag)
		if errTags != nil {
			return project, errTags
		}
		project.Tags = append(project.Tags, tag)
	}

	// Now get the components + component Data.

	componentsQuery := `SELECT 
	c.id,
	 c.title,
	 c.description,
	 COALESCE(cd.id, -1)  as cdID,
	 COALESCE(cd.key, '')  as cdKey,
	 COALESCE(cd.value, '')  as cdValue
FROM components c 
LEFT OUTER JOIN component_data cd ON cd.component_id = c.id where c.project_id = $1`

	rowComponents, errComponents := r.Pool.Query(context.Background(), componentsQuery, project.ID)
	if errComponents != nil {
		return project, errComponents
	}

	project.Components = make([]models.ProjectComponent, 0)
	for rowComponents.Next() {
		var (
			id          int
			title       string
			description string
			dataID      int
			dataKey     string
			dataValue   string
		)
		errComponents = rowTags.Scan(&id, &title, &description, &dataID, &dataKey, &dataValue)
		if errComponents != nil {
			return project, errComponents
		}
		c := models.ProjectComponent{
			ID:          id,
			Title:       title,
			Description: description,
			Data:        map[string]string{},
		}
		//test above and add the data for the component
		project.Components = append(project.Components, c)
	}

	//test above
	if err != nil {
		return project, err
	}

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

	// Insert the components

	for _, component := range project.Components {
		var componentID int
		componentQuery := `INSERT INTO "components" ("project_id", "title", "description") VALUES ($1, $2, $3) RETURNING id`
		err = r.Pool.QueryRow(context.Background(), componentQuery, id, component.Title, component.Description).Scan(&componentID)
		if err != nil {
			return -1, err
		}

		for key, value := range component.Data {
			var cdID int
			componentDataQuery := `INSERT INTO "component_data" ("component_id", "key", "value") VALUES ($1,$2,$3) RETURNING id`
			err = r.Pool.QueryRow(context.Background(), componentDataQuery, componentID, key, value).Scan(&cdID)
			if err != nil {
				return -1, err
			}
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
