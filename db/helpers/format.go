package helpers

import (
	"fmt"
	"log"

	"github.com/alex-305/bookbackend/models"
)

func Format(query string, o models.SortOptions) string {
	log.Printf("%+v", o)
	return query + fmt.Sprintf(` ORDER BY %s %s LIMIT %d OFFSET %d`, o.GetBy(), o.GetDirection(), o.GetLimit(), o.GetPage()*o.GetLimit())
}
