package main
import "fmt"

type Student struct {
 // Isi struct ini
name string
    age int
}

func (s *Student) Introduction() {
   // Tulis kode disini
   nama := s.name
   umur := s.age
   fmt.Printf("Hello, my name is %s. Im %d years old", nama, umur)
}

func main() {
 // Tulis kode disini
 var student1 = Student {
        name: "Sammy",
        age: 17,
    }
  student1.Introduction()
}