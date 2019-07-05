package models

import (
	"errors"
	"strconv"
	"time"
)

var (
	TicketList map[string]*Ticket
)

type Ticket struct {
	TicketId   string
	UserId     string
	Ticketcode string
	Usestatus  int
	Ticketdesc string
	Exp        string
	Field      string
}

func AddTicket(t Ticket) string {
	t.TicketId = "ticket_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	TicketList[t.TicketId] = &t
	return t.TicketId
}

func GetTicket(ticketId string) (u *Ticket, err error) {
	if t, ok := TicketList[ticketId]; ok {
		return t, nil
	}
	return nil, errors.New("Ticket not exists")
}

func GetAllTickets() map[string]*Ticket {
	return TicketList
}

func UpdateTicket(uid string, tt *Ticket) (a *Ticket, err error) {
	if t, ok := TicketList[uid]; ok {

		if tt.UserId != "" {
			t.UserId = tt.UserId
		}
		if tt.Ticketcode != "" {
			t.Ticketcode = tt.Ticketcode
		}
		if tt.Ticketdesc != "" {
			t.Ticketcode = tt.Ticketcode
		}
		if tt.Usestatus != -1 {
			t.Usestatus = tt.Usestatus
		}
		if tt.Exp != "" {
			t.Exp = tt.Exp
		}
		if tt.Field != "" {
			t.Field = tt.Field
		}
		return t, nil
	}
	return nil, errors.New("Ticket Not Exist")
}

func DeleteTicket(uid string) {
	delete(TicketList, uid)
}

// type Admin struct {
// 	ID        int64
// 	Username  string
// 	Password  string
// 	Totalcode int64
// 	Exp       string
// 	Userlimit int64
// }

// type User struct {
// 	ID         int64
// 	Username   string
// 	Password   string
// 	CreatedAt string
// 	UpdatedAt string
// }
