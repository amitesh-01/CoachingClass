package main

import "gofr.dev/pkg/gofr"

type Students struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Class      int    `json:"class"`
	Teacher_id int    `json:"teacher_id"`
	Fees       int    `json:"fees"`
	Address    string `json:"address"`
	Mobile_num int    `json:"mobile_no"`
}

type Teachers struct {
	Teach_id     int    `json:"teach_id"`
	Teach_name   string `json:"teach_name"`
	Teach_age    int    `json:"teach_age"`
	Teach_salary int    `json:"teach_salary"`
	Teach_num    int    `json:"teach_num"`
	Teach_add    string `json:"teach_add"`
}

func main() {
	app := gofr.New()

	// Create , Read , Update , Delete operations on teacher database

	app.GET("/teacher", func(ctx *gofr.Context) (interface{}, error) {
		var Teacher []Teachers

		rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM Teacher_details")
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var Teach Teachers
			if err := rows.Scan(&Teach.Teach_id, &Teach.Teach_name, &Teach.Teach_age, &Teach.Teach_salary,
				&Teach.Teach_num, &Teach.Teach_add); err != nil {
				return nil, err
			}

			Teacher = append(Teacher, Teach)
		}

		return Teacher, nil
	})

	app.POST("/teacher/{teach_id}/{teach_name}/{teach_age}/{teach_salary}/{teach_num}/{teach_add}", func(ctx *gofr.Context) (interface{}, error) {
		Teach_id := ctx.PathParam("teach_id")
		Teach_name := ctx.PathParam("teach_name")
		Teach_age := ctx.PathParam("teach_age")
		Teach_salary := ctx.PathParam("teach_salary")
		Teach_num := ctx.PathParam("teach_num")
		Teach_add := ctx.PathParam("teach_add")

		_, err := ctx.DB().ExecContext(ctx,
			"INSERT INTO Teacher_details (teach_id,teach_name,teach_age,teach_salary,teach_num,teach_add) VALUES (?,?,?,?,?,?)",
			Teach_id, Teach_name, Teach_age, Teach_salary, Teach_num, Teach_add)

		return nil, err
	})

	app.PUT("/teacher/{teach_id}/{teach_name}/{teach_age}/{teach_salary}/{teach_num}/{teach_add}", func(ctx *gofr.Context) (interface{}, error) {
		Teach_id := ctx.PathParam("teach_id")
		Teach_name := ctx.PathParam("teach_name")
		Teach_age := ctx.PathParam("teach_age")
		Teach_salary := ctx.PathParam("teach_salary")
		Teach_num := ctx.PathParam("teach_num")
		Teach_add := ctx.PathParam("teach_add")
		_, err := ctx.DB().ExecContext(ctx,
			"UPDATE Teacher_details SET teach_name=?, teach_age=?,teach_salary=?, teach_num=?, teach_add=? WHERE teach_id=?",
			Teach_name, Teach_age, Teach_salary, Teach_num, Teach_add, Teach_id)

		return nil, err
	})

	app.DELETE("/teacher/{teach_id}", func(ctx *gofr.Context) (interface{}, error) {
		Teach_id := ctx.PathParam("teach_id")

		_, err := ctx.DB().ExecContext(ctx,
			"DELETE FROM Teacher_details WHERE teach_id=?",
			Teach_id)

		return nil, err
	})

	// Create , Read , Update , Delete operations on student database

	app.GET("/student", func(ctx *gofr.Context) (interface{}, error) {
		var Studs []Students

		rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM Student_details")
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var Stud Students
			if err := rows.Scan(&Stud.ID, &Stud.Name, &Stud.Age, &Stud.Class,
				&Stud.Teacher_id, &Stud.Fees, &Stud.Address, &Stud.Mobile_num); err != nil {
				return nil, err
			}

			Studs = append(Studs, Stud)
		}

		return Studs, nil
	})

	app.POST("/student/{name}/{age}/{class}/{teacher_id}/{fees}/{address}/{mobile_no}", func(ctx *gofr.Context) (interface{}, error) {
		name := ctx.PathParam("name")
		age := ctx.PathParam("age")
		class := ctx.PathParam("class")
		teacher_id := ctx.PathParam("teacher_id")
		fees := ctx.PathParam("fees")
		address := ctx.PathParam("address")
		mobile_no := ctx.PathParam("mobile_no")

		_, err := ctx.DB().ExecContext(ctx,
			"INSERT INTO Student_details (name,age,class,teacher_id,fees,address,mobile_no) VALUES (?,?,?,?,?,?,?)",
			name, age, class, teacher_id, fees, address, mobile_no)

		return nil, err
	})

	app.PUT("/student/{id}/{name}/{age}/{class}/{teacher_id}/{fees}/{address}/{mobile_no}", func(ctx *gofr.Context) (interface{}, error) {
		id := ctx.PathParam("id")
		name := ctx.PathParam("name")
		age := ctx.PathParam("age")
		class := ctx.PathParam("class")
		teacher_id := ctx.PathParam("teacher_id")
		fees := ctx.PathParam("fees")
		address := ctx.PathParam("address")
		mobile_no := ctx.PathParam("mobile_no")

		_, err := ctx.DB().ExecContext(ctx,
			"UPDATE Student_details SET name=?, age=?,class=?, teacher_id=?, fees=?, address=? ,mobile_no=? WHERE id=?",
			name, age, class, teacher_id, fees, address, mobile_no, id)

		return nil, err
	})

	app.DELETE("/student/{id}", func(ctx *gofr.Context) (interface{}, error) {
		id := ctx.PathParam("id")

		_, err := ctx.DB().ExecContext(ctx,
			"DELETE FROM Student_details WHERE id=?",
			id)

		return nil, err
	})

	//using the Teacher's id , it retrives the data of all the students that have the same teacher.

	app.GET("/teacher/{teach_id}", func(ctx *gofr.Context) (interface{}, error) {
		var Studs []Students

		Teach_id := ctx.PathParam("teach_id")
		rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM Student_details WHERE teacher_id=?", Teach_id)
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var Stud Students
			if err := rows.Scan(&Stud.ID, &Stud.Name, &Stud.Age, &Stud.Class,
				&Stud.Teacher_id, &Stud.Fees, &Stud.Address, &Stud.Mobile_num); err != nil {
				return nil, err
			}

			Studs = append(Studs, Stud)
		}

		return Studs, nil

	})

	app.Start()
}

// docker exec -it gofr-mysql mysql -uroot -proot123 test_db -e "CREATE TABLE Student_details (id INT AUTO_INCREMENT PRIMARY KEY NOT NULL, name VARCHAR(255) NOT NULL ,age INT ,class INT ,teacher_id INT NOT NULL ,fees INT NOT NULL ,address VARCHAR(255) NOT NULL ,mobile_no VARCHAR(255) NOT NULL);"

//localhost:9000/student/Amitesh/18/12/137/1500/bengali Square indore/9770672296

//docker exec -it gofr-mysql mysql -uroot -proot123 test_db -e "CREATE TABLE Teacher_details (teach_id INT PRIMARY KEY, teach_name VARCHAR(255) NOT NULL ,teach_age INT ,teach_salary INT NOT NULL,teach_num VARCHAR(255) NOT NULL ,teach_add VARCHAR(255) NOT NULL);"
//localhost:9000/teacher/137/Dr Vipin/45/1200000/987654321/abc
