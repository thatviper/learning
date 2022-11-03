package main

import (
	"context"
	"fmt"
	"log"
	_ "math/rand"
	"net"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	pb "main/pb"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type Employee struct {
	gorm.Model
	Name         string
	Email        string
	Role         string
	Manager      *Employee
	ManagerId    *int32
	DepartmentId int32
}

type Department struct {
	gorm.Model
	Name      string
	Employees []*Employee
}

type EmployeeManagementServer struct {
	pb.UnimplementedEmployeeManagementServer
	db *gorm.DB
}

func (s *EmployeeManagementServer) CreateDepartment(ctx context.Context, in *pb.NewDepartment) (*pb.Department, error) {
	fmt.Println("creating department...")
	//manager := &Employee{Name: in.GetName()}
	//s.db.Debug().Where(manager).Find(manager)
	//fmt.Println(*manager)

	newDept := &Department{
		Model:     gorm.Model{},
		Name:      in.GetName(),
		Employees: []*Employee{},
	}
	savedDept := &Department{}
	s.db.Debug().Save(newDept).Find(savedDept)
	fmt.Println(savedDept)
	out := &pb.Department{
		Name:      in.GetName(),
		Employees: []*pb.Employee{},
	}
	return out, nil
}

func (s *EmployeeManagementServer) CreateEmployee(ctx context.Context, in *pb.Employee) (*pb.Employee, error) {
	fmt.Printf("creating employee... with managerID %v", in.ManagerId)
	var newEmp *Employee
	if in.ManagerId != 0 {
		manager := &Employee{Model: gorm.Model{ID: uint(in.ManagerId)}}
		s.db.Debug().Where(manager).Find(manager)
		//s.db.Debug().Where("id = ?", in.ManagerId).Find(manager)
		fmt.Println(*manager)
		x := in.GetManagerId()
		newEmp = &Employee{
			Name:         in.GetName(),
			Email:        in.GetEmail(),
			Role:         in.GetRole(),
			Manager:      manager,
			ManagerId:    &x,
			DepartmentId: in.DepartmentId,
		}
	} else {
		newEmp = &Employee{
			Name:         in.GetName(),
			Email:        in.GetEmail(),
			Role:         in.GetRole(),
			Manager:      nil,
			ManagerId:    nil,
			DepartmentId: in.DepartmentId,
		}
	}
	result := &Employee{}
	s.db.Debug().Save(newEmp).Find(result)
	fmt.Println(result)
	return in, nil
}

func (s *EmployeeManagementServer) GetEmployeeById(ctx context.Context, in *pb.Id) (*pb.Employee, error) {
	employee := &Employee{Model: gorm.Model{ID: uint(in.GetId())}}
	s.db.Debug().Preload("Manager").Where(employee).Find(employee)
	emp := &pb.Employee{
		Name:         employee.Name,
		Email:        employee.Email,
		Role:         employee.Role,
		ManagerId:    *employee.ManagerId,
		DepartmentId: employee.DepartmentId,
	}
	fmt.Println(employee.Manager)
	return emp, nil
}

func (s *EmployeeManagementServer) GetAllDepartments(ctx context.Context, in *pb.EmptyParams) (*pb.AllDepartments, error) {
	allDepts := &pb.AllDepartments{Departments: []*pb.Department{}}
	departments := []*Department{}
	s.db.Debug().Preload("Employees").Find(&departments)
	for _, dpt := range departments {
		emps := []*pb.Employee{}
		for _, emp := range dpt.Employees {
			e := &pb.Employee{
				Name:         emp.Name,
				Email:        emp.Email,
				Role:         emp.Role,
				ManagerId:    *emp.ManagerId,
				DepartmentId: emp.DepartmentId,
			}
			emps = append(emps, e)
		}
		allDepts.Departments = append(allDepts.Departments, &pb.Department{
			Name:      dpt.Name,
			Employees: emps,
		})
	}
	return allDepts, nil
}

func (s *EmployeeManagementServer) UpdateEmployee(ctx context.Context, in *pb.Employee) (*pb.Employee, error) {
	emp := &Employee{Model: gorm.Model{ID: uint(in.Id)}}
	fmt.Println(in.ManagerId)
	updatedParams := &Employee{}
	if in.DepartmentId != 0 {
		updatedParams.DepartmentId = in.GetDepartmentId()
	}
	if in.Email != "" {
		updatedParams.Email = in.GetEmail()
	}
	if in.ManagerId != 0 {
		x := in.GetManagerId()
		updatedParams.ManagerId = &x
	}
	if in.Name != "" {
		updatedParams.Name = in.GetName()
	}
	if in.Role != "" {
		updatedParams.Role = in.GetRole()
	}
	s.db.Debug().Model(emp).Update(updatedParams).Find(emp)
	fmt.Println(emp)

	return &pb.Employee{
		Name:         emp.Name,
		Email:        emp.Email,
		Role:         emp.Role,
		ManagerId:    *emp.ManagerId,
		DepartmentId: emp.DepartmentId,
	}, nil
}

func (s *EmployeeManagementServer) DeleteEmployeeById(ctx context.Context, in *pb.Id) (*pb.Employee, error) {
	emp := &Employee{Model: gorm.Model{ID: uint(in.Id)}}
	jnr := &Employee{Model: gorm.Model{ID: 2}}
	s.db.Debug().Model(emp).Find(emp).Delete(emp)
	s.db.Debug().Preload("Manager").Model(jnr).Find(jnr)
	fmt.Println(jnr.Manager)
	return &pb.Employee{
		Name:         emp.Name,
		Email:        emp.Email,
		Role:         emp.Role,
		ManagerId:    *emp.ManagerId,
		DepartmentId: emp.DepartmentId,
	}, nil
}

func newServer() *EmployeeManagementServer {
	dbString := fmt.Sprintf("user=postgres password=root dbname=employees-gorm-grpc")
	db, err := gorm.Open("postgres", dbString)
	if err != nil {
		fmt.Printf("error occured %v", err)
	}
	//db.DropTable(&Employee{})
	//db.DropTable(&Department{})
	//db.Debug().AutoMigrate(&Employee{}, &Department{})
	fmt.Println("database connected!")

	s := &EmployeeManagementServer{db: db}
	return s
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := newServer()
	grpcServer := grpc.NewServer()
	pb.RegisterEmployeeManagementServer(grpcServer, s)
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
