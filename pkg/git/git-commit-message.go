package git

func GenerateCommitMessage(status string) string {

    message := "🚀 Update project"
    message = "🚧 Work in progress"
    message = "✨ Add new feature"

    message = "🔨 Refactor code"
    message = "👌 Improve code structure"
    message = "🐎 Improve performance"

    message = "🗑️ Remove unused code"
    message = "🔥 Remove code or files"

    message = "🐛 Fix bugs"
    message = "🚑 Fix critical bug"
    message = "🔒 Fix security issues"
    message = "🚨 Fix linter warnings"

    message = "👷 Add CI build system"
    message = "🔧 Add or update configuration files"

    message = "🚀 Create new version"
    message = "🔖 Release version"
    message = "🚀 Deploy stuff"

    message = "📝 Update documentation"
    message = "📝 Update license"

    return message
}
