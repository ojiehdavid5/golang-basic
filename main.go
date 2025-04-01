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

 
//normal method
// x := make([]float64,5,10);

// x[0] = 98
// x[1] = 93
// x[2] = 77
// x[3] = 82z
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

// elements := make(map[string]string)
// elements["H"] = "Hydrogen"
// elements["He"] = "Helium"
// elements["Li"] = "Lithium"
// elements["Be"] = "Beryllium"
// elements["B"] = "Boron"
// elements["C"] = "Carbon"
// elements["N"] = "Nitrogen"
// elements["O"] = "Oxygen"
// elements["F"] = "Fluorine"
// elements["Ne"] = "Neon"



// name, ok := elements["H"]
// fmt.Println(name, ok)


elements := map[string]map[string]string{
    "H": map[string]string{
    "name":"Hydrogen",
    "state":"gas",
    },
    "He": map[string]string{
    "name":"Helium",
    "state":"gas",
    },
    "Li": map[string]string{
    "name":"Lithium",
    "state":"solid",
    },"Be": map[string]string{
"name":"Beryllium",
"state":"solid",
},
"B": map[string]string{
"name":"Boron",
"state":"solid",
},
"C": map[string]string{
"name":"Carbon",
"state":"solid",
},
"N": map[string]string{
"name":"Nitrogen",
"state":"gas",
},
"O": map[string]string{
"name":"Oxygen",
"state":"gas",
},
"F": map[string]string{
"name":"Fluorine",
"state":"gas",
},
"Ne": map[string]string{
"name":"Neon",
"state":"gas",
},
}
if el, ok := elements["Ne"]; ok {
fmt.Println(el["name"], el["state"]);
}
}

