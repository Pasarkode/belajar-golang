package controllers

import (
  "aahframework.org/aah.v0"
  "aah-app/app/models"
)

// App struct application controller
type App struct {
  *aah.Context
}

// Index method is application root API endpoint.

func (a *App) Index() {
  a.Reply().Ok().JSON(models.Greet{
    Message: "Welcome to aah framework - API application",
  })
}
