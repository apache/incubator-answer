package templaterender

import (
	"math"

	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service"
	"github.com/answerdev/answer/internal/service/tag"
	"github.com/google/wire"
)

// ProviderSetTemplateRenderController is template render controller providers.
var ProviderSetTemplateRenderController = wire.NewSet(
	NewTemplateRenderController,
)

type TemplateRenderController struct {
	questionService *service.QuestionService
	userService     *service.UserService
	tagService      *tag.TagService
}

func NewTemplateRenderController(
	questionService *service.QuestionService,
	userService *service.UserService,
	tagService *tag.TagService,

) *TemplateRenderController {
	return &TemplateRenderController{
		questionService: questionService,
		userService:     userService,
		tagService:      tagService,
	}
}

// Paginator page
// page : now page
// pageSize : Number per page
// nums : Total
// Returns the contents of the page in the format of 1, 2, 3, 4, and 5. If the contents are less than 5 pages, the page number is returned
func Paginator(page, pageSize int, nums int64) *schema.Paginator {
	if pageSize == 0 {
		pageSize = 10
	}

	var prevpage int //Previous page address
	var nextpage int //Address on the last page
	//Generate the total number of pages based on the total number of nums and the number of prepage pages
	totalpages := int(math.Ceil(float64(nums) / float64(pageSize))) //Total number of Pages
	if page > totalpages {
		page = totalpages
	}
	if page <= 0 {
		page = 1
	}
	var pages []int
	switch {
	case page >= totalpages-5 && totalpages > 5: //The last 5 pages
		start := totalpages - 5 + 1
		prevpage = page - 1
		nextpage = int(math.Min(float64(totalpages), float64(page+1)))
		pages = make([]int, 5)
		for i, _ := range pages {
			pages[i] = start + i
		}
	case page >= 3 && totalpages > 5:
		start := page - 3 + 1
		pages = make([]int, 5)
		prevpage = page - 3
		for i, _ := range pages {
			pages[i] = start + i
		}
		prevpage = page - 1
		nextpage = page + 1
	default:
		pages = make([]int, int(math.Min(5, float64(totalpages))))
		for i, _ := range pages {
			pages[i] = i + 1
		}
		prevpage = int(math.Max(float64(1), float64(page-1)))
		nextpage = page + 1
	}
	paginator := &schema.Paginator{}
	paginator.Pages = pages
	paginator.Totalpages = totalpages
	paginator.Prevpage = prevpage
	paginator.Nextpage = nextpage
	paginator.Currpage = page
	return paginator
}
