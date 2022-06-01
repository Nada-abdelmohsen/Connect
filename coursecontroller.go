package coursecontroller

import (
	"example/data-access/entities"
	"example/data-access/models"
	"html/template"
	"net/http"
	"strconv"
)

func Index(response http.ResponseWriter, request *http.Request) {
	var coursecontroller models.CourseModel
	courses, _ := coursecontroller.FindAll()
	data := map[string]interface{}{
		"courses": courses,
	}
	tmp, _ := template.ParseFiles("views/course/index.html")
	tmp.Execute(response, data)
}

// func Add(response http.ResponseWriter, request *http.Request) {

// 	tmp, _ := template.ParseFiles("views/course/add.html")
// 	tmp.Execute(response, nil)
// }

func Edit(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var courseModel models.CourseModel
	course, _ := courseModel.Find(id)
	data := map[string]interface{}{
		"course": course,
	}
	tmp, _ := template.ParseFiles("views/course/edit.html")
	tmp.Execute(response, data)

}
func Update(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	var course entities.Course
	course.ID, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
	course.NAME = request.Form.Get("name")
	course.CREDITS, _ = strconv.ParseInt(request.Form.Get("credits"), 10, 64)
	course.DEP = request.Form.Get("department")
	var courseModel models.CourseModel
	courseModel.Update(course)
	http.Redirect(response, request, "/course", http.StatusSeeOther)
}
