package main


import(
        "fmt"
        "net"
)



var ip string = "192.168.0."


func main() {
        for a:=1; a<=255; a++ {
                num := fmt.Sprintf("%s%d",ip,a)
                for p:=1; p<=65000; p++ {
                        conn, err := net.Dial("tcp",fmt.Sprintf("%s:%d", num, p))
                        if err == nil {
                                fmt.Printf("open %s:%d",num,p)
                                conn.Close()
                        }
                }
        }
}
//I_made_it
