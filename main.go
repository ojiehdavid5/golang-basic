package main

import "fmt"

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

 

x := make([]float64,5,10);

x[0] = 98
x[1] = 93
x[2] = 77
x[3] = 82
x[4] = 83


fmt.Println(x)




}