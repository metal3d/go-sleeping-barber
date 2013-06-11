package barbershop

import "fmt"
import "time"

// AFAIK a client is an Human
type Client struct {
    Human
}

// Client enters shop and tries to find a seat
func (c* Client) EnterShop(shop *Shop) {
    // try to find a seat 3 times...
    for i:=0 ; i<3; i++ {
        if i > 0 {
            // client is back
            fmt.Printf(">> %s is back\n", c.GetName())
        }
        select {
        case shop.WaitingRoom <- c:
            //le client a trouvé de la place
            fmt.Printf(":) %s found a seat\n", c.GetName())
            select {
                // on vérifie que le barbier dort pas
                case shop.Barbers.WakeMe<-c:
                    fmt.Printf("WAKE UP ! screams %s\n", c.GetName())
                default:
                    //rien à faire
            }
            return


        default:
            fmt.Printf("... %s will be back later ...\n", c.GetName())
            time.Sleep(time.Millisecond * 100)
        }
    }
    // no seat, get out
    fmt.Printf(":( %s didn't found a seat, get out...\n", c.GetName())
}

// create a new client (pointer)
func NewClient(name string) *Client {
    c := new(Client) // retourne un pointeur
    c.name = name
    return c
}
