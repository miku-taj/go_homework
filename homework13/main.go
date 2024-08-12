package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

//1
type StringProcessor interface {
	Length() int
	WordCount() int
}

type MyString struct {
	Text string
}

func (s MyString) Length() int {
	return len([]rune(s.Text))
}

func (s MyString) WordCount() int {
	words := strings.Split(s.Text, " ")
	return len(words)
}

//2
type Formatter interface{
	ToUpper()
	ToLower()
}
type MyFormatter struct{
	Text string
}
func (f *MyFormatter) ToUpper(){
	f.Text = strings.ToUpper(f.Text)
}
func (f *MyFormatter) ToLower(){
	f.Text = strings.ToLower(f.Text)
}

//3
type PointerOperations interface{
	Increment()
	Decrement()
}
type IntPointer struct{
	Pointer *int
}
func (i IntPointer) Increment(){
	*(i.Pointer) ++
	fmt.Println(*(i.Pointer))
}
func (i IntPointer) Decrement(){
	*(i.Pointer) --
	fmt.Println(*(i.Pointer))
}

//4
type StringCleaner interface{
	TrimSpaces()
	RemoveSpaces()
}
type TextCleaner struct{
	Text string
}
func (t *TextCleaner) TrimSpaces(){
	t.Text = strings.Trim(t.Text, " ")
}
func (t *TextCleaner) RemoveSpaces(){
	t.Text = strings.ReplaceAll(t.Text, " ", "")
}

//5
type StringConcatenator interface{
	Concat() string
	Join(sep string) string
}
type StringJoiner struct{
	StringsSlice []string
}
func (sj *StringJoiner) Concat() string{
	result := ""
	for _, s := range sj.StringsSlice{
		result += s
	}
	return result
}
func (sj *StringJoiner) Join(sep string) string{
	result := ""
	for _, s := range sj.StringsSlice{
		result += s + sep
	}
	return strings.TrimSuffix(result, sep)
}

//6
type SubstringExtractor interface{
	Substr(start, length int) string
	Suffix(length int) string
}
type TextExtractor struct{
	Text string
}
func (te TextExtractor) Substr(start, length int) string{
	if start + length < len([]rune(te.Text)){
		return te.Text[start:start + length]
	} else {
		fmt.Println("Substring can't be extracted; the string is too short")
		return ""
	}
}
func (te TextExtractor) Suffix(length int) string{
	if len(te.Text) - length >= 0{
		return te.Text[len(te.Text) - length:]
	} else {
		fmt.Println("Substring can't be extracted; the string is too short")
		return ""
	}
}

//7
type StringReverser interface{
	Reverse() string
	ReverseWords() string
}
type TextReverser struct{
	Text string
}
func (tr TextReverser) Reverse() string{
	result := ""
	for _, char := range tr.Text{
		result = string(char) + result
	}
	return result
}
func (tr TextReverser) ReverseWords() string{
	words := strings.Split(tr.Text, " ")
	result := ""
	word_result := ""
	for _, word := range words{
		word_result = ""
		for _, char := range word{
			word_result = string(char) + word_result
		}
		result += word_result + " "
	}
	return strings.TrimSuffix(result, " ")
}

//8
type StringInserter interface{
	InsertAt(index int, substring string) string
	InsertAfterWord(word, substring string) string
}
type TextInserter struct{
	Text string
}
func (ti TextInserter) InsertAt(index int, substring string) string{
	if len([]rune(ti.Text)) < index + 1{
		fmt.Println("The text is too short, empty string returned")
		return ""
	} else {
		return ti.Text[:index] + substring + ti.Text[index:]
	}
}
func (ti TextInserter) InsertAfterWord(word, substring string) string{
	words := strings.Split(ti.Text, " ")
	result := ""
	inserted := false
	for _, w := range words {
		result += w + " "
		if w == word{
			result += word + " "
			inserted = true
		}
	}
	if !inserted{
		fmt.Printf("The %s word wasn't found, original text returned\n", word)
	}
	return result
}

//9
type PalindromeChecker interface{
	IsPalindrome() bool
	IsWordPalindrome() bool
}
type TextPalindromeChecker struct{
	Text string
}
func (tp TextPalindromeChecker) IsPalindrome() bool {
	is_p := true
	for i := range tp.Text{
		if tp.Text[i] != tp.Text[len([]rune(tp.Text)) - 1 - i] {
			is_p = false
			break
		}
	}
	return is_p
}
func (tp TextPalindromeChecker) IsWordPalindrome() bool{
	is_p := true
	words := strings.Split(tp.Text, " ")
	for i := range words{
		if words[i] != words[len(words) - 1 - i] {
			is_p = false
			break
		}
	}
	return is_p
}

//10
type DuplicateRemover interface{
	RemoveDuplicates()
	RemoveDuplicatesCaseInsensitive()
}
type TextDuplicateRemover struct{
	Text string
}
func (tr *TextDuplicateRemover) RemoveDuplicates() {
	words := strings.Split(tr.Text, " ")
	result := ""
	for i := range words{
		for j := i + 1; j < len(words); j ++{
			if words[i] == words[j] {
				if j < len(words) - 1{
					words = append(words[:j], words [j+1:]...)
				} else {
					words = words[:j]
				}
			}
		}
	}
	for i := range words{
		result += words[i]+ " "
	}
	tr.Text = strings.TrimSuffix(result, " ")
}
func (tr *TextDuplicateRemover) RemoveDuplicatesCaseInsensitive(){
	Words := strings.Split(tr.Text, " ")
	WordsLower := strings.Split(strings.ToLower(tr.Text), " ")
	result := ""
	for i := range WordsLower {
		for j := i + 1; j < len(WordsLower); j ++{
			if WordsLower[i] == WordsLower[j]{
				if j < len(WordsLower) - 1{
					WordsLower = append(WordsLower[:j], WordsLower [j+1:]...)
					Words = append(Words[:j], Words[j+1:]...)
				} else {
					WordsLower = WordsLower[:j]
					Words = Words[:j]
				}
				
			}
		}
	}
	for i := range Words{
		result += Words[i] + " "
	}
	tr.Text = strings.TrimSuffix(result, " ")
}

//11
type Budget interface{
	AddIncome(amount float64)
	AddExpense(amount float64)
}
type MonthlyBudget struct{
	Budget float64
}
func (b *MonthlyBudget) AddIncome(amount float64){
	b.Budget += amount
}
func (b *MonthlyBudget) AddExpense(amount float64){
	b.Budget -= amount
}
func (b *MonthlyBudget) GetBalance() float64{
	return b.Budget
}

//12
type CurrencyConverter interface{
	ToUSD(amount float64)
	ToEUR(amount float64)
}
type Exchange struct{
	USDExchangeRate float64
	EURExchangeRate float64
}
func (e *Exchange) ToUSD(amount float64) float64{
	return amount * e.USDExchangeRate
}
func (e *Exchange) ToEUR(amount float64) float64{
	return amount * e.EURExchangeRate
}

//13
type BankAccount interface{
	Deposit(amount float64)
	Withdraw(amount float64)
}
type Account struct{
	Balance float64
}
func (a *Account) Deposit(amount float64){
	a.Balance += amount
}
func ( a *Account) Withdraw(amount float64){
	a.Balance -= amount
}
func (a *Account) GetBalance() float64{
	return a.Balance
}

//14
type TaxCalculator interface{
	CalculateIncomeTax(income float64)
	CalculateVAT(amount float64)
}
type SimpleTaxCalculator struct{
	IncomeTaxRate float64
	VATTaxRate float64
}
func (c *SimpleTaxCalculator) CalculateIncomeTax(income float64) float64{
	return income * c.IncomeTaxRate
}
func (c *SimpleTaxCalculator) CalculateVAT(amount float64) float64{
	return amount * c.VATTaxRate
}

//15
type Invoice interface{
	AddItem(name string, price float64)
	Total() float64
}
type SimpleInvoice struct{
	Items map[string]float64
}
func (s *SimpleInvoice) AddItem(name string, price float64){
	if _, exists := s.Items[name]; !exists{
		s.Items[name] = price
	} else {
		fmt.Println("Position already exists")
		//s.Items[name] += price
	}
}
func (s *SimpleInvoice) Total() float64{
	total := 0.0
	for _, val := range s.Items{
		total += val
	}
	return total
}

//16
type StockPortfolio interface{
	Buy(symbol string, quantity int)
	Sell(symbol string, quantity int)
}
type Portfolio struct{
	Shares map[string]int 
}
func (p *Portfolio) Buy(symbol string, quantity int){
	p.Shares[symbol] += quantity
}
func (p *Portfolio) Sell(symbol string, quantity int){
	p.Shares[symbol] -= quantity
}
func (p *Portfolio) GetHoldings() map[string]int{
	return p.Shares
}

//17
type ATM interface{
	Deposit(amount float64)
	Withdraw(amount float64)
}
type ATMSystem struct{
	Balance float64
}
func (s *ATMSystem) Deposit(amount float64){
	s.Balance += amount
}
func (s *ATMSystem) Withdraw(amount float64){
	s.Balance -= amount
}
func (s *ATMSystem) GetBalance() float64{
	return s.Balance
}
//18
type Payroll interface{
	AddEmployee(name string, salary float64)
	GetSalary(name string) float64
}
type CompanyPayroll struct{
	EmployeeSalaries map[string]float64
}
func (p *CompanyPayroll) AddEmployee(name string, salary float64){
	p.EmployeeSalaries[name] = salary
}
func (p *CompanyPayroll) GetSalary(name string) float64{
	return p.EmployeeSalaries[name]
}
func (p *CompanyPayroll) TotalPayroll() float64{
	total := 0.0
	for _, val := range p.EmployeeSalaries{
		total += val
	}
	return total
}
//19
type Library interface{
	AddBook(title, author string)
	GetBooks() []string
}
type BookLibrary struct{
	Books map[string]string
}
func (b *BookLibrary) AddBook(title, author string){
	b.Books[title] = author
}
func (b *BookLibrary) GetBooks() []string{
	titles := []string{}
	for key := range b.Books{
		titles = append(titles, key)
	}
	return titles
}
//20
type OrderManager interface{
	AddOrder(orderID string, amount float64)
	GetTotalSales() float64
}
type OnlineStore struct{
	Orders map[string]float64
}
func (s *OnlineStore) AddOrder(orderID string, amount float64){
	s.Orders[orderID] = amount
}
func (s *OnlineStore) GetTotalSales() float64{
	total := 0.0
	for _, val := range s.Orders{
		total += val
	}
	return total
}
//21
type TransactionManager interface{
	AddTransaction(amount float64)
	GetStatement() []string
}
//type BankAccount struct{}

//22
type ProductCatalog interface{
	AddProduct(name string, price float64)
	GetProductPrice(name string) float64
}
type Catalog struct{
	ProductPrices map[string]float64
}
func (c *Catalog) AddProduct(name string, price float64){
	c.ProductPrices[name] = price
}
func (c *Catalog) GetProductPrice(name string) float64{
	return c.ProductPrices[name]
}
func (c *Catalog) GetProducts() map[string]float64{
	return c.ProductPrices
}
//23
type RestaurantOrderManager interface{
	AddOrder(orderID string, items []string) 
	GetOrderDetails(orderID string) []string
}
type Restaurant struct{
	Orders map[string][]string
}
func (r *Restaurant) AddOrder(orderID string, items []string){
	r.Orders[orderID] = items
}
func (r *Restaurant) GetOrderDetails(orderID string) []string{
	return r.Orders[orderID]
}
func (r *Restaurant) GetAllOrders() map[string][]string{
	return r.Orders
}

//24
type MortgageCalculator interface{
	CalculateMonthlyPayment(principal, annualRate float64, years int)
	CalculateTotalPayment(principal, annualRate float64, years int)
}
type Mortgage struct{

}

func (m *Mortgage) CalculateMonthlyPayment(principal, annualRate float64, years int) float64{
	final_sum := principal * math.Pow((1 + annualRate), float64(years))
	return final_sum / (float64(years) * 12)
}
func (m *Mortgage) CalculateTotalPayment(principal, annualRate float64, years int) float64{
	return principal * math.Pow((1 + annualRate), float64(years))
}

//25
type CustomerManager interface{
	AddCustomer(name string, accountNumber string)
	GetCustomerDetails(accountNumber string) (string, string)
}
type Bank struct{
	Customers map[string]string
}
func (b *Bank) AddCustomer(name string, accountNumber string){
	b.Customers[accountNumber] = name
}
func (b *Bank) GetCustomerDetails(accountNumber string) (string, string){
	return accountNumber, b.Customers[accountNumber]
}
func (b *Bank) GetAllCustomers() map[string]string{
	return b.Customers
}

//26
type TimeTracker interface{
	ClockIn(employeeID string)
	ClockOut(employeeID string)
}
type WorkTime struct{
	Employees map[string][2]time.Time
}

func (w *WorkTime) ClockIn(employeeID string){
	w.Employees[employeeID] = [2]time.Time{time.Now(), time.Now()}
}
func (w *WorkTime) ClockOut(employeeID string) {
	w.Employees[employeeID] = [2]time.Time{w.Employees[employeeID][0], time.Now()}
}
func (w *WorkTime) GetWorkingHours(employeeID string) float64{
	time_dif := w.Employees[employeeID][1].Sub(w.Employees[employeeID][0])
	return time_dif.Hours() + time_dif.Minutes()/60
}

//27
type StudentManager interface{
	AddStudent(name string, studentID string)
	GetStudentDetails(studentID string) (string, string)
}
type University struct{
	Students map[string]string
}
func (u *University) AddStudent(name string, studentID string){
	u.Students[studentID] = name
}
func (u *University) GetStudentDetails(studentID string) (string, string){
	return studentID, u.Students[studentID]
}
func (u *University) GetAllStudents() map[string]string{
	return u.Students
}

//28
type ExpenseManager interface{
	AddExpense(category string, amount float64)
	GetTotalExpenses(category string) float64
}
type Enterprise struct{
	Expenses map[string]float64
}
func (e *Enterprise) AddExpense(category string, amount float64){
	e.Expenses[category] += amount
}
func (e *Enterprise) GetTotalExpenses(category string) float64{
	return e.Expenses[category]
}
func (e *Enterprise) GetAllExpenses() map[string]float64{
	return e.Expenses
}

//29
type CourseManager interface{
	AddCourse(courseName string, credits int)
	GetCourseCredits(courseName string) int
}
type EducationInstitution struct{
	Courses map[string]int
}
func (ei *EducationInstitution) AddCourse(courseName string, credits int){
	ei.Courses[courseName] = credits
}
func (ei *EducationInstitution) GetCourseCredits(courseName string) int{
	return ei.Courses[courseName]
}
func (ei *EducationInstitution) GetAllCourses() map[string]int{
	return ei.Courses
}

//30
type UtilityBillManager interface{
	AddBill(utilityType string, amount float64)
	GetTotalBills(utilityType string) float64
}
type Household struct{
	Bills map[string]float64
}
func (h *Household) AddBill(utilityType string, amount float64){
	h.Bills[utilityType] += amount
}
func (h *Household) GetTotalBills(utilityType string) float64 {
	return h.Bills[utilityType]
}
func (h *Household) GetAllBills() map[string]float64{
	return h.Bills
}




func main() {
	// //1
	// var string_processor StringProcessor
	// string_processor = MyString{
	// 	Text: "this is a text",
	// }
	// fmt.Printf("string_processor (%T, %#v) has %d symbols and %d words\n", string_processor, string_processor, string_processor.Length(), string_processor.WordCount())

	// //3
	// var pointer_op PointerOperations
	// pointer_op = IntPointer{
	// 	Pointer: new(int),
	// }
	// pointer_op.Increment()
	// pointer_op.Increment()
	// pointer_op.Increment()
	// pointer_op.Decrement()

	// //7
	// var str_reverser StringReverser
	// str_reverser = TextReverser{
	// 	Text: "hello friend",
	// }
	// fmt.Printf("Total reverse: %s, Word Reverse: %s\n", str_reverser.Reverse(), str_reverser.ReverseWords())

	//
	// t := TextDuplicateRemover {
	// 	Text: "hey girl Hey i met you girl hEy",
	// }
	// t.RemoveDuplicates()
	// fmt.Println(t.Text)
	// t.RemoveDuplicatesCaseInsensitive()
	// fmt.Println(t.Text)



}

type Cat struct{
	Sound string
}
func (c Cat) MakeSound(){
	fmt.Println(c.Sound)
}
