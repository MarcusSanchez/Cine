package tmdb

type Genre int

const (
	Action                Genre = 28
	ActionAdventure       Genre = 10759
	Adventure             Genre = 12
	Animation             Genre = 16
	Comedy                Genre = 35
	Crime                 Genre = 80
	Documentary           Genre = 99
	Drama                 Genre = 18
	Family                Genre = 10751
	Fantasy               Genre = 14
	History               Genre = 36
	Horror                Genre = 27
	Kids                  Genre = 10762
	Music                 Genre = 10402
	Mystery               Genre = 9648
	News                  Genre = 10763
	Reality               Genre = 10764
	Romance               Genre = 10749
	ScienceFiction        Genre = 878
	ScienceFictionFantasy Genre = 10765
	Soap                  Genre = 10766
	Talk                  Genre = 10767
	TVMovie               Genre = 10770
	Thriller              Genre = 53
	War                   Genre = 10752
	WarAndPolitics        Genre = 10768
	Western               Genre = 37
)

func (g Genre) String() string {
	switch g {
	case Action:
		return "Action"
	case ActionAdventure:
		return "Action & Adventure"
	case Adventure:
		return "Adventure"
	case Animation:
		return "Animation"
	case Comedy:
		return "Comedy"
	case Crime:
		return "Crime"
	case Documentary:
		return "Documentary"
	case Drama:
		return "Drama"
	case Family:
		return "Family"
	case Fantasy:
		return "Fantasy"
	case History:
		return "History"
	case Horror:
		return "Horror"
	case Kids:
		return "Kids"
	case Music:
		return "Music"
	case Mystery:
		return "Mystery"
	case News:
		return "News"
	case Reality:
		return "Reality"
	case Romance:
		return "Romance"
	case ScienceFiction:
		return "Science Fiction"
	case ScienceFictionFantasy:
		return "Science Fiction & Fantasy"
	case Soap:
		return "Soap"
	case Talk:
		return "Talk"
	case TVMovie:
		return "TV Movie"
	case Thriller:
		return "Thriller"
	case War:
		return "War"
	case WarAndPolitics:
		return "War & Politics"
	case Western:
		return "Western"
	default:
		return "Unknown"
	}
}
