package model

var Users = map[string]*User{
	"user": {
		Id:    "user",
		Email: "user1@example.com",
		Permissions: map[string]bool{
			ViewExperience: true,
			Ialamuie:       true,
		},
	},
	"supplier": {
		Id:    "supplier",
		Email: "supplier@supplier.com",
		Permissions: map[string]bool{
			ViewExperience: true,
			Ialamuie:       true,
			Dalamuie:       true,
		},
	},
	"admin": {
		Id:    "admin",
		Email: "admin@boss.com",
		Permissions: map[string]bool{
			ViewExperience:     true,
			Dalamuie:           true,
			ManageAvailability: true,
			Ialamuie:           true,
		},
	},
}
