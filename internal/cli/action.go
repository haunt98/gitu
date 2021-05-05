package cli

const (
	appName  = "gitu"
	appUsage = "switch git user quickly"

	// flags
	nameFlag     = "name"
	emailFlag    = "email"
	nicknameFlag = "nickname"
	allFlag      = "all"

	// commands
	addCommand    = "add"
	switchCommand = "switch"
	statusCommand = "status"
	listCommand   = "list"
	deleteCommand = "delete"

	// flag usage
	nameUsage     = "gitconfig user.name"
	emailUsage    = "gitconfig user.email"
	nicknameUsage = "nickname to choose"
	allUsage      = "select all nicknames"

	// command usage
	addUsage    = "add git user"
	switchUsage = "switch git user"
	statusUsage = "show current name and email"
	listUsage   = "list all saved names and emails"
	deleteUsage = "delete saved name and email"
)

var (
	// flag aliases
	allAliases = []string{"a"}

	// command aliases
	addAliases    = []string{"a"}
	switchAliases = []string{"sw"}
	statusAliases = []string{"st"}
	listAliases   = []string{"l"}
	deleteAliases = []string{"d"}
)
