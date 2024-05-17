package access

import "errors"

func HasOwnershipAccess(username, ownername string) error {
	if username != ownername {
		return errors.New("user does not have access to this resource")
	}
	return nil
}
