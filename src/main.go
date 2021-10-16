package main

import (
	"fmt"
	"reflect"
	"runtime"
	"runtime/debug"
	"strconv"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
)

const myNumber int = 0

const(
	number_a = iota
	number_b
	number_c
)

const(
	_ = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

const(
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

type Student struct {
	Id int
	FirstName string
	LastName string
	Grades []int
}

type Animal struct{
	Name string
	Age int
}

type Bird struct {
	Animal
	Speed float64
	Area string
}

type TagStruct struct{
	Name string `required max:"100"`
	Family string `ImportantField,AndOtherField`
}

func main() {

	//Constant
	printConstant()
	printFileSizes(4000000000.)
	printRole()

	//Array - Slice
	printArray()
	print2DArray()
	printSlice()
	printDefineOtherSlices()
	printAppendAndMake()
	printRemoveElementInSlice()

	//Maps - Struct
	printMaps()
	printStruct()
	printStrucAsRefByValue()
	printCompositInStruct()
	printUsingTagInStruct()

	//FlowControl : IF - SWITCH
	printIfStatement()
	printSwitchStatement(10)
	printSwitchStatement(2)
	printSwitchStatementTaglessSyntax(12)
	printSwitchStatementUsingFallThrough(15)
	printTypeSwitchStatement()

	//Looping
	printSimpleLooping()
	printRangeArray()
	printRangeAccessOnlyToKeyOrValue()

	//ControlFlow : Defer, Panic and Recover
	printDeferStd()
	printDeferStdShowLIFO()
	printDeferEvaluateFunc()
	printPanic()
	printRecover()
	printRecoverAndPanic()

	//POINTERS
	printPointer()
	printPointerAddress()
	printNilInPointers()

	//Functions
	printGenericSum(1,2,3,4,5,6,7)
	printGoStackVariablesToSharedMemory()
	printNameReturnValueFunc()
	printAnonymousFunction()
	printFuncAsVariable()
	printUsingMethod()
	printUsingMethodAsPointer()

	//Interfaces
	printInterfaces()

	//GoRoutines
	printSimpleGoRoutines() // use go run --race for detect race condition
	printSimpleGoRoutinesV2()
	printSimpleGoRoutinesV3()
	stdThread := runtime.GOMAXPROCS(-1)
	printSimpleGoRoutinesV4(stdThread+100)
	printSimpleGoRoutinesV4(stdThread)
	printMAXThreadAvailable()

	//Channels
	printSimpleChannel()
	printDirectionalChannel()
	printChannelBufferSize()
	printChannelBufferSizeV2()
	printUsingSignalChannel()

	//go.mod tidy
	printUsingDiscordTestLib()
}

func printUsingDiscordTestLib(){
	discord, err := discordgo.New("")
	if err != nil{
		fmt.Println("Could not start session")
	}
	fmt.Println("Session: ", discord)

}

var wgChannel = sync.WaitGroup{}
var wgChannelV2 = sync.WaitGroup{}
var wgChannelV3 = sync.WaitGroup{}
var wgChannelV4 = sync.WaitGroup{}

func printUsingSignalChannel(){
	fmt.Println("============================")
	wgChannelV4.Add(2)
	chMessage := make(chan string)
	chCloseSignal := make(chan struct{})

	loggerMethod := func(){
		isLoopActive := true
		for isLoopActive {
			select {
			case entry := <-chMessage:
				fmt.Println("Thread 1: ", entry)
			case <-chCloseSignal:
				fmt.Println("SIGNAL CLOSED 1")
				isLoopActive = false
			}
		}
		fmt.Println("Logger is shutdown 1")
		wgChannelV4.Done()

	}
	loggerMethodv2 := func(){
		isLoopActive := true
		for isLoopActive {
			select {
			case entry := <-chMessage:
				fmt.Println("Thread 2: ",entry)
			case <-chCloseSignal:
				fmt.Println("SIGNAL CLOSED 2")
				isLoopActive = false
			}
		}
		fmt.Println("Logger is shutdown 2")
		wgChannelV4.Done()

	}

	go loggerMethod()
	go loggerMethodv2()

	for i:= 0;i<100;i++{
		chMessage <- "This is Message (" + strconv.Itoa(i) + ")"
	}

	chCloseSignal <- struct{}{}
	chCloseSignal <- struct{}{}
	wgChannelV4.Wait()

	fmt.Println("APPLICATION CLOSED")
}

func printChannelBufferSizeV2(){
	fmt.Println("============================")
	ch  := make(chan int, 20)
	wgChannelV3.Add(2)

	go func(ch <- chan int){	//Receive Method
		for {
			if i, ok := <- ch; ok{
				fmt.Printf("YOUR NUMBER IS %v \n", i)
			}else{
				break
			}
		}
		wgChannelV3.Done()
	}(ch)
	go func(ch chan <- int){  //Send Method
		defer close(ch)
		for j:=0;j<10;j++{
			ch <- (4200 + j)
		}
		wgChannelV3.Done()
	}(ch)

	wgChannelV3.Wait()
}


func printChannelBufferSize(){
	fmt.Println("============================")
	ch  := make(chan int, 20)
	wgChannelV3.Add(2)

	go func(ch <- chan int){	//Receive Method
		for i := range ch{
			fmt.Printf("YOUR NUMBER IS %v \n", i)
		}
		wgChannelV3.Done()
	}(ch)
	go func(ch chan <- int){  //Send Method
		defer close(ch)
		for j:=0;j<10;j++{
			ch <- (420 + j)
		}
		wgChannelV3.Done()
	}(ch)

	wgChannelV3.Wait()
}

func printDirectionalChannel(){
	fmt.Println("============================")
	ch  := make(chan int)
	wgChannelV2.Add(2)

	go func(ch <- chan int){	//Receive Method
		for j:=0;j<10;j++{
			i := <- ch
			fmt.Printf("YOUR NUMBER IS %v \n", i)
		}
		wgChannelV2.Done()
	}(ch)
	go func(ch chan <- int){  //Send Method
		for j:=0;j<10;j++{
			ch <- (42 + j)

		}
		wgChannelV2.Done()
	}(ch)

	wgChannelV2.Wait()
}

func printSimpleChannel(){
	fmt.Println("============================")
	ch  := make(chan int)
	wgChannel.Add(2)

	go func(){	//Receive Method
		for j:=0;j<10;j++{
			i := <- ch
			ch <- 0
			fmt.Printf("YOUR NUMBER IS %v \n", i)
		}
		wgChannel.Done()
	}()
	go func(){  //Send Method
		for j:=0;j<10;j++{
			ch <- (42 + j)
			result := <- ch
			fmt.Println("Consumer Response: ", result)
		}
		wgChannel.Done()
	}()

	wgChannel.Wait()
}

var msgV1 string = "Hello Goroutines"
var wgV2 = sync.WaitGroup{}
var wgV3 = sync.WaitGroup{}
var wgV4 = sync.WaitGroup{}
var counterV4 int = 0
var mutexV4 = sync.RWMutex{}

func printMAXThreadAvailable(){
	fmt.Println("============================")
	fmt.Printf("MAX Thread Available: %v \n", runtime.GOMAXPROCS(-1))
}

func printSimpleGoRoutinesV4(maxThread int){
	counterV4 = 0
	runtime.GOMAXPROCS(maxThread)
	fmt.Println("============================")
	msg := "Hello Goroutines v4"


	for i:=0;i<10;i++{
		msgThread := msg + "Thread(" + strconv.Itoa(i) + ") and counter="
		wgV4.Add(2)
		go _insideGoroutinesPrintHellV4(msgThread)
		go _insideIncrementCountV4()
	}

	msg = "Goodbye"
	wgV4.Wait()
	fmt.Println("Application Finished: msg =",msg)
}

func _insideGoroutinesPrintHellV4(msg string){
	mutexV4.RLock()
	fmt.Println(msg, counterV4)
	mutexV4.RUnlock()
	wgV4.Done()
}
func _insideIncrementCountV4(){
	mutexV4.Lock()
	counterV4++
	mutexV4.Unlock()
	wgV4.Done()
}


func printSimpleGoRoutinesV3(){
	fmt.Println("============================")
	msg := "Hello Goroutines v3"

	for i:=0;i<10;i++{
		msgThread := msg + "Thread(" + strconv.Itoa(i) + ")\n"
		wgV3.Add(1)
		go func(msg string){
			fmt.Println(msg)
			wgV3.Done()
		}(msgThread)
	}

	msg = "Goodbye"
	wgV3.Wait()
	fmt.Println("Application Finished: msg =",msg)
}


func printSimpleGoRoutinesV2(){
	fmt.Println("============================")
	msg := "Hello Goroutines v2"
	wgV2.Add(1)
	go func(msg string){
		fmt.Println(msg)
		wgV2.Done()
	}(msg)
	msg = "Goodbye"
	wgV2.Wait()
	fmt.Println("Application Finished: msg =",msg)
}

func printSimpleGoRoutines(){
	fmt.Println("============================")

	go _insiedPrintSimpleGoRoutines()
	msgV1 = "Goodbye Goroutines" //race condition ==> using go run --race .
	time.Sleep(100 * time.Microsecond)
}

func _insiedPrintSimpleGoRoutines(){
	fmt.Println(msgV1)
}

func printInterfaces(){
	fmt.Println("============================")
	var w Writer = ConsoleWriter{}
	w.Write([]byte("HELLO ALI MOHIT"))
}

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct {}

func (cw ConsoleWriter) Write(data []byte) (int, error){
	n,err := fmt.Println(string(data))

	return n,err
}

func printUsingMethodAsPointer(){
	fmt.Println("============================")
	p := NumberHolder{}

	for i:=0;i<10;i++{
		fmt.Println("P.Number is Change: ", p.Number)
		p.PlusPlus()
	}
}

type NumberHolder struct {
	Number int
}

func (g *NumberHolder) PlusPlus(){
	(*g).Number++
}

func printUsingMethod(){
	fmt.Println("============================")
	book := Book{
		Name: "C++ Programming",
		Author: "Daitel by Daitel",
		ISBN: "121313FA1231234",
	}

	fmt.Println("Using ToString() method: ",book.ToString())
}

type Book struct{
	Name string
	Author string
	ISBN string
}

func (g Book) ToString() string {
	return fmt.Sprintf("Name: %v - Author: %v - ISBN: %v", g.Name, g.Author, g.ISBN)
}

func printFuncAsVariable(){
	fmt.Println("============================")
	var funcVariable func() = func(){
		fmt.Println("THIS IS VARIABLES METHOD")
	}

	funcVariable()
}


func printAnonymousFunction(){
	fmt.Println("============================")
	func(){
		fmt.Println("Anonymous Function")
	}()
}

func printNameReturnValueFunc() (result int){
	fmt.Println("============================")
	result = 2*2*2
	fmt.Println("NamedReturn Value is: ",result)
	return
}

func printGoStackVariablesToSharedMemory(){
	fmt.Println("============================")
	result := _innerPrintGoStackVariablesToSharedMemory(10,20,30,40,50)
	fmt.Println(result)
	fmt.Println(*result)
}

func _innerPrintGoStackVariablesToSharedMemory(values ...int) *int {
	result := 0
	for _,item := range values{
		result += item
	}
	fmt.Println("RESULT is not in STACK MEMORY:", &result)
	return &result
}

func printGenericSum(values ...int){
	fmt.Println("============================")
	result := 0

	for _,item := range values{
		result += item
	}
	fmt.Println(result)
}


func printNilInPointers(){
	fmt.Println("============================")
	var item *myStruct
	fmt.Println(item)
	item = new(myStruct)
	(*item).Id = 20   // item.Id because of compiler
	fmt.Println(item)
}

type myStruct struct {
	Id int
}

func printPointerAddress(){
	fmt.Println("============================")
	arrayItem := [...]int{1,2,3}
	indexOne := &arrayItem[0]
	indexTwo := &arrayItem[1]
	indexThree := &arrayItem[2]

	fmt.Printf("ARRAY IS %p %v and ADDRESSES: %v %v %v\n",
		&arrayItem, arrayItem, indexOne, indexTwo, indexThree)
}

func printPointer(){
	fmt.Println("============================")
	a := 50
	b := a
	fmt.Println(a, b)
	a = 100
	fmt.Println(a, b)

	// For Define Ptr using * before type
	// For Using Address using & before variable
	var c int = 1000
	var d *int = &c
	fmt.Printf("C = (%v) %v -- D = (%v) %v\n",&c,c, d, *d)
	*d = 1010
	fmt.Printf("C = (%v) %v -- D = (%v) %v\n",&c,c, d, *d)
}

func printRecoverAndPanic(){
	fmt.Println("============================")

	_, err1 := _insidePrintRecoverAndPanic(3)
	if err1 != nil{
		fmt.Println("ERROR RESULT1: ", err1.Error())

	}
	_, err2 := _insidePrintRecoverAndPanic(0)
	if err2 != nil{
		fmt.Println("ERROR RESULT2: ", err2.Error())
	}
	fmt.Println("HANDLING FINISHED")
}

func _insidePrintRecoverAndPanic(input int) (res int,myerr error){
	defer func(){
		if err := recover();err != nil{
			myerr = fmt.Errorf("%v\nSTACK TRACE: %v\n", err, string(debug.Stack()))
		}
	}()
	if input % 2 == 0{
		panic("This is not a odd number!!!!!")
	}
	fmt.Println("INPUT NUMBER IS:", input)
	res = input
	return
}

func printRecover(){
	fmt.Println("============================")
	fmt.Println("I want to call PANIC_METHOD")
	_insidePrintRecover()
	fmt.Println("END OF PANIC_RECOVERY")
}

func _insidePrintRecover(){
	tempFunc := func(){
		err := recover()
		if err != nil{
			fmt.Println("ERROR HANDLING MESSAGE: ",err)
			//debug.PrintStack()
		}
	}
	defer tempFunc()

	fmt.Println("Starting to PANICING...")
	panic("OHHHH MY GODDDDDDDDDDDD")
}

func printPanic(){
	defer func(){
		err := recover()
		if err != nil{
			fmt.Println("PANIC RAISED")
			fmt.Println(err)
		}
	}()
	fmt.Println("============================")
	fmt.Println("BEFORE PANIC Raised...")
	panic("OHHHHHHHHHHHHHH, PANIC RAISED")
	fmt.Println("This line would not printed")
}

func printDeferEvaluateFunc(){
	fmt.Println("============================")
	a := "start"
	fmt.Println("Must Be Print (start) == ",a)
	a = "end"
}

func printDeferStdShowLIFO(){
	fmt.Println("============================")
	defer fmt.Println("START")
	defer fmt.Println("MIDDLE")
	defer fmt.Println("END")
}

func printDeferStd(){
	fmt.Println("============================")
	fmt.Println("START")
	fmt.Println("MIDDLE")
	fmt.Println("END")

	fmt.Println("+++++++++++ NOW USING DEFER +++++++++")
	fmt.Println("START")
	defer fmt.Println("DEFER MIDDLE")
	fmt.Println("END")
}

func printRangeAccessOnlyToKeyOrValue(){
	fmt.Println("============================")
	grades := map[string]int{
		"Ali": 20,
		"Mohammad": 20,
		"Fatemeh": 19,
		"AhmadReza": 15,
	}

	fmt.Print("KEYS: ")
	for key := range grades{
		fmt.Printf("%v, ", key)
	}
	fmt.Println()

	fmt.Print("VALUES: ")
	for _, value := range grades{
		fmt.Printf("%v, ", value)
	}
	fmt.Println()
}

func printRangeArray(){
	fmt.Println("============================")
	numberArray := [...]int{20,30,40}
	city := map[string]int{
		"Tehran": 14000000,
		"Esfehan": 7000000,
	}
	for k,v := range numberArray{
		fmt.Printf("Index , Value: %v, %v \n", k,v)
	}

	for key,value := range city{
		fmt.Println(key , value)
	}

	var message string = "HI My Name is Ali Mohit!"

	for index, character := range message{
		fmt.Println(index, string(character))
	}

}

func printSimpleLooping(){
	fmt.Println("============================")
	fmt.Print("Your Numbers is : ")
	for i := 0; i<10; i++{
		if i != 0{
			fmt.Print(", ")
		}
		fmt.Printf("%v", i)
	}
	fmt.Println("")

	fmt.Print("Tuple : ")
	for i,j := 0, 0; i<10; i,j = i+1, (j+i)*2{
		if i != 0{
			fmt.Print(", ")
		}
		fmt.Printf("(%v, %v)", i, j)
	}
	fmt.Println("")

	var isActive bool = true
	var number int = 0
	for isActive{
		if number > 100{
			isActive = false
		}

		number++
		if number % 2 ==0{
			continue
		}
		if number != 0{
			fmt.Print(", ")
		}
		fmt.Printf("%v", number)
	}
	fmt.Println("")
}

func printTypeSwitchStatement(){
	fmt.Println("============================")
	var item interface{} = 1

	switch item.(type) {
	case int:
		fmt.Println("This is a integer number")
		break
		fmt.Println("Must not be execute")
	case float64:
		println("This is a float64 number")
	case float32:
		println("This is a float32 number")
	case string:
		println("This is a Text")
	default:
		println("Can not find type of object")
	}
}

func printSwitchStatementUsingFallThrough(number int) {
	fmt.Println("============================")
	switch {
	case number <= 10:
		fmt.Println("Less than equal ten")
	case number <= 20:
		fmt.Println("Less than equal twenty")
		//fallthrough
	case number <= 30:
		fmt.Println("Less than equal thirty")
	default:
		fmt.Println("Bigger than twenty!")
	}
}

func printSwitchStatementTaglessSyntax(number int) {
	fmt.Println("============================")
	switch {
	case number <= 10:
		fmt.Println("Less than equal ten")
	case number <= 20:
		fmt.Println("Less than equal twenty")
	default:
		fmt.Println("Bigger than twenty!")
	}
}

func printSwitchStatement(number int){
	fmt.Println("============================")

	switch number {
	case 1:
		fmt.Println("This is one!")
	case 2:
		fmt.Println("This is Two!")
	case 10,20,30:
		fmt.Println("This Ten, Twenty or Thirty")
	default:
		fmt.Println("not one or two!")
	}
}

func printIfStatement(){
	fmt.Println("============================")

	if true{
		fmt.Println("The test is true")
	}

	cityPopulation := map[string]int {
		"Tehran": 14000000,
		"Esfehan": 7000000,
		"Tabriz": 4000000,
	}
	if population, ok := cityPopulation["Tehran"]; ok{
		fmt.Printf("Tehran Populations: %v \n", population)
	}

}

func printUsingTagInStruct(){
	fmt.Println("============================")

	t := reflect.TypeOf(TagStruct{})
	fieldName,_ := t.FieldByName("Name")
	fieldFamily,_ := t.FieldByName("Family")

	fmt.Printf("Field Name's Tag: %v \n", fieldName.Tag)
	fmt.Printf("Field Family's Tag: %v \n", fieldFamily.Tag)
}

func printCompositInStruct(){
	fmt.Println("============================")

	birdData := Bird{}
	birdData.Name = "COCO"
	birdData.Age = 2
	birdData.Speed = 20
	birdData.Area = "Africa"

	birdItem := Bird{
		Animal: Animal{
			Name: "Eagle",
			Age: 1,
		},
		Speed: 10,
		Area: "Asia",
	}

	fmt.Printf("My BirdData: %v \n", birdData)
	fmt.Printf("My EagleData: %v \n", birdItem)
}

func printStrucAsRefByValue(){
	fmt.Println("============================")
	mainStudent := Student{
		Id: 20,
		FirstName: "Ali",
		LastName: "Mohit",
		Grades: []int{100,100},
	}

	secondStudent := mainStudent
	secondStudent.FirstName = "Mohammad"

	sameStuden := &mainStudent

	fmt.Printf("First Student: %v \n", mainStudent)
	fmt.Printf("Second Student: %v \n", secondStudent)


	sameStuden.FirstName = "Niloufar"
	fmt.Printf("After Change Main Student: %v \n", mainStudent)
	fmt.Printf("Ref to Main Student: %v \n", sameStuden)

}

func printStruct(){
	fmt.Println("============================")

	aliMohit := Student{
		Id: 10,
		FirstName: "Ali",
		LastName: "Mohit",
		Grades: []int{20,20,20,19},
	}

	//Positional Syntax
	mohammadMohit := Student{
		11,
		"Mohammad",
		"Mohit",
		[]int{19,19,20,20},
	}

	fmt.Printf("Student: %v \n", aliMohit)
	fmt.Printf("Student: %v \n", mohammadMohit)
	fmt.Printf("Type Student: %T \n", aliMohit)
}

func printMaps(){
	fmt.Println("============================")

	cityPopulation := make(map[string]int)
	cityPopulation = map[string]int{
		"Tehran": 14000000,
		"Esfehan": 7000000,
		"Hamedan": 2000000,
	}
	cityPopulation["Saveh"] = 1000000
	fmt.Printf("Size of Map: %v \n", len(cityPopulation))
	fmt.Printf("City Population: %v \n", cityPopulation)

	fmt.Printf("Tehran Population: %v \n", cityPopulation["Tehran"])
	fmt.Printf("Tabriz Population: %v \n", cityPopulation["Tabriz"])

	delete(cityPopulation, "Saveh")
	delete(cityPopulation, "Tabriz")
	fmt.Printf("After DELETE: %v \n", cityPopulation)

	//If Data Not Exists
	noExistCity, ok := cityPopulation["Ghom"]
	esfehanCity, esfehan_ok := cityPopulation["Esfehan"]
	fmt.Printf("GHOM: %v - %v\n", noExistCity, ok)
	fmt.Printf("GHOM: %v - %v\n", esfehanCity, esfehan_ok)


	//Map is Pointer
	sp := cityPopulation
	delete(sp, "Hamedan")
	fmt.Printf("After Assign New Ptr: %v \n", sp)
	fmt.Printf("After Assign Old Ptr: %v \n", cityPopulation)


}

func printRemoveElementInSlice(){
	fmt.Println("============================")

	a := []int{1,2,3,4,5,6}
	removeFirstElement := a[1:]
	removeLastElement := a[:len(a)-1]
	removeMiddleElement := append(a[:2], a[3:]...)

	fmt.Printf("Main Array: %v \n", a)
	fmt.Printf("Removed FirstElement: %v \n", removeFirstElement)
	fmt.Printf("Removed LastElement: %v \n", removeLastElement)
	fmt.Printf("Removed MiddleElement: %v \n", removeMiddleElement)

}

func printAppendAndMake(){
	a := []int{}

	fmt.Println("============================")
	fmt.Printf("A: %v \n", a)
	fmt.Printf("A's length: %v \n", len(a))
	fmt.Printf("A's capacity: %v \n", cap(a))

	a = append(a, 1)
	fmt.Printf("A: %v \n", a)
	fmt.Printf("A's length: %v \n", len(a))
	fmt.Printf("A's capacity: %v \n", cap(a))

	b := make([]int, 0, 10)
	b = append(b, 10,20,30)

	//spread the array values
	b = append(b, []int{100,200,300}...)

	fmt.Printf("B: %v \n", b)
	fmt.Printf("B's length: %v \n", len(b))
	fmt.Printf("B's capacity: %v \n", cap(b))

}

func printDefineOtherSlices(){
	fmt.Println("============================")

	mainArray := []int{0,1,2,3,4,5,6,7,8}
	a := mainArray
	b := mainArray[:]
	c := mainArray[3:]
	d := mainArray[3:6]

	fmt.Printf("MainArray: %v \n", mainArray)
	fmt.Printf("A: %v \n", a)
	fmt.Printf("B: %v \n", b)
	fmt.Printf("C: %v \n", c)
	fmt.Printf("D: %v \n", d)

}

func printSlice(){
	a := []int{1,2,3}
	b := a

	fmt.Println("============================")
	fmt.Printf("1st array: %v \n", a)
	fmt.Printf("array's length: %v \n", len(a))
	fmt.Printf("array's capacity: %v \n", cap(a))

	b[1] = 5
	fmt.Printf("B is refrenced type of A: %v \n", b)
	fmt.Printf("A value is: %v \n", a)
}

func print2DArray(){
	fmt.Println("============================")
	var identifyMatrix [3][3]int
	identifyMatrix[0] = [3]int {1,0,0}
	identifyMatrix[1] = [3]int {0,1,0}
	identifyMatrix[2] = [3]int {0,0,1}

	secondMatrix := identifyMatrix
	secondMatrix[0][2] = 10

	thirdMatrix := &identifyMatrix
	thirdMatrix[0][1] = 5

	fmt.Printf("Identify Matrix: %v \n", identifyMatrix)
	fmt.Printf("Second Matrix: %v \n", secondMatrix)
	fmt.Printf("Third matrix must be equal identifyMatrix: %v \n", thirdMatrix)

}

func printArray(){
	fmt.Println("============================")
	grades := [3]int{100,90,80}
	names := [...]string{"Ali", "Mohammad", "Niloufar"}
	var families [3]string

	fmt.Printf("Grades: %v \n", grades)
	fmt.Printf("Names: %v \n", names)

	families[0] = "Mohit"
	families[1] = "Mohit"
	families[2] = "Mardiha"
	fmt.Printf("Families: %v \n", families)
	fmt.Printf("Number of Students: %v \n", len(names))

}

func printConstant(){
	fmt.Println("============================")

	fmt.Printf("Value: %v, Type: %T\n", myNumber, myNumber)
	fmt.Printf("Value: %v, Type: %T\n", number_a, number_a)
	fmt.Printf("Value: %v, Type: %T\n", number_b, number_b)
	fmt.Printf("Value: %v, Type: %T\n", number_c, number_c)
}


func printFileSizes(file_size float32){
	fmt.Println("============================")
	fmt.Printf("File Size: %.2fGB \n", file_size/GB)
}

func printRole(){
	var roles byte = isAdmin | canSeeAfrica | canSeeEurope
	fmt.Println("============================")
	fmt.Printf("roles: %b\n", roles)
	fmt.Printf("Is Admin? \t%v \n", roles & isAdmin == isAdmin)
	fmt.Printf("Is HeadQuarters? \t%v \n", roles & isHeadquarters == isHeadquarters)

}