package main
// import ("fmt"; "math")
// import "fmt";
import (
    "fmt"
    // "time"
    // "math"
    // "math/rand"
    // "strings"
    // "oi"
    "os"
    // "io/ioutil"
    // "path/filepath"
    // "errors"
// "sort"
// "hash/crc32"
// "net/http"
)


// import "io"

// import "os"

func main() {
    fmt.Println("Hello, world!")
//      x  :=18+1;
//     fmt.Println(x);
//  q  := "hello";
//   z := "world";


//  fmt.Println(q==z);
//  fmt.Println("ojieh is ",x);
    
// var (a=3
//     b=4
//     c=5
// )



// fmt.Println(a,b,c)

// fmt.Print("Enter a number: ")
// var input float64
// fmt.Scanf("%f", &input)
// output := input * 2
// fmt.Println(output)
//






//for loop

// i :=1;

// for i <=10{
//     fmt.Println(i);
//     i=i+1;

// }




// for i :=1; i<=10; i++{
//     fmt.Println(i);
// }



//if statement


// for i:=1; i<=10; i++{
//     if i%2==0{
//         fmt.Println(i,"even");
//     }else{
//         fmt.Println(i,"odd");
//     }
// }



//Switch statement
// i:=1;
// switch i {
// case 0:fmt.Println("zero")
// case 1:fmt.Println("one")
// case 2:fmt.Println("two")
// case 3:fmt.Println("three")
// case 4:fmt.Println("four")
// case 5:fmt.Println("five")
// case 6:fmt.Println("six")
// case 7:fmt.Println("seven")
// case 8:fmt.Println("eight")
// case 9:fmt.Println("nine")
    
// }


// for i:=1; i<=100; i++{
//     if i%3==0{
//         fmt.Println(i)
//     }
// }


//Array slices and maps

// var x [5] int;
// x[4]=100;
// fmt.Println(x);

// var x [5]float64
// x[0] = 98
// x[1] = 93
// x[2] = 77
// x[3] = 82
// x[4] = 83

// var total float64=0;


// for i:=0; i<len(x); i++{
//     total+=x[i];
// }

// fmt.Println(total / float64(len(x)));

//Slice
//A slice is a segment of an array. Like arrays slices are
// indexable and have a length. Unlike arrays this length
// is allowed to change

 
//normal method
// x := make([]float64,5,10);

// x[0] = 98
// x[1] = 93
// x[2] = 77
// x[3] = 82
// x[4] = 83


// fmt.Println(x)


// arr := [5]float64{1,2,3,4,5}


// x := arr[0:5]

// fmt.Println(x);

//append 
// slice1 := []int{1,2,3}
// slice2 := append(slice1, 4, 5)
// fmt.Println(slice1, slice2)

//Copy

// slice1 := []int{1,2,3,6,8,9}
// slice2 := make([]int, 8)
// copy(slice2, slice1)
// fmt.Println(slice1, slice2)



//MAPPING


//var x map[string]int;
// x:= make(map[string]int)
// x["key"] = 10
// x["world"] = 11

// y:= make(map[int]int);
// y[1]=100;

// fmt.Println(y);


// delete(x, "key");
// fmt.Println(x)

elements := make(map[string]string)
elements["H"] = "Hydrogen"
elements["He"] = "Helium"
elements["Li"] = "Lithium"
elements["Be"] = "Beryllium"
elements["B"] = "Boron"
elements["C"] = "Carbon"
elements["N"] = "Nitrogen"
elements["O"] = "Oxygen"
elements["F"] = "Fluorine"
elements["Ne"] = "Neon"



name, ok := elements["H"]
fmt.Println(name, ok)


// elements := map[string]map[string]string{
//     "H": map[string]string{
//     "name":"Hydrogen",
//     "state":"gas",
//     },
//     "He": map[string]string{
//     "name":"Helium",
//     "state":"gas",
//     },
//     "Li": map[string]string{
//     "name":"Lithium",
//     "state":"solid",
//     },"Be": map[string]string{
// "name":"Beryllium",
// "state":"solid",
// },
// "B": map[string]string{
// "name":"Boron",
// "state":"solid",
// },
// "C": map[string]string{
// "name":"Carbon",
// "state":"solid",
// },
// "N": map[string]string{
// "name":"Nitrogen",
// "state":"gas",
// },
// "O": map[string]string{
// "name":"Oxygen",
// "state":"gas",
// },
// "F": map[string]string{
// "name":"Fluorine",
// "state":"gas",
// },
// "Ne": map[string]string{
// "name":"Neon",
// "state":"gas",
// },
// }
// if el, ok := elements["Ne"]; ok {
// fmt.Println(el["name"], el["state"]);
// }

//Function
// xs := []float64{98,93,77,82,83}


// fmt.Println(average(xs));




// x,y:=f();
// fmt.Println(x,y);




//variadic function



// fmt.Println(add(1,2,3,4,5,6,7,8,9,7,6,5,4,3,2,4,6))



    // add := func(x , y int) int {
    // return x + y
    // }
    // fmt.Println(add(1,1))



//     x := 0
// increment := func() int {
// x++
// return x
// }
// fmt.Println(increment())
// fmt.Println(increment())

// fmt.Println(factorial(3))


//Defer ,Panic & Recover

//  defer second()
//  first()

//  f, _ := os.Open()
// defer f.Close()


// defer func() {
//     str := recover()
//     fmt.Println(str)
//     }()
//     panic("PANIC")

//Pointers

// 

//Stuct and interface




// var rx1, ry1 float64 = 0, 0
// var rx2, ry2 float64 = 10, 10
// var cx, cy, cr float64 = 0, 0, 5
// fmt.Println(rectangleArea(rx1, ry1, rx2,
// ry2))
// fmt.Println(circleArea(cx, cy, cr))

// type Circle struct{
//     x,y,r float64;
// }



// c:=Circle{1,4,5}
// r:=Rectangle{1,0,3,4}


// fmt.Println(r.area())


// p:=Person{"chuks"}
// p.Talk();


// a:=new(Android);
// a.Talk();

// fmt.Println(totalArea(&c, &r))









//CONCURRENCY

//    go f(0)
// var input string
// fmt.Scanln(&input)
// fmt.Println( "this is the input",input)


    // for i := 0; i < 10; i++ {
    // go f(i)
    // }
    // var input string
    // fmt.Scanln(&input)


    //CHANNNELS
//     var c chan string = make(chan string)
// go pinger(c)
// go ponger(c)
// go printer(c)
// var input string
// fmt.Scanln(&input)


//CHANNEL DIRECTION


//SELECT
    // c1 := make(chan string)
    // c2 := make(chan string)
    // go func() {
    // for {
    // c1 <- "from 1"
    // time.Sleep(time.Second * 0)
    // }
    // }()
    // go func() {
    // for {
    // c2 <- "from 2"
    // time.Sleep(time.Second * 0)
    // }
    // }()
    // go func() {
    // for {
    //     select {
    //     case msg1 := <- c1:
    //     fmt.Println("Message 1", msg1)
    //     case msg2 := <- c2:
    //     fmt.Println("Message 2", msg2)
    //     case <- time.After(time.Second):
    //     fmt.Println("timeout")

    // default:
    //     fmt.Println("nothing ready")
        
    //     }

    


    






    
    // }
    // }()
    // var input string
    // fmt.Scanln(&input)

        // xs := []float64{1,2,3,4}
        // avg := m.Average(xs)




        
        // fmt.Println(avg)


        //TESTING




        //strings

//         fmt.Println(
//             // true
//             strings.Contains("test", "es"),
//             // 2
//             strings.Count("test", "t"),
//             // true
//             strings.HasPrefix("test", "te"),
//             // true
//             strings.HasSuffix("test", "st"),
//             // 1
//             strings.Index("test", "s"),
//             // "a-b"
//             strings.Join([]string{"a","b"}, "-"),
//             // == "aaaaa"
//             strings.Repeat("a", 5),
//             // "bbaa"
//             strings.Replace("aaaa", "a", "b", 3),
//             // []string{"a","b","c","d","e"}
//             strings.Split("a,b,c,d,e", ","),
//             // "test"
//         strings.ToLower("TEST"),
// // "TEST"
//         strings.ToUpper("test"),
// )


// fmt.Println(arr, str);

//READ AND WRITE IO

    
//  arr := []byte("test")
//  str := string([]byte{'t','e','s','t'})

      
    
//  var buf bytes.Buffer
//  buf.Write([]byte("test"))



// file, err := os.Open("test.txt")
// if err != nil {
// // handle the error here
// fmt.Println("err:",err)

// return
// ;
// }


// defer file.Close()
// // get the file size
// stat, err := file.Stat()
// if err != nil {
//     fmt.Println()
// return
// }
// // fmt.Println(stat.Size())
// // read the file
// bs := make([]byte, stat.Size())
// _, err = file.Read(bs)
// if err != nil {
// return
// }
// str := string(bs)
// fmt.Println(str)

//SHORTER METHOD TO READ


// bs,err:=os.ReadFile("test.txt");
// if err!=nil{
//     return;

// }
// str:=string(bs);

// fmt.Println(str);


//CREATE AND WRITE FILE

file,err:=os.Create("chuks.txt");

if err!=nil{
    fmt.Println(err)
    return

}
defer file.Close();

file.WriteString("CHUKSONCHAIN");
fmt.Println("file created");

txt,err:=os.ReadFile("ojieh.txt");

fmt.Println(string(txt));



// dir, err := os.Open(".")
// if err != nil {
// return
// }
// defer dir.Close()
// fileInfos, err := dir.Readdir(-1)
// if err != nil {
// return
// }

// for i:=0; i<len(fileInfos); i++{
// fmt.Println(fileInfos[i].Name());
// }
// for _, fi := range fileInfos {
// fmt.Println(fi.Name())
// }
//

//FILEPATH

// filepath.Walk(".",func(path string,info os.FileInfo,err error)error{
//     fmt.Println(path)
//     return nil
// })


//Errors
// err := errors.New("error message")
// fmt.Println(err)


//container and sort
//sort
// kids := []Person{
//     {"Jill",9},
//     {"Jack",10},
//     }
//     sort.Sort(ByName(kids))
//     fmt.Println(kids)
//


//HASHES CRYPTOCRAPHY

// h := crc32.NewIEEE()
// h.Write([]byte("test"))
// v := h.Sum32()
// fmt.Println(v)

// h1,err:=getHash("ojieh.txt");

// if err!=nil{
//     return 
// }
// h2,err:=getHash("test.txt");

// if err!=nil{
//     return 
// }


// fmt.Println(h1,h2,h1==h2);


// http.HandleFunc("/", handler)

// // Start the server on port 8080
// fmt.Println("Server is listening on port 8080...")
// err := http.ListenAndServe(":8080", nil)
// if err != nil {
//     fmt.Println("Error starting server:", err)
// }


// http.HandleFunc("/hello", hello)
// http.ListenAndServe(":9000", nil)


//rpc



}


// func average(xs []float64) float64 {
//     // panic("Not Implemented")

//     total := 0.0;
//     for i:=0; i<len(xs); i++{
//         total+=xs[i];

//     }
//     return total / float64(len(xs));


    

// }



// func f()(int,int){
//     return 5,0
// }



//variadic function
// func add(args ...int) int {
//     total := 0
//     for _, v := range args {
//     total += v
//     }
//     return total
//     }




//recursion

// func factorial(x uint) uint {
//     if x == 0 {
//     return 1
//     }
//     return x * factorial(x-1)
//     }


//Defer ,Panic & Recover

// func first() {
//     fmt.Println("1st")
//     }
//     func second() {
//     fmt.Println("2nd")
//     }


//pointer
// func zero(xPtr *int) {
//     *xPtr = 0
//     }



//Sturct and interface
//     func rectangleArea(x1, y1, x2, y2 float64) float64 {
//         l := distance(x1, y1, x1, y2)
//         w := distance(x1, y1, x2, y1)
//         return l * w
    
//     }
                
    


//     func circleArea(x, y, r float64) float64 {
//     return math.Pi  *  r * r
//     }

// type Circle struct{
//     x,y,r float64;
// }




// func (c *Circle) area() float64 {
//     return math.Pi * c.r * c.r
// }
// func distance(x1, y1, x2, y2 float64) float64 {
//     a := x2 -  x1;
//     b := y2 - y1
//     return math.Sqrt(a*a + b*b)
//     }
//     type Rectangle struct{
//         x1,y1,x2,y2 float64
//     }


//     func (r *Rectangle) area() float64{

//         l:=distance(r.x1,r.y2,r.x1,r.y2)
//         w:=distance(r.x1,r.y1,r.x2,r.y1)
//         return l*w;

//     }











// //EMBEDDED TYPES

// type Person struct {



//     Name string
//     }
//     func (p *Person) Talk() {
//     fmt.Println("Hi, my name is", p.Name)
//     }



//     type Android struct {
//         Person
//         Model string
//         }



// //INTERFACES
// type Shape interface {
//     area() float64
//     }


//     func totalArea(shapes ...Shape) float64 {
//         var area float64
//         for _, s := range shapes {
//         area += s.area()
//         }
//         return area
//         }



    //CONCURRENCY


    //GOROUNTINUES

    // func f(n int) {
    //     for i := 0; i < 10; i++ {
    //     fmt.Println(n, ":", i)
    //     }
    //     }


    // func f(n int) {
    //     for i := 0; i < 10; i++ {
    //     fmt.Println(n, ":", i)
    //     amt := time.Duration(rand.Intn(250))
    //     time.Sleep(time.Millisecond * amt)
    //     }
    //     }




    //CHANNELS


    // func pinger(c chan string) {


    //     for i := 0; ; i++ {
    //     c <- "ping"
    //     }
    //     }
    //     func printer(c chan string) {
    //     for {
    //     // msg := <- c
    //     fmt.Println(<- c)
    //            amt := time.Duration(rand.Intn(250))

    //     time.Sleep(time.Millisecond * amt)
    //     }
    // }


    // func ponger(c chan string) {
    //     for i := 0; ; i++ {
    //     c <- "pong"
    //     }
    //     }


    //READ AND WRITE

    // func Copy(dst oi.Writer, src oi.Reader) (written int64, err error) 

    //sort


    // type Person struct {
    //     Name string
    //     Age int
    //     }
    //     type ByName []Person
    //     func (this ByName) Len() int {
    //     return len(this)
    //     }
    //     func (this ByName) Less(i, j int) bool {
    //     return this[i].Name < this[j].Name
    //     }
    //     func (this ByName) Swap(i, j int) {
    //     this[i], this[j] = this[j], this[i]
    //     }


    //HAshes


    // func getHash(filename string)(uint32,error){
    //     bs,err:=os.ReadFile(filename);
    //     if err != nil{
    //         return 0,err
    //     }
    //     h :=crc32.NewIEEE()
    //     h.Write(bs)
    //     return h.Sum32(),nil
    
    // }
    //SERVER
    // func handler(res http.ResponseWriter, req *http.Request) {
    //     // Set the Content-Type header to text/html
    //     res.Header().Set("Content-Type", "text/html")
    
    //     // Write the HTML response
    //     io.WriteString(res, `<!DOCTYPE html>
    //     <html>
    //     <head>
    //         <title>Hello my name is chuks!</title>
    //     </head>
    //     <body>
    //         <h1>Hello, World!</h1>
    //     </body>
    //     </html>`)
    // }


    // func hello(res http.ResponseWriter, req *http.Request) {
    //     res.Header().Set(
    //     "Content-Type",
    //     "text/html",
    //     )
    //     io.WriteString(
    //     res,
    //     `<doctype html>
    //     <html>
    //     <head>
    //             <title>Hello World</title>

    //     </head>
    //     <body>
    //     </body>
    //     </html>`,
    //     )
    //     }


    //rpc