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
func sortRowsInToProjectsNew(rows pgx.Rows) ([]models.Project, error) {
	projects := make([]models.Project, 0)

	projectsMap := make(map[int]models.Project)

	for rows.Next() {
		var project models.Project
		var tag string

		var component models.ProjectComponent
		var data models.ComponentData

		if err := rows.Scan(&project.ID, &project.Title, &project.Description, &project.Priority, &tag, &component.ID, &component.Title, &component.Description, &data.ID, &data.Key, &data.Value); err != nil {
			return projects, err
		}

		if data.ID != -1 {
			component.Data = []models.ComponentData{data}
		}
		p, ok := projectsMap[project.ID]

		if ok {
			if tag != "" {
				// check if the tag exists
				tagExists := false
				for _, t := range p.Tags {
					if t == tag {
						tagExists = true
						break
					}
				}

				if !tagExists {
					p.Tags = append(p.Tags, tag)
					projectsMap[project.ID] = p
				}
			}

			// Check if the component exists

			if component.ID != -1 {
				componentExists := false

				for i, c := range p.Components {
					if c.ID == component.ID {
						componentExists = true
						// we know it exists, now check if the data exists
						if data.ID != -1 {
							// check if the data exists in the components
							dataExists := false
							for _, y := range c.Data {
								if y.ID == data.ID {
									dataExists = true
									break
								}
							}
							// we only care if it does not exist.
							if !dataExists {
								c.Data = append(c.Data, data)
								p.Components[i] = c
							}
						}
						break
					} else {
						p.Components = append(p.Components, component)
					}
				}

				// update

				if !componentExists {
					// the component doesnt exists, so add the component and any data
					// p.Components = append(p.Components, component)
				}
				projectsMap[project.ID] = p
			}
			continue
		}

		if tag != "" {
			project.Tags = []string{tag}
		} else {
			project.Tags = make([]string, 0)
		}

		if component.ID != -1 {
			project.Components = []models.ProjectComponent{component}
		} else {
			project.Components = make([]models.ProjectComponent, 0)
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

	//this gets everything in one go, but a lot of data filtering has to happen
	query := `SELECT
		p.id,
		p.title,
		p.description,
		pr.name,
		COALESCE(X.name,''),
	        COALESCE(c.id,-1)  as componentID,
	        COALESCE(c.title, '')  as componentTitle,
	        COALESCE(c.description, '')  as componentDescription,
	        COALESCE(cd.id, -1)  as cdID,
	        COALESCE(cd.key, '')  as cdKey,
	        COALESCE(cd.value, '')  as cdValue

	FROM projects p
		INNER JOIN priorities pr ON p.priority_id = pr.id

		LEFT OUTER JOIN
	(select * from project_tags pt
		INNER JOIN tags t ON pt.tag_id = t.id) as X ON p.id = X.project_id
				LEFT OUTER JOIN
					components c  ON c.project_id = p.id
	LEFT OUTER JOIN component_data cd ON cd.component_id = c.id
	WHERE p.id = $1`

	rows, err := r.Pool.Query(context.Background(), query, ID)
	if err != nil {
		return project, err
	}
	projects, err := sortRowsInToProjectsNew(rows)
	if err != nil {
		return project, err
	}
	return projects[0], err
}

func (r *ProjectsRepositoryPostgres) GetByUser(user models.User) ([]models.Project, error) {
	projects := make([]models.Project, 0)
	query := `SELECT
		p.id,
		p.title,
		p.description,
		pr.name,
		COALESCE(X.name,'') as projectName,
	        COALESCE(c.id,-1)  as componentID,
	        COALESCE(c.title, '')  as componentTitle,
	        COALESCE(c.description, '')  as componentDescription,
	        COALESCE(cd.id, -1)  as cdID,
	        COALESCE(cd.key, '')  as cdKey,
	        COALESCE(cd.value, '')  as cdValue

	FROM projects p
		INNER JOIN priorities pr ON p.priority_id = pr.id

		LEFT OUTER JOIN
	(select * from project_tags pt
		INNER JOIN tags t ON pt.tag_id = t.id) as X ON p.id = X.project_id
				LEFT OUTER JOIN
					components c  ON c.project_id = p.id
	LEFT OUTER JOIN component_data cd ON cd.component_id = c.id
	WHERE p.user_id = $1`
	rows, err := r.Pool.Query(context.Background(), query, user.ID)
	if err != nil {
		return projects, err
	}
	return sortRowsInToProjectsNew(rows)
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
