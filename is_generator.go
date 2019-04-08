package main

import "fmt"
import "math/big"
/*******
This funcion take in a large prime hex number as string, p, and checks if hex number g, as a string, is a generator $
*****/
func is_generator(p_string *string, g_string *string) bool{

    big1 := big.NewInt(1)
    big2 := big.NewInt(2)
    p := new(big.Int)
    modresult := new(big.Int)
    q := new(big.Int)
    g := new(big.Int)

    p.SetString(*p_string, 16)
    g.SetString(*g_string, 16)
    // If p is not prime return false
    if !p.ProbablyPrime(20){
        return false
    }
    qtemp := big.NewInt(0).Sub(p, big.NewInt(1))  // qtemp is p-1
    q = q.Div(qtemp, big2)                        // q = p-1 / 2
    // If q is not prime return false
    if !q.ProbablyPrime(20){
        return false
    }
    // Check is g mod p = 1, if yes g is not a generator
    modresult = modresult.Mod(g, p)
    if modresult.Cmp(big1) == 0{
        return false
    }
    //fmt.Println("g mod p = ", modresult)
    // Check if g^2 mod p = 1, if yes g is not a generator
    modresult = modresult.Exp(g, big2, p)
    if modresult.Cmp(big1) == 0{
        return false
    }
    //fmt.Println("g^2 mod p = ", modresult)
    // Check is g^q mod p = 1, if yes g is not a generator
    modresult = modresult.Exp(g, q, p)
    //fmt.Println("g^q mod p = ", modresult)
    if modresult.Cmp(big1) == 0{
        return false
    } else{
        return true
    }
}


func main(){
    p := "FFFFFFFFFFFFFFFFC90FDAA22168C234C4C6628B80DC1CD129024E088A67CC74020BBEA63B139B22514A08798E3404DDEF9519B3CD$
    g := "1f"
    fmt.Println(is_generator(&p, &g))
    
}
