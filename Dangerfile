# Make sure if PR have assignee
failure "This PR does not have any assignees yet." unless github.pr_json["assignee"]

# Make sure one of the reviewer is from O2O
# Requested Reviewer
#requested_reviewers = github.pr_json["requested_reviewers"]
# Actual Reviewer
# pr_num = github.pr_json["number"]
# reviews = github.api.pull_request_reviews("mpermperpisang/golang-automation", pr_num)
# actual_reviewers = reviews.map {|u| u["user"]}
# PR Reviewer
# reviewers = requested_reviewers + actual_reviewers
# pr_reviewers = reviewers.map {|u| u["login"]}
# O2O Reviewer
# o2o_reviewer = ["mpermperpisang"]

# unless o2o_reviewer.any?{|x| pr_reviewers.include?(x)}
#   failure "Please request a review from mpermperpisang"
# end

# Make it more obvious that a PR is a work in progress and shouldn't be merged yet
warn "PR is classed as Work in Progress" if github.pr_labels.include? "WIP"
warn "PR is classed as Work in Progress" if github.pr_title.include? "[WIP]"

# Warn when there is a big PR
warn("Big PR") if git.insertions > 500

# Provide PR Summary
if github.pr_body.length < 5
    failure "Please provide a summary in the Pull Request description"
end

# Provide automation running screenshot
format = [".png", ".jpg", ".gif", ".jpeg"]

unless format.any?{|x| github.pr_body.include?(x)}
  failure "Please provide automation running screenshot to prove there is no syntax error"
end

can_merge = github.pr_json["mergeable"]
warn("This PR cannot be merged yet.", sticky: false) unless can_merge

# if git.modified_files.include? "main.go"
#   warn "Please assign @mpermperpisang as reviewer"
# end
