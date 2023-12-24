package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ishanshre/GoRestApiExample/internals/models"
)

type Authors map[string][]models.Author

func (h *handler) HomeHandlerHtmx(c *gin.Context) {
	authors, err := h.repo.GetAllAuthors()
	if err != nil {
		c.HTML(http.StatusBadRequest, "index.tmpl", gin.H{
			"error": "Error getting authors",
		})
		return
	}
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"authors": authors,
	})
}

func (h *handler) AddAuthorHandlerHtmx(c *gin.Context) {
	if err := c.Request.ParseForm(); err != nil {
		c.HTML(http.StatusBadRequest, "index.tmpl", gin.H{
			"error": "error parsing form",
		})
		return
	}
	first_name := c.Request.Form.Get("first_name")
	last_name := c.Request.Form.Get("last_name")
	email := c.Request.Form.Get("email_address")
	author := models.Author{
		FirstName: first_name,
		LastName:  last_name,
		Email:     email,
	}
	data, err := h.repo.CreateAuthor(author)
	if err != nil {
		c.HTML(http.StatusBadRequest, "index.tmpl", gin.H{
			"error": "error in creating author",
		})
		return
	}
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"authors": data,
	})
}
