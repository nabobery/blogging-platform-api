package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"blogging-platform-api/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// BlogPostController holds the DB instance for use in HTTP handlers.
type BlogPostController struct {
	DB *gorm.DB
}

// NewBlogPostController returns a new instance of BlogPostController.
func NewBlogPostController(db *gorm.DB) *BlogPostController {
	return &BlogPostController{DB: db}
}

// CreateBlogPost handles POST /posts and creates a new blog post.
func (c *BlogPostController) CreateBlogPost(w http.ResponseWriter, r *http.Request) {
	var post models.BlogPost
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate required fields
	if post.Title == "" || post.Content == "" || post.Category == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	post.CreatedAt = time.Now().UTC()
	post.UpdatedAt = post.CreatedAt

	if err := c.DB.Create(&post).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(post); err != nil {
		log.Printf("Error encoding response in CreateBlogPost: %v", err)
	}
}

// GetBlogPost handles GET /posts/{id} to retrieve a single blog post.
func (c *BlogPostController) GetBlogPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	var post models.BlogPost
	if err := c.DB.First(&post, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Post not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(post); err != nil {
		log.Printf("Error encoding response in GetBlogPost: %v", err)
	}
}

// GetBlogPosts handles GET /posts and supports optional filtering by term.
func (c *BlogPostController) GetBlogPosts(w http.ResponseWriter, r *http.Request) {
	var posts []models.BlogPost
	term := r.URL.Query().Get("term")
	query := c.DB

	// If a search term is provided, perform a wildcard search on title, content, or category.
	if term != "" {
		likeTerm := "%" + term + "%"
		query = query.Where("title ILIKE ? OR content ILIKE ? OR category ILIKE ?", likeTerm, likeTerm, likeTerm)
	}

	if err := query.Find(&posts).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		log.Printf("Error encoding response in GetBlogPosts: %v", err)
	}
}

// UpdateBlogPost handles PUT /posts/{id} to update an existing blog post.
func (c *BlogPostController) UpdateBlogPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	var existingPost models.BlogPost
	if err := c.DB.First(&existingPost, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Post not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	var updatedPost models.BlogPost
	if err := json.NewDecoder(r.Body).Decode(&updatedPost); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate required fields
	if updatedPost.Title == "" || updatedPost.Content == "" || updatedPost.Category == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	existingPost.Title = updatedPost.Title
	existingPost.Content = updatedPost.Content
	existingPost.Category = updatedPost.Category
	existingPost.Tags = updatedPost.Tags
	existingPost.UpdatedAt = time.Now().UTC()

	if err := c.DB.Save(&existingPost).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(existingPost); err != nil {
		log.Printf("Error encoding response in UpdateBlogPost: %v", err)
	}
}

// DeleteBlogPost handles DELETE /posts/{id} to delete a blog post.
func (c *BlogPostController) DeleteBlogPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	if err := c.DB.Delete(&models.BlogPost{}, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Successful deletion returns 204 No Content
	w.WriteHeader(http.StatusNoContent)
}
