package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Square of 5: ", square(5))

	//17
	p1 := Person{FirstName: "Mark"}
	person_description(p1)

	//20
	b1 := Book {
		Author: "Mark Twen",
		Title: "Tom Sawyer",
	}
	b2 := Book {
		Author: "Joanne Rowling",
		Title: "Harry Potter",
	}
	b3 := Book {
		Author: "Three Comrades",
		Title: "Erich Remarque",
	}
	l1 := Library {
		Name: "The National Library",
		Books: []Book{b1, b2, b3},
	}
	lib_description(l1)

	//26
	p := Person {
		FirstName: "Stacy",
		Address: Address{Street: "Coleen's 20", City: "New-York"},
	}

	person_address(p)

	//33
	n1 := Node {Value: 1 }
	n2 := Node {Value: 2}
	n1.Next = &n2
	n3 := Node {Value: 3}
	n2.Next = &n3
	n4 := Node {Value: 3}
	n3.Next = &n4
	n5 := Node {Value: 5}
	n4.Next = &n5
	through_node_chain(&n1)
	delete_node(&n1, 3)
	through_node_chain(&n1)
	project := AssignedProject{Tasks: []AssignedTask{AssignedTask{Title: "Go task"}, AssignedTask{Title: "Cpp task"}, AssignedTask{Title: "Marketing task"}} , CompletionPercentage: 10.2345}
	fmt.Println(about_assigned_project(project))

	//44
	e1 := Employee{
		ID: 000001,
		Name: "Josh",
	}
	e2 := Employee{
		ID: 000002,
		Name: "Gemma",
	}
	e1.Next = &e2
	e3 := Employee{
		ID: 000003,
		Name: "Korra",
	}
	e2.Next = &e3
	e4 := Employee{
		ID: 000004,
		Name: "Zukoo",
	}
	e3.Next = &e4
	print_employees(&e1)

	//45
	c1 := Client{
		ID: 000001,
		Name: "Naruto",
	}
	c2 := Client{
		ID: 000002,
		Name: "Albus",
	}
	c1.Next = &c2
	c3 := Client{
		ID: 000003,
		Name: "Kurenai",
	}
	c2.Next = &c3
	c4 := Client{
		ID: 000004,
		Name: "Sasuke",
	}
	add_clinet(&c1, c4)
	print_clients(&c1)


}

//1
type Temperature float64
func weather_forecast(temp Temperature) {
	if temp > 0 {
		fmt.Println("Temperature is higher than zero")
	} else if temp == 0 {
		fmt.Println("Temperature is equal to zero")
	} else {
		fmt.Println("Temperature is lower then zero")
	}
}

//2
type Age int
func is_teen(n Age) string {
	if n >= 13 && n <= 19 {
		return "Teenager"
	} else {
		return "Not a teenager"
	}
}

//3
type Speed float64
func check_speed(s Speed) string {
	if s >= 0 && s <= 120 {
		return "Permissible"
	} else {
		return "Not Permissible"
	}
}

//4
type Score int
func check_score(s Score) string {
	if s > 0 {
		return "Posititve"
	} else if s == 0 {
		return "Zero"
	} else {
		return "Negative"
	}
}

//5
type BMI float64
func identify_bmi(i BMI) string {
	if i < 18.5 {
		return "Underweight"
	} else if i < 25 {
		return "Normal weight"
	} else if i < 30 {
		return "Overweight"
	} else {
		return "Obese"
	}
}

//6
type UnaryOperation func(int) int
var square UnaryOperation = func(x int) int {
	return x * x
}

//7
var double UnaryOperation = func(x int) int {
	return x * 2
}

//8
type Check func(int) bool
var is_even Check = func(x int) bool {
	if x % 2 == 0 {
		return true
	} else {
		return false
	}
}

//9
var is_positive Check = func(x int) bool {
	if x > 0 {
		return true
	} else {
		return false
	}
}

//10
type Operation func(int, int) int
func massive_operation(a, b int, array []Operation) []int{
	result := []int{}
	for _, value := range array {
		result = append(result, value(a, b))
	}
	return result
}

//11
type Count int
func countdown(c Count) {
	for i:= c; i >= 0; i-- {
		fmt.Println(i)
	}
}

//12
type BatteryLevel int
func battery_info(i BatteryLevel) string{
	if i < 15 {
		return "Low battery level"
	} else if i < 60 {
		return "Sufficient battery level"
	} else {
		return "High battery level"
	}
}

//13
type Weight float64
func weight_info(w Weight) string {
	if w < 3 {
		return "Light"
	} else if w < 8 {
		return "Medium"
	} else {
		return "Heavy"
	}
}

//14

type Grade int
func is_satisfying(g Grade) string {
	if g >= 50 {
		return "Satisfying"
	} else {
			return "Not Satisfying"
	}

}

//15
type BloodPressure float64

//16
type Product struct {
	Name string
	Price float64
}
func product_desciption(p Product) {
	fmt.Printf("%+v\n", p)
}

//17
//Person с полями FirstName, LastName и Age
type Person struct {
	FirstName string
	LastName string
	Age
	Address
}
func person_description(p Person) {
	if p.Age == 0 {
		fmt.Printf("Name: %s %s, Age: Not specified\n", p.FirstName, p.LastName)	
	} else {
		fmt.Printf("Name: %s %s, Age: %d\n", p.FirstName, p.LastName, p.Age)
	}
}

//18
func costlier_product(p1, p2 Product) Product{
	if p1.Price >= p2.Price {
		return p1
	} else {
		return p2
	}
}
//19
type Item struct{
	Name string
	Quantity int
}
func total_quantity(item_array []Item) int {
	total := 0
	for _, item := range item_array {
		total += item.Quantity
	}
	return total
}

//20
type Book struct {
	Title string
	Author string
}
type Library  struct {
	Name string
	Books []Book
}
func lib_description(l Library) {
	fmt.Printf("%s has following books:\n", l.Name)
	for i, book := range l.Books {
		fmt.Printf("%d. %s by %s\n", i + 1, book.Title, book.Author)
	}
}

//21
func update_price(productPtr *Product, new_price float64) {
	productPtr.Price = new_price
}

//22
func person_aging(personPtr *Person) {
	personPtr.Age += 1
}

//23
func update_product(productPtr *Product, new_name string, new_price float64) {
	productPtr.Name = new_name
	productPtr.Price = new_price
}

//24
func item_supply(itemPtr *Item, supplied_amount int) {
	itemPtr.Quantity += supplied_amount
}

//25
type Task struct {
	Title string
	Completed bool
}
func completed_task(task_array []Task, task_id int) {
	task_array[task_id].Completed = true
}

//26
type Address struct {
	Street string
	City string
}

func  person_address(p Person) {
	fmt.Printf("Name: %s %s, Address: %+v\n", p.FirstName, p.LastName, p.Address)
}

//27
type Company struct {
	Name string
	Location Address
}

func about_company(c Company) {
	fmt.Printf("Comapany name: %s, Location: %+v\n", c.Name, c.Location)
}

//28
type Course struct{
	Title string
	Instructor Person
}

func about_course(c Course) {
	fmt.Printf("Course title: %s, Instructor: %s %s", c.Title, c.Instructor.FirstName, c.Instructor.LastName)
}

//29
type Event struct {
	Title string
	Location Address
}

func about_event(e Event) {
	fmt.Printf("Event title: %s, Location: %s, %s", e.Title, e.Location.City, e.Location.Street)
}

//30
type Project struct {
	Name string
	Manager Person
}

func about_project(p Project) {
	fmt.Printf("Project name: %s\nProject manager info:\nName: %s %s\nAddress: %s, %s\n", p.Name, p.Manager.FirstName, p.Manager.LastName, p.Manager.Address.City, p.Manager.Address.Street)
}

//31
type Node struct {
	Value int
	Next *Node
}

func two_nodes() {
	n1 := Node{ Value: 1}
	n2 := Node{ Value: 2, Next: &n1}
	n1.Next = &n2
	fmt.Printf("Node1 value: %d\nNext node value: %d", n1.Value, n1.Next.Value)
}

//можно сделать хоть тысячу связанных узлов!
func  create_node_chain(n int) {
	result := []Node{}

	var n1, n2 *Node
	n1 = new(Node)
	for i := 0; i < n; i ++ {
		n2 = new(Node)
		n1.Next = n2
		result = append(result, *n1)
		n1 = n2
	}
	result[len(result)-1].Next = &result[0]

	for _, node := range result {
		fmt.Println("Node1 value:", node, "Next node value:", node.Next)
	}	
}

//32
func three_nodes() {
	n1 := Node{ Value: 1}
	n2 := Node{ Value: 2}
	n3 := Node{ Value: 3, Next: &n1}
	n1.Next = &n2
	n2.Next = &n3
	fmt.Printf("Node1 value: %d\nNode2 value: %d\nNode3 value: %d\n", n1.Value, n1.Next.Value, n1.Next.Next.Value)
}

//33
func through_node_chain(n *Node) {
	fmt.Print("\n", n.Value, " ")
	for n.Next != nil {
		n = n.Next
		fmt.Print(n.Value, " ")
	}
}

//34
func add_node(n *Node, new_n Node) {
	for n.Next != nil {
		n = n.Next
	}
	n.Next = &new_n
}

//35
//работает если есть только 1 нод с удаляемым значением
func delete_node(n *Node, delete_value int) {
	
	if n.Value == delete_value {
		//указателем на первый эл нода станет указатель на 2-ой. Так мы удалим первый нод
		*n = *n.Next
	} else {
		n1 := n
		for n1 != nil && n1.Next != nil {
			if n1.Next.Value == delete_value {
				n1.Next = n1.Next.Next
			}
			n1 = n1.Next
		}
	}
}

//работает для нескольких удаляемых нодов, но если они в середине цепи. Не удаляет >1 нода в начале или на конце
//func delete_node(n *Node, delete_value int) {
//	
//	if n.Value == delete_value {
//		//указателем на первый эл нода станет указатель на 2-ой. Так мы удалим первый нод
//		*n = *n.Next
//	} else {
//		n1 := n
//		for n1 != nil && n1.Next != nil {
//			n2 := n1
//			for n2.Next.Value == delete_value && n2.Next != nil {
//				n2.Next = n2.Next.Next
//			}
//			n1 = n2.Next
//		}
//	}
//}

//func delete_node(nPtr *Node, delete_value int) {
//	n0 := Node{Next: nPtr}
//	n := &n0
//	var cur_n *Node
//	for n.Next != nil {
//		cur_n = n
//		for cur_n.Next != nil {
//			if cur_n.Next.Value == delete_value {
//				cur_n = cur_n.Next
//			} else {
//				break
//			}
//			n.Next = cur_n.Next
//		}
//	}
//}

//36
var p struct {
	Name string
	Price float64
} = struct {
	Name string
	Price float64
}{
	Name: "Lola",
	Price: 60,
}

func about_p(p struct {
	Name string
	Price float64
}) {
	fmt.Printf("The anonymous product %s costs %g\n", p.Name, p.Price)
}

//37
var pers struct {
	FirstName string
	LastName string
	Age Age
} = struct {
	FirstName string
	LastName string
	Age Age
} {
	FirstName: "Nick",
	LastName: "Holding",
	Age: Age(21),
} 

func about_pers(p struct {
	FirstName string
	LastName string
	Age Age
}) {
	fmt.Printf("The anonymous person %s %s is %d years old\n", p.FirstName, p.LastName, p.Age)
}

//38
var car struct {
	Make string
	Model string
	Year int
}
func about_car(c struct {
	Make string
	Model string
	Year int
}) {
	fmt.Printf("The car make is %s, model %s, manufactured in year %d", c.Make, c.Model, c.Year)
}

//39
var a_book struct {
	Title string
	Author Person
}
func about_book(b struct {
	Title string
	Author Person
}) {
	fmt.Printf("The book title is %s, by %s %s", b.Title, b.Author.FirstName, b.Author.LastName)
}

//40
var secret_event struct {
	Title string
	Date time.Time
	Location Address
}

func about_secret_event(event struct {
	Title string
	Date time.Time
	Location Address
}) {
	fmt.Printf("The %s event will take place on %s, %d in %s, %s", event.Title, event.Date.Month(), event.Date.Day(), event.Location.City, event.Location.Street)
}


//EXTRA:

//41
type Student struct {
	ID int
	Name string
	Grades []float64
}

func average_mark(s Student) float64{
	var sum float64 =0
	for _, value := range s.Grades {
		sum += value
	}
	return sum / float64(len(s.Grades))
}

func about_student(s Student) string {
	return fmt.Sprintf("The student name is %s; The average gpa is %g", s.Name, average_mark(s))
}

func best_student(student_slice []Student) Student {
	max_score := average_mark(student_slice[0])
	best := student_slice[0]
	for _, s := range student_slice {
		if average_mark(s) > max_score {
			max_score = average_mark(s)
			best = s
		}
	}
	return best
}

//42
type AssignedTask struct {
	Title string
	Description string
	Status string
}
type AssignedProject struct {
	Name string
	Tasks []AssignedTask
	CompletionPercentage float64
}

func about_assigned_project(p AssignedProject) string{
	result := fmt.Sprintf("The %s includes following tasks: ", p.Name)
	for i, task := range p.Tasks {
		if i != 0 {
			result += ", "
		}
		result += task.Title
	}
	result += fmt.Sprintf("; It's %.2f percent competed ", p.CompletionPercentage)
	return result
}

func closest_completion_project(ps []AssignedProject) AssignedProject {
	max := ps[0].CompletionPercentage
	p := ps[0]
	for _, p0 := range ps {
		if p0.CompletionPercentage > max {
			max = p0.CompletionPercentage
			p = p0
		}
	}
	return p
}

//43 Account с полями AccountNumber и Balance
type Account struct{
	AccountNumber int
	Balance float64
}
type Bank  struct{
	Name string
	Accounts []Account
	TotalAssets float64
}
func about_bank(b Bank) string{
	return fmt.Sprintf("The %s has %d accounts with total assets of %.2f dollars", b.Name, len(b.Accounts), b.TotalAssets)
}

//44
type Employee struct{
	ID int
	Name string
	Next *Employee
}
func print_employees(e *Employee) {
	fmt.Println("ID: ", e.ID, " Name: ", e.Name)
	if e.Next != nil {
		print_employees(e.Next)
	}
}

//45
type Client struct{
	ID int
	Name string
	Next *Client
}
func add_clinet(first_client *Client, new_client Client){
	if first_client.Next != nil {
		add_clinet(first_client.Next, new_client)
	} else {
		first_client.Next = &new_client
	}
}
func print_clients(e *Client) {
	fmt.Println("ID: ", e.ID, " Name: ", e.Name)
	if e.Next != nil {
		print_clients(e.Next)
	}
}

//46
type BookNode struct{
	Title string
	Author string
	Next *BookNode
}
//работает если есть только 1 нод с удаляемым значением
func delete_book_node(n *BookNode, delete_title string) {
	
	if n.Title == delete_title {
		//указателем на первый эл нода станет указатель на 2-ой. Так мы удалим первый нод
		*n = *n.Next
	} else {
		n1 := n
		for n1 != nil && n1.Next != nil {
			if n1.Next.Title == delete_title {
				n1.Next = n1.Next.Next
			}
			n1 = n1.Next
		}
	}
}

//47
type Order struct{
	OrderID int
	Amount int
	Next *Order
}
func total_order_amount(o *Order) {
	total := o.Amount
	for o.Next != nil {
		o = o.Next
		total += o.Amount
	}
}

//48
type TaskNode struct{
	ID int
	Name string
	Next *TaskNode
}
type ProjectNode struct{
	Name string
	Tasks *TaskNode
}
func tasks_number(p *ProjectNode) int {
	t := p.Tasks
	number := 0
	for t!= nil {
		number+=1
		t = t.Next
	}
	return number
}

//49
var anon_event struct{
	EventID int
	Name string
	Location Address
	Date string
}
func about_anon_event(e struct{
	EventID int
	Name string
	Location Address
	Date string
}) string{
	return fmt.Sprintf("The event %d (%s) will take place on %s, %s, %s", e.EventID, e.Name, e.Date, e.Location.City, e.Location.Street)
}

//50
var anon_product struct{
	ProductID int
	Name string
	Category string
	Price float64
}
func about_anon_product(p struct{
	ProductID int
	Name string
	Category string
	Price float64
}) string{
	return fmt.Sprintf("Product %d (%s) belongs to %s Category and costs %g", p.ProductID, p.Name, p.Category, p.Price)
}

//51
var anon_order struct{
	OrderID int
	Product string
	Quantity int
}
func about_anon_order(o struct{
	OrderID int
	Product string
	Quantity int
}) string{
	return fmt.Sprintf("The Order ID%d: Product - %s; Quantity - %d", o.OrderID, o.Product, o.Quantity)
}

//52
var anon_client struct{
	ClientID int
	Name string
	Email string
	Phone string
}

func about_anon_client(c struct{
	ClientID int
	Name string
	Email string
	Phone string
}) string{
	return fmt.Sprintf("Client ID%d (%s) contact information:\nEmail: %s\nPhone: %s", c.ClientID, c.Name, c.Email, c.Phone)
}

//53
var anon_lib struct {
	LibraryID int
	Name string
	Location Address
}
func abou_anon_lib(l struct {
	LibraryID int
	Name string
	Location Address
}) string{
	return fmt.Sprintf("The Library Id%d (%s) is located in %s, %s", l.LibraryID, l.Name, l.Location.City, l.Location.Street)
}


