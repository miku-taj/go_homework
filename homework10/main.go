package main

import (
	"fmt"
	"strings"
)

func main() {
	t := Text {
		Content: "The cardboard box was standing on my table. I peeked inside the box. There lay a few flowers and a postcard.",
	}
	t.ReplaceWord("box", "envelope")
	fmt.Println(t.Content)

}

//1

type Book struct {
	Title  string
	Author string
}
type Library struct{
	Books []Book
}

func (b Book) GetDetails() string {
	return fmt.Sprint("Book title: %s; Book author: %s", b.Title, b.Author)
}
func (l *Library) AddBook(book Book){
	l.Books = append(l.Books, book)
}

//2
type Student struct{
	Grades []int
}
func (s *Student) AddGrade(grade int) {
	s.Grades = append(s.Grades, grade)
}
func (s Student) AverageGrade() float64{
	result := 0
	for _, i := range s.Grades{
		result+= i
	}
	return float64(result)/float64(len(s.Grades))
}

//3 
var PI = 3.14
type Circle struct{
	Radius float64
}
func (c Circle) Area() float64{
	return PI * c.Radius * c.Radius
}
func (c Circle) Circumference() float64{
	return 2 * PI * c.Radius
}

//4
type Item struct{
	Name string
	Price float64
}
type Inventory struct{
	Items []Item
}
func (i *Inventory) AddItem(item Item){
	i.Items = append(i.Items, item)
}
func (i *Inventory) TotalValue() float64{
	value := 0.0
	for _, item := range i.Items{
		value += item.Price
	}
	return value
}

//5
type Transaction struct {
	Amount      float64
	Description string
}
type Account struct{
	Transactions []Transaction
}
func (t Transaction) Process() {
	fmt.Printf("The transaction amount is: %g; Description: %s", t.Amount, t.Description)
}
func (a *Account) AddTransaction(amount float64, description string) {
	t := Transaction{amount,description,}
	a.Transactions = append(a.Transactions, t)
}

//6

type Task struct{
	Title string
	Completed bool
}
type TaskList struct{
	Tasks []Task
}
func (tl *TaskList) AddTask(title string) {
	t:= Task{Title: title, Completed: false,}
	tl.Tasks = append(tl.Tasks, t)
}
func (tl TaskList) CompletedTasks() int{
	q := 0
	for _, task := range tl.Tasks{
		if task.Completed {
			q++
		}
	}
	return q
}

//7
type Temperature struct{
	Celsius float64
}
func (t Temperature) ToFahrenheit() float64 {
	return (t.Celsius * 9 / 5) + 32
}
func (t Temperature) ToKelvin() float64 {
	return t.Celsius +  273.15 
}

//8
type Student1 struct{
	Name string
	Age int
}
type Classroom struct {
	Students []Student1
}
func (c *Classroom) AddStudent(name string, age int) {
	c.Students = append(c.Students, Student1{Name: name, Age: age,})
}
func (c Classroom) AverageAge() float64 {
	total_age := 0
	for _, s := range c.Students{
		total_age += s.Age
	}
	return float64(total_age)/float64(len(c.Students))
}

//9
type Product struct{
	Name string
	Quantity int
}
type Warehouse struct{
	Products []Product
}
func (w *Warehouse) AddProduct(name string, quantity int) {
	w.Products = append(w.Products, Product{Name: name, Quantity: quantity,})
}
func (w Warehouse) TotalQuantity() int{
	total := 0
	for _, p := range w.Products{
		total+= p.Quantity
	}
	return total
}

//10
type Note struct{
	Content string
	Tags []string
}
type Notebook struct{
	Notes []Note
}
func (nb *Notebook) AddNote(content string, tags []string){
	nb.Notes = append(nb.Notes, Note{Content: content, Tags: tags,})
}
func (nb Notebook) NotesWithTag(tag string) []Note {
	var result []Note
	for _, note := range nb.Notes{
		for _, t := range note.Tags{
			if t == tag {
				result = append(result, note)
				break
			}
		}
	}
	return result
}

//11
type Employee struct{
	Name string
	Salary float64
}
type Payroll struct{
	Employees []Employee
}
func (p *Payroll) AddEmployee(name string, salary float64){
	p.Employees = append(p.Employees, Employee{Name: name, Salary: salary,})
}
func (p Payroll) AverageSalary() float64{
	total_salary := 0.0
	for _, emp := range p.Employees{
		total_salary+= emp.Salary
	}
	return total_salary/float64(len(p.Employees))
}

//12
type Budget struct{
	Amount float64
}
func (b *Budget) Allocate(amount float64) bool{
	if amount <= b.Amount {
		b.Amount -= amount
		return true
	}
	return false
}
func (b Budget) Remaining() float64{
	return b.Amount
}

//13
type Order struct{
	ID int
	TotalAmount float64
}
type OrderManager struct{
	Orders []Order
}
func (om *OrderManager) AddOrder(id int, totalAmount float64){
	om.Orders = append(om.Orders, Order{id, totalAmount,})
}
func (om *OrderManager) TotalOrdersAmount() float64{
	total := 0.0
	for _, ord := range om.Orders{
		total += ord.TotalAmount
	}
	return total
}

//14
type Account1 struct{
	Balance float64
}
func (a *Account1) Deposit(amount float64){
	a.Balance += amount
}
func (a *Account1) Withdraw(amount float64) bool {
	if a.Balance >= amount {
		a.Balance -= amount
		return true
	} else {
		return false
	}
}
func (a *Account1) GetBalance() float64{
	return a.Balance
}

//15
type Text struct{
	Content string
}
func (t *Text) AddText(text string){
	t.Content += "\n" + text
}
func (t *Text) ReplaceWord(oldWord, newWord string) {
	split_text := strings.Split(t.Content, " ")
	final_text:= ""
	for i := range split_text{
		if split_text[i] == oldWord {
			split_text[i] = newWord
		}
		final_text+=split_text[i]+" "
	}
	(*t).Content = final_text
}
func (t *Text) Length() int{
	split_text := strings.Split(t.Content, " ")
	return len(split_text)
}

//16
type ShoppingItem struct{
	Name string
	Price float64
}
type ShoppingList struct{
	ShoppingItems []ShoppingItem
}
func (sl *ShoppingList) AddItem(name string, price float64) {
	sl.ShoppingItems = append(sl.ShoppingItems, ShoppingItem{name, price})
}
func (sl ShoppingList) TotalPrice() float64{
	total := 0.0
	for _, item := range sl.ShoppingItems{
		total += item.Price
	}
	return total
}

//17
type Event struct{
	Name string
	Date string
}
type Calendar struct{
	Events []Event
}
func (c *Calendar) AddEvent(name string, date string){
	c.Events = append(c.Events, Event{name, date})
}
/*
func (c *Calendar) EventsAfterDate(date string) []Event{
	result := []Event{}
	d:= strings.Split(date, ".")
	indicatedDay, indicatedMonth, indicatedYear := int(d[0]), int(d[1]), int(d[2])
	var day, month, year int
	for _, event := range c.Events{
		d = strings.Split(event.Date, ".")
		day, month, year := d[0], d[1], d[2]
		if year > indicatedYear {
			result = append(result, event)
		}

	}
}
*/
type Order1 struct{
	OrderID int
	Product string
	Quantity int
}
type Store struct{
	Orders []Order1
}
func (s *Store) AddOrder(orderID int, product string, quantity int){
	s.Orders = append(s.Orders, Order1{orderID, product, quantity,})
}
func (s Store) TotalQuantityByProduct(product string) int{
	total := 0
	for _, ord := range s.Orders{
		if ord.Product == product {
			total += ord.Quantity
		}
	}
	return total
}

//19
type Trip struct{
	Distance float64
	CostPerMile float64
}
func (t *Trip) SetCostPerMile(cost float64){
	t.CostPerMile = cost
}
func (t *Trip) TotalCost() float64{
	return t.CostPerMile * t.Distance
}

//20
type Movie struct{
	Title string
	Rating float64
}
type MovieList struct{
	Movies []Movie
}
func (ml *MovieList) AddMovie(title string, rating float64){
	ml.Movies = append(ml.Movies, Movie{title, rating})
}
func (ml MovieList) AverageRating() float64{
	rating := 0.0
	for _, m := range ml.Movies{
		rating+= m.Rating
	}
	return rating/float64(len(ml.Movies))
}

//21
type Order2 struct{
	ID int
	Product string
	Quantity int
	Price float64
}
type Discount struct{
	Percentage float64
}
func (o *Order2) ApplyDiscount(discount Discount){
	o.Price = o.Price * (100 - discount.Percentage) / 100
}
func (o Order2) TotalAmount() float64{
	return o.Price * float64(o.Quantity)
}
//22
type Account2 struct{
	Balance float64
	History []string
}
func (a *Account2) Deposit(amount float64) {
	a.Balance += amount
}
func (a *Account2) Withdraw(amount float64) bool{
	if a.Balance >= amount {
		a.Balance -=amount
		return true
	} else {
		return false
	}
}
func (a *Account2) TransactionHistory() []string{
	return a.History
}

//23
type Sale struct{
	Item string
	Amount float64
}
type SalesReport struct{
	Sales []Sale
}
func (sr *SalesReport) AddSale(item string, amount float64){
	sr.Sales = append(sr.Sales, Sale{item, amount})
}
func (sr SalesReport) TotalSales() float64{
	total := 0.0
	for _, s := range sr.Sales{
		total += s.Amount
	}
	return total
}

//24
type Transaction1 struct{
	TransactionType string
	Amount float64
}
type Account3 struct{
	Transactions []Transaction1
}
func (a *Account3) AddTransaction(transactionType string, amount float64){
	a.Transactions = append(a.Transactions, Transaction1{transactionType, amount,})

	
}
func (a Account3) Balance() float64{
	result := 0.0
	for _, t := range a.Transactions{
		if t.TransactionType == "debit"{
			result+= t.Amount
		} else {
			result -= t.Amount
		}
	}
	return result
}

//25
type OrderManager2 struct{
	Orders []Order2
}
func (om *OrderManager2) AddOrder(id int, product string, quantity int, price float64){
	om.Orders = append(om.Orders, Order2{id, product, quantity, price})
}
func (om *OrderManager2) ReturnOrder(id int, quantity int){
	for i, o := range om.Orders{
		if o.ID == id {
			if quantity < o.Quantity {
				o.Quantity -= quantity
			} else {
				om.Orders = append(om.Orders[:i], om.Orders[i+1:]...)
			}
		}
	}
}
//!!!!
func (om *OrderManager2) TotalActiveOrders() int{
	return len(om.Orders)
}

//26
type Task2 struct{
	Title string
	Description string
	Status string
}
type Project struct{
	Tasks []Task2
}
func (p *Project) AddTask(title, description, status string){
	p.Tasks = append(p.Tasks, Task2{title, description, status,})
}
func (p Project) CountTasksByStatus(status string) int{
	n := 0
	for _, el := range p.Tasks{
		if el.Status == status{
			n++
		}
	}
	return n
}

//27
type User struct{
	Username string
	Email string
	Age int
}
type UserManager struct{
	Users []User
}
func (u *UserManager) AddUser(username, email string, age int){
	u.Users = append(u.Users, User{username, email, age})
}
func (u UserManager) UsersOlderThan(age int) []User{
	older := []User{}
	for _, user := range u.Users{
		if user.Age > age {
			older = append(older, user)
		}
	}
	return older
}

//28
type Contract struct{
	ContractID int
	Client string
	amount float64
}
type ContractManager struct{
	Contracts []Contract
}
func (cm *ContractManager) AddContract(contractID int, client string, amount float64){
	cm.Contracts = append(cm.Contracts, Contract{contractID, client, amount})
}
func (cm ContractManager) TotalAmountForClient(client string) float64{
	total := 0.0
	for _, c := range cm.Contracts{
		if c.Client == client {
			total += c.amount
		}
	}
	return total
}
//29
type Book2 struct{
	Title string
	Author string
	Year int 
}
type Library2 struct{
	Books []Book2
}
func (l *Library2) AddBook(title, author string, year int){
	l.Books = append(l.Books, Book2{title, author, year})
}
func (l Library2) BooksByAuthorsAfterYear(year int) []Book2{
	books := []Book2{}
	for _, b := range l.Books{
		if b.Year > year{
			books = append(books, b)
		}
	}
	return books
}
//30
type Activity struct{
	ActivityType string
	Timestamp time.Time
}
type UserActivityTracker struct{
	Activities []Activity
}
func (tracker *UserActivityTracker) AddActivity(activityType string, timestamp time.Time){
	tracker.Activities = append(tracker.Activities, Activity{activityType, timestamp})
}
func (tracker UserActivityTracker) ActivitiesAfterTime(timestamp time.Time) []Activity{
	activities := []Activity{}
	for _, a := range tracker.Activities{
		if a.Timestamp.After(timestamp){
			activities = append(activities, a)
		}
	}
	return activities
}

















