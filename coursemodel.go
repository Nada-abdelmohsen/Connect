package models

import (
	"example/data-access/config"
	"example/data-access/entities"
	"html/template"
	"net/http"
	"strconv"
)

type CourseModel struct {
}

func (*CourseModel) FindAll() ([]entities.Course, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("select * from course")
		if err2 != nil {
			return nil, err2
		} else {
			var courses []entities.Course
			for rows.Next() {
				var course entities.Course
				rows.Scan(&course.ID, &course.NAME, &course.CREDITS, &course.DEP)
				courses = append(courses, course)
			}
			return courses, nil

		}
	}
}

func (*CourseModel) Find(id int64) (entities.Course, error) {
	db, err := config.GetDB()
	if err != nil {
		return entities.Course{}, err
	} else {
		rows, err2 := db.Query("select * from course where id = ?", id)
		if err2 != nil {
			return entities.Course{}, err2
		} else {
			var course entities.Course
			for rows.Next() {
				rows.Scan(&course.ID, &course.NAME, &course.CREDITS, &course.DEP)
			}
			return course, nil
		}
	}
}
func (*CourseModel) Create(course *entities.Course) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("insert into course(course_id,course_name,course_cred,course_dept) values(?,?,?,?)", course.ID, course.NAME, course.CREDITES, course.DEP)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}

	}
}
func (*CourseModel) Update(course entities.Course) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("update course set id= ?,name= ?,credits= ?,department= ? where id = ?", course.ID, course.NAME, course.CREDITS, course.DEP)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}

	}
}
func Edit(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var courseModel models.CourseModel
	course, _ := courseModel.Find(id)
	data := map[string]interface{}{
		"course": course,
	}
	tmp, _ := template.ParseFiles("views/course/add.html")
	tmp.Execute(response, data)
}
