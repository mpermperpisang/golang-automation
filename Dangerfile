# Make sure if PR have assignee
failure "This PR does not have any assignees yet." unless github.pr_json["assignee"]

# Mention potential reviewer
mention.run

# Make sure one of the reviewer is from official reviewer
# Requested Reviewer
requested_reviewers = github.pr_json["requested_reviewers"]
# Actual Reviewer
pr_num = github.pr_json["number"]
reviews = github.api.pull_request_reviews("mpermperpisang/golang-automation", pr_num)
actual_reviewers = reviews.map {|u| u["user"]}
# PR Reviewer
reviewers = requested_reviewers + actual_reviewers
pr_reviewers = reviewers.map {|u| u["login"]}
# official Reviewer
official_reviewer = ["mpermperpisang", "mmpisang", "mpermper321"]

unless official_reviewer.any?{|x| pr_reviewers.include?(x)}
  failure "Please request a review from mpermperpisang or mmpisang or mpermper321"
end

# Make sure one of the approval is from official reviewer
list_approval = []

reviews.map {|u|
  if u["state"] == 'APPROVED'
    list_approval.push(u["user"]["login"])
  end
}

unless official_reviewer.any?{|x| list_approval.include?(x)}
  failure "Please get an approval from mmpisang"
end

# Provide automation running screenshot
formats = [".png", ".jpg", ".gif", ".jpeg"]

unless formats.any?{|x| github.pr_body.include?(x)}
  if git.modified_files.include? "features/*"
    failure "Modifying file must provide automation result screenshot!"
  end

  if git.added_files.include? "features/*"
    failure "Adding file must provide automation result screenshot!"
  end
end

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

if git.modified_files.include? "main.go"
  warn "Please assign @mpermperpisang or @mmpisang as reviewer"
end

# Looks Good To Me
lgtm.check_lgtm
