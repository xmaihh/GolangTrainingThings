package controllers

import (
	"beeticket/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about Ticket
type TicketController struct {
	beego.Controller
}

// @Title Create
// Description create ticket
// @Param	body	body	models.Ticket 	true	"The Ticket content"
// @Success 200 {string} models.Ticket.TicketId
// @Failure 403 body is empty
// @router	/	[post]
func (t *TicketController) Post() {
	var ticket models.Ticket
	json.Unmarshal(t.Ctx.Input.RequestBody, &ticket)
	ticketId := models.AddTicket(ticket)
	t.Data["json"] = map[string]string{"ticketId": ticketId}
	t.ServeJSON()
}

// @Title Get
// @Description get a ticket by ticketId
// @Param 	ticketId 	path 	string 	true 	"The ticketId you want to get"
// @Success 200 {object} models.Ticket
// @Failure 403 :ticketId is empty
// @router	/:ticketId [get]
func (t *TicketController) Get() {
	ticketId := t.Ctx.Input.Param(":ticketId")
	if ticketId != "" {
		ticket, err := models.GetTicket(ticketId)
		if err != nil {
			t.Data["json"] = err.Error()
		} else {
			t.Data["json"] = ticket
		}
	}
	t.ServeJSON()
}

// @Title GetAll
// @Description get all tickets
// @Success 200 {object} models.Ticket
// @Failure 403 :ticketId is empty
// @router / [get]
func (t *TicketController) GetAll() {
	tickets := models.GetAllTickets()
	t.Data["json"] = tickets
	t.ServeJSON()
}

// @Title Update
// @Description update the ticket
// @Param	ticketId 	path	string	true		"The ticketId you want to update"
// @Param	body	body	models.Ticket	true		"body for ticket content"
// @Success 200 {object} models.Ticket
// @Failure 403	:ticketId is empty
// @router /:ticketId [put]
func (t *TicketController) Put() {
	ticketId := t.GetString(":ticketId")
	if ticketId != "" {
		var ticket models.Ticket
		json.Unmarshal(t.Ctx.Input.RequestBody, &ticket)

		tt, err := models.UpdateTicket(ticketId, &ticket)
		if err != nil {
			t.Data["json"] = err.Error()
		} else {
			t.Data["json"] = tt
		}
	}
	t.ServeJSON()
}

// @Title Delete
// @Description delete the ticket
// @Param	ticketId 	path 	string	true	"The ticketId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 ticketId is empty
// @router /:ticketId [delete]
func (t *TicketController) Delete() {
	ticketId := t.Ctx.Input.Param(":ticketId")
	models.DeleteTicket(ticketId)
	t.Data["json"] = "delete success!"
	t.ServeJSON()
}
