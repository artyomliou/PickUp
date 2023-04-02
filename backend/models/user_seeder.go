package models

func (s *Seeder) Users() {
	emails := []string{
		"artyomliou@gmail.com",
	}
	passwordHash, _ := HashPassword("12345678")
	for _, email := range emails {
		u := &User{
			Email: email,
		}
		DB().FirstOrCreate(u, User{
			Password: passwordHash,
		})
	}
}
