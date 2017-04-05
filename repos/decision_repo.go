package repos

import (
	"fmt"

	"github.com/fatih/structs"
	"github.com/jeffmcnd/clik/models"
)

func DbCreateDecision(id int, decision *models.DecisionForm) error {
	_, err := db.NamedExec(fmt.Sprintf(`
	insert into decisions (from_user_id, to_user_id, likes)
	values (%d, :UserId, :Likes)
	`, id), structs.Map(decision))

	if err != nil {
		return err
	}

	return nil
}
