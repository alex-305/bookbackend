package helpers

import (
	"fmt"

	"github.com/alex-305/bookbackend/internal/models"
)

func Format(query string, o models.SortOptions) string {
	return query + fmt.Sprintf(` ORDER BY %s %s LIMIT %d OFFSET %d`, o.GetBy(), o.GetDirection(), o.GetLimit(), o.GetPage()*o.GetLimit())
}
