package repos

import "fmt"

func DbDeleteMatch(id int) error {
	_, err := db.Exec(fmt.Sprintf(`
	delete from matches where id = %d
	`, id))

	return err
}
