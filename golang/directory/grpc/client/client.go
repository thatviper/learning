package main

import (
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	pb "main/pb"
	_ "math/rand"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewEmployeeManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// emp1 := &pb.Employee{
	// 	Name:         "Vamshi",
	// 	Email:        "vamshi@beautifulcode.in",
	// 	Role:         "Senior Developer",
	// 	DepartmentId: 1,
	// }
	// emp2 := &pb.Employee{
	// 	Name:         "Naga",
	// 	Email:        "naga@beautifulcode.in",
	// 	Role:         "Junior Developer",
	// 	DepartmentId: 1,
	// 	ManagerId:    1,
	// }

	//dpt1 := &pb.NewDepartment{
	//	Name: "Tester",
	//}

	//r1, err := c.CreateNewDepartment(ctx, dpt1)

	// r, err := c.CreateNewDepartment(ctx, dpt1)
	// if err != nil {
	// 	log.Fatalf("couldnt create deparment: %v", err)
	// }
	// fmt.Println(r.GetName())

	//r1, err := c.CreateNewEmployee(ctx, emp1)
	//r2, err := c.CreateNewEmployee(ctx, emp2)

	//r, err := c.GetEmployeeById(ctx, &pb.Id{Id: 2})

	//r, err := c.GetAllDepartments(ctx, &pb.EmptyParams{})

	r, err := c.UpdateEmployee(ctx, &pb.Employee{Id: 1, Name: "Mani", Email: "mani@beautifulcode.in", Role: "Senior Developer"})

	//r, err := c.DeleteEmployeeById(ctx, &pb.Id{Id: 1})

	if err != nil {
		log.Fatalf("couldnt create deparment: %v", err)
	}
	fmt.Print(r)
}
