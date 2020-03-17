# Greet Devs!
welcome_message.greet

# Golang linter for code documentation
golint.lint

# Make sure if PR have assignee
failure "This PR does not have any assignees yet." unless github.pr_json["assignee"]

# Make it more obvious that a PR is a work in progress and shouldn't be merged yet
warn "PR is classed as Work in Progress" if github.pr_labels.include? "WIP"
warn "PR is classed as Work in Progress" if github.pr_title.include? "[WIP]"

# Warn when there is a big PR
warn("Big PR") if git.insertions > 500

# Provide PR Summary
if github.pr_body.length < 5
    failure "Please provide a summary in the Pull Request description"
end

can_merge = github.pr_json["mergeable"]
warn("This PR cannot be merged yet.", sticky: false) unless can_merge
