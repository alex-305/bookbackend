package helpers

import (
	"fmt"

	"github.com/alex-305/bookbackend/internal/models"
)

func ListFormat(o models.SortOptions) string {
	return orderby(o) + paginate(o)
}

func orderby(o models.SortOptions) string {
	return fmt.Sprintf(` ORDER BY %s %s`, o.GetBy(), o.GetDirection())
}

func paginate(o models.SortOptions) string {
	return fmt.Sprintf(` LIMIT %d OFFSET %d`, o.GetLimit(), o.GetPage()*o.GetLimit())
}
