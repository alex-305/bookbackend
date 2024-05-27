package access

import (
	"errors"

	"github.com/alex-305/bookbackend/internal/models"
)

func HasOwnershipAccess(username models.UserID, ownername models.UserID) error {
	if username != ownername {
		return errors.New("user does not have access to this resource")
	}
	return nil
}
