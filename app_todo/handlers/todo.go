package handlers

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"todo/db"
	"net/http"

)

// Handler for GET /todo
func Handler(ctx *gin.Context)  {
	todos, err := db.GetTodos()
	if err != nil {
		 ctx.String(http.StatusBadRequest,"failed to get todos\n")
		 return
	}
	ctx.JSON(http.StatusOK, todos)	
}

type Requirement struct {
	Purpose string `json:"purpose"`
}

type Completed struct {
	Done bool `json:"done"`
}

// Handler for POST /todo
func AddTodo(ctx *gin.Context)  {
	var req Requirement

	err := ctx.BindJSON(&req)
	if err != nil {
		 ctx.String(http.StatusNotFound,"Invalid input")
		return
	}
	if req.Purpose == "" {
		 ctx.String(http.StatusNotFound,"Todo empty.")
		return 
	}
	err = db.AddTodo(req.Purpose)
	if err != nil {
		 ctx.String(http.StatusBadRequest,"Invalid Todo: couldn't be added")
		 return
	}
	 ctx.String(http.StatusOK,"Added Todo list")
}

// Handler for DELETE /todo/:id
func DeleteTodo(ctx *gin.Context)  {
	String_id := ctx.Param("id")
	Number_id, err := strconv.Atoi(String_id)
	if err != nil {
		ctx.String(http.StatusBadRequest,"Invalid input: not a number ID")
		return
	}
	if err := db.DeleteTodo(Number_id); err != nil {
		 ctx.String(http.StatusNotFound,"Id does not exist")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "the todo is deleted succesfully"})
}

// Handler for PATCH /todo/:id
func Update(ctx *gin.Context)  {
	String_id := ctx.Param("id")
	Number_id, err := strconv.Atoi(String_id)
	if err != nil {
		 ctx.String(http.StatusBadRequest,"Invalid input: not a number ID")
		return
	}
	var req Completed
	err = ctx.BindJSON(&req)
	if err != nil {
		 ctx.String(http.StatusNotFound,"Failed to parse Body")
		return
	}
	if err := db.UpdateTodo(Number_id, req.Done); err != nil {
		 ctx.String(http.StatusBadRequest,"Failed to update todo")
		 return
	}
	 ctx.String(http.StatusOK,"Fully updated the todo")
}