package main

import (
	alert "directory/alerting"
	"flag"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Employee struct {
	gorm.Model
	Name      string
	Email     string
	Role      string
	Manager   *Employee
	ManagerId *uint
	CompanyId uint
}

type Company struct {
	gorm.Model
	Name      string
	Employees []Employee
}

func (emp *Employee) AfterCreate() {
	alert.SendWelcomeEmail(emp.Email)
}

func main() {
	dbname := flag.String("dbname", "", "a string")
	flag.Parse()
	dbString := fmt.Sprintf("user=postgres password=root dbname=%s", *dbname)
	db, err := gorm.Open("postgres", dbString)
	fmt.Println(*dbname)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	db.DropTable(&Employee{})
	db.DropTable(&Company{})
	db.Debug().AutoMigrate(&Employee{}, &Company{})

	fmt.Println("database connected!")
	company := Company{
		Name: "BeautifulCode",
		Employees: []Employee{
			{Name: "Naga Narasimha", Role: "Developer", Email: "naga@beautifulcode.in"},
			{Name: "GP", Role: "Senior Developer", Email: "gp@beautifulcode.in"},
			{Name: "Vamshi", Role: "Senior Developer", Email: "vamshi@beautifulcode.in"}},
	}

	db.Debug().Save(&company)
	m := Employee{Name: "Vamshi"}
	e := Employee{Name: "Naga Narasimha"}
	db.Debug().Where(&m).Find(&m)
	db.Debug().Where(&e).Find(&e)
	e.Manager = &m
	e.ManagerId = &m.ID
	db.Debug().Save(&e)
	db.Debug().Where(&e).Find(&e)
	fmt.Println(*e.ManagerId)
}
