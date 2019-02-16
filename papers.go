package arxivlib

import (
	"errors"
	"time"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// An Paper is a research Paper residing in arXiv
type Paper struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	ArxivID    string             `json:"arxiv_id"`
	Title      string             `json:"title"`
	Published  time.Time          `json:"published"`
	Updated    time.Time          `json:"updated"`
	Abstract   string             `json:"abstract"`
	Authors    []string           `json:"authors"`
	LinkPDF    string             `json:"link_pdf"`
	LinkPage   string             `json:"link_page"`
	Categories []string           `json:"categories"`
	Ratings    []Rating           `json:"ratings"`
}

// A Rating is a score given to a paper by a User
type Rating struct {
	UserID  primitive.ObjectID `json:"user_id" bson:"user_id"`
	Score   int                `json:"score" bson:"score"`
	Comment string             `json:"comment" bson:"comment"`
	Date    time.Time          `json:"date" bson:"date"`
}

// PapersService interacts with the paper-related endpoints in arxivlib's API
type PapersService interface {
	// Get a paper
	Get(id primitive.ObjectID) (*Paper, error)

	// List papers
	List(opt *PaperListOptions) ([]*Paper, error)

	// Update a paper
	Update(paper *Paper) (updated bool, err error)

	// Upload paper(s)
	Upload(paper []*Paper) (uploaded bool, err error)

	// Remove a paper
	Remove(id primitive.ObjectID) (removed bool, err error)

	// Add a rating to a paper
	AddRating(id primitive.ObjectID, r *Rating) (added bool, err error)
}

var (
	// ErrPaperNotFound is a failure to find a specified paper
	ErrPaperNotFound = errors.New("paper not found")

	// ErrRatingNotFound is a failure to find a rating for a specified paper
	ErrRatingNotFound = errors.New("rating not found")
)

// A PaperListOptions represents a search filter for listing papers
type PaperListOptions struct {
	Title      string    `json:"title,omitempty" form:"title,omitempty"`
	Author     string    `json:"author,omitempty" form:"author,omitempty"`
	Updated    time.Time `json:"updated,omitempty" form:"updated,omitempty"`
	Abstract   string    `json:"abstract,omitempty" form:"abstract,omitempty"`
	Categories []string  `json:"categories,omitempty" form:"categories,omitempty"`
	MaxResults int       `json:"max_results,omitempty" form:"max_results,omitempty"`
}
