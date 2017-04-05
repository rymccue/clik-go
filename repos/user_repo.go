package repos

import (
	"fmt"

	"github.com/fatih/structs"
	"github.com/jeffmcnd/clik/models"
	"github.com/jeffmcnd/clik/utils/crypto"
)

func DbGetUser(id int) (*models.User, error) {
	user := &models.User{}
	tx := db.MustBegin()
	err := tx.Get(user, `select * from users where id = $1`, id)
	return user, err
}

func DbCreateUser(user *models.User) error {
	user.Salt = crypto.GenerateSalt()
	user.Password = crypto.HashPassword(user.Password, user.Salt)
	_, err := db.NamedExec(`
	insert into users
	(age, birthday, career, email, start_age, end_age, gender, info, first_name, last_name, looking_for, school, password, salt)
	values
	(:Age, to_date(:Birthday, 'YYYY-MM-DD'), :Career, :Email, :StartAge, :EndAge, :Gender, :Info, :FirstName, :LastName, :LookingFor, :School, :Password, :Salt)
	`, structs.Map(user))

	if err != nil {
		return err
	}

	tx := db.MustBegin()
	err = tx.Get(user, `select * from users where email = $1`, user.Email)

	return err
}

/*
DbGetUserQueue returns a queue of users that haven't been decided on by the requesting user
based on his/her criteria.DbCreateUser

This is done by joining the users and images tables, retrieving all the users that both
match the requesting users criteria and whose criteria include the requesting users info.
Only users that haven't been decided on by the requesting user will be returned.
*/
func DbGetUserQueue(id int) (*models.UserQueue, error) {
	user, err := DbGetUser(id)
	if err != nil {
		return nil, err
	}

	userQueue := models.UserQueue{}

	/*
		Note the use of Sprintf to use the user.Age property twice. This is necessary because
		the PrepareNamed function returns a statement using sequential $n values where property
		values should be inserted even if you reuse the property name.

		For example:
		"users.start_age <= :Age and users.end_age >= :Age"
		becomes
		"users.start_age <= $1 and users.end_age >= $2"

		So, instead we insert the value using Sprintf ahead of time.
	*/
	stmt, err := db.PrepareNamed(fmt.Sprintf(`
	select users.*, images.url
	from users join images
	on users.id = images.user_id
	where users.id not in (
		select to_user_id
		from decisions
		where from_user_id = :Id
	)
	and users.gender = :LookingFor
	and users.looking_for = :Gender
	and users.start_age <= %v and users.end_age >= %v
	and users.age between :StartAge and :EndAge
	and images.index = 0
	`, user.Age, user.Age))

	if err != nil {
		return nil, err
	}

	err = stmt.Select(&userQueue, structs.Map(user))

	return &userQueue, err
}

func DbGetUserMatches(id int) (*models.Matches, error) {
	user, err := DbGetUser(id)
	if err != nil {
		return nil, err
	}

	matches := models.Matches{}

	stmt, err := db.PrepareNamed(fmt.Sprintf(`
	select users.id, users.age, users.career, users.first_name, users.last_name, users.school, images.url, matches.id as match_id
	from users join images on users.id = images.user_id
	join matches on users.id = matches.user_1_id or users.id = matches.user_2_id
	where images.index = 0
	and ((users.id in (
		select user_2_id
		from matches
		where user_1_id = %v
	)
	or users.id in (
		select user_1_id
		from matches
		where user_2_id = %v
	)))
	`, id, id))

	if err != nil {
		return nil, err
	}

	err = stmt.Select(&matches, structs.Map(user))

	return &matches, err
}
