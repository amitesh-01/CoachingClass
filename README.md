# gofr-project

REQUIREMENTS :-
Golang , MySQL , Docker Desktop , Postman(optional)

Open the terminal in project location and run the following command 

    go mod init github.com/example            (here we have to paste a link of a empty repo of users account)
    go get gofr.dev

  run the command in terminal

      go mod tidy

  ** command for connection of mySQL run the following commands
  
    docker run --name gofr-mysql -e MYSQL_ROOT_PASSWORD=root123 -e MYSQL_DATABASE=test_db -p 3307:3306 -d mysql:8.0.30

    docker exec -it gofr-mysql mysql -uroot -proot123 test_db -e "CREATE TABLE Student_details (id INT AUTO_INCREMENT PRIMARY KEY NOT NULL, name VARCHAR(255) NOT NULL ,age INT ,class INT ,teacher_id INT NOT NULL ,fees INT NOT NULL ,address VARCHAR(255) NOT NULL ,mobile_no VARCHAR(255) NOT NULL);"

    docker exec -it gofr-mysql mysql -uroot -proot123 test_db -e "CREATE TABLE Teacher_details (teach_id INT PRIMARY KEY, teach_name VARCHAR(255) NOT NULL ,teach_age INT ,teach_salary INT NOT NULL,teach_num VARCHAR(255) NOT NULL ,teach_add VARCHAR(255) NOT NULL);"

    go run main.go 

  Now To check , use postman (also can be done using chrome browser)

  -to insert data of student ,type the following url using POST method

      localhost:9000/student/Amitesh/18/12/137/1500/bengali Square indore/9770672296
      
      The above url resembles with the one given below
      localhost:9000/student/{name}/{age}/{class}/{teacher_id}/{fees}/{address}/{mobile_no}
    

  -to read data of students ,type the following url using GET method

      localhost:9000/student

  -to update data of a student ,type the following url using PUT method

      localhost:9000/student/1/Amitesh SONI/18/12/137/1500/bengali Square indore/9770672296

      The above url resembles with the one given below
      localhost:9000/student/{id}/{name}/{age}/{class}/{teacher_id}/{fees}/{address}/{mobile_no}

-to delete data of a student ,type the following url using DELETE method

      localhost:9000/student/1

      The above url resembles with the one given below
      localhost:9000/student/{id}

  Similarly, the url's for the database of Teachers

  -to insert data of teacher ,type the following url using POST method

      localhost:9000/teacher/137/Dr Vipin/45/1200000/987654321/abc
      
      The above url resembles with the one given below
      localhost:9000/teacher/{teach_id}/{teach_name}/{teach_age}/{teach_salary}/{teach_num}/{teach_add}
    

  -to read data of students ,type the following url using GET method

      localhost:9000/teacher

  -to update data of a student ,type the following url using PUT method

      ocalhost:9000/teacher/137/Dr Vipin/45/1200000/987654321/xyz

      The above url resembles with the one given below
      localhost:9000/teacher/{teach_id}/{teach_name}/{teach_age}/{teach_salary}/{teach_num}/{teach_add}

-to delete data of a student ,type the following url using DELETE method

      localhost:9000/teacher/137

      The above url resembles with the one given below
      localhost:9000/teacher/{id}

Now , if someone gives teacher's id than we can show the students whom the teacher is teaching by following url

    localhost:9000/teacher/137

    The above url resembles with the one given below
    localhost:9000/teacher/{teach_id}


  
    
    
