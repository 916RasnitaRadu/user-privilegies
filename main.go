package main

import (
	"fmt"
	"log"
	"net/http"
	"users-privi/handlers"
	"users-privi/middleware"
	"users-privi/model"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	// only different permissions on different routes
	r.Post("/view-experience", middleware.RequirePermission([]string{model.ViewExperience}, handlers.ViewExperience))
	r.Post("/esti-gata", middleware.RequirePermission([]string{model.Dalamuie, model.Ialamuie}, handlers.HandleEstiGata))
	r.Post("/iala", middleware.RequirePermission([]string{model.Ialamuie}, handlers.HandleIalamuie))
	r.Post("/dala", middleware.RequirePermission([]string{model.Dalamuie}, handlers.HandleDalamuie))
	r.Post("/manage-availability", middleware.RequirePermission([]string{model.EditExperience, model.ManageAvailability}, handlers.HandleManageAvailability))

	// using roles
	r.Group(func(r chi.Router) {
		r.Use(middleware.OnlyUser)

		r.Post("/user/view-experience", handlers.ViewExperience)
		r.Post("/user/manage-availability", handlers.HandleManageAvailability)
	})

	r.Group(func(r chi.Router) {
		r.Use(middleware.OnlySupplier)

		r.Post("/supplier/iala", handlers.HandleIalamuie)
		r.Post("/supplier/dala", handlers.HandleDalamuie)
		r.Post("/supplier/esti-gata", handlers.HandleEstiGata)
	})

	fmt.Println("🚀 Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
