package pointers

import "fmt"

func zeroval(ival int) {
    ival = 0
}

func zeroptr(iptr *int) {
    *iptr = 0
}

func PlayWithPointers(){
	i := 1;
    a := 5
    o := 24

    var optr,aptr *int 

    // vemos como se modifican o no, ya que zeroval al ser un parametro por valor, crea una copia en el scope de la funcion, contrario a zeroptr que se maneja con parametro por referencia

    fmt.Println("initial:", i)

    zeroval(i)
    fmt.Println("zeroval:", i)

    zeroptr(&i)
    fmt.Println("zeroptr:", i)

    fmt.Println("pointer:", &i)

    fmt.Println("initial a:", a) 
    
    fmt.Println("initial o:", o) 

    aptr = &a  // le asigno donde apuntar

    *aptr = o // mi pointer toma el valor de o

    optr = &o

    *aptr = 115
    
    aptr = &o

    *aptr = 666

    *optr -=1

    fmt.Println(" a value:", a)
    fmt.Println(" o value:", o)
    fmt.Println("pointer a:", &a)
    fmt.Println("pointer o:", &o)

    fmt.Println("pointer aptr:", aptr) // ambos se encuentran apuntando a o
    fmt.Println("pointers:", optr)
}