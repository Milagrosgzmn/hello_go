package my_structs

import "fmt"

func PlayWIthStructs(){
	type def1 struct{
        name string
        edad int
    }
    d :=  def1{
        "pepe",
        24, //si uso nueva linea debo aclarar con la ,
    }

    test := &d 
    fmt.Println("test value:", test.edad) // puedes acceder a los campos de un struct en un struct pointer - el pointer se desreferencia de forma automatica.
	
	fmt.Println("test value:", test.name) 
	
	test.name = "marta" // los structs son mutables

	fmt.Println("test value:", test.name) 

	dog := struct { // si solo lo vas a usar una vez, puedes crear un struct anonimo
        name   string
        isGood bool
    }{
        "Rex",
        true,
    }
	fmt.Println(dog);

}