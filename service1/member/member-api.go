package member

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
* Add a new member via REST call
 */
func AddMember(c *gin.Context) {

	var newMember Member

	if err := c.BindJSON(&newMember); err != nil {
		return
	}

	Members = append(Members, newMember)

	ctx := context.Background()
	Produce(newMember, ctx)

	c.IndentedJSON(http.StatusCreated, newMember)
}
