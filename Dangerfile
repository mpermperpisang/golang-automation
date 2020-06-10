# frozen_string_literal: true

# Welcome messages
welcome_message.greet

# Make sure if PR have assignee
failure 'This PR does not have any assignees yet.' unless github.pr_json['assignee']

# Suggest code changes through inline comments in pull requests
rubocop.lint(
  force_exclusion: true,
  inline_comment: true
)
system('bundle exec rubocop --auto-correct')

# Suggest code changes through inline comments in pull requests
suggester.suggest

# Make sure one of the reviewer is from official reviewer
# Requested reviewer
requested_reviewers = github.pr_json['requested_reviewers']
# Actual reviewer
repo = 'mpermperpisang/golang-automation'
pr_num = github.pr_json['number']
reviews = github.api.pull_request_reviews(repo, pr_num)
actual_reviewers = reviews.map { |u| u['user'] }
# PR reviewer
reviewers = requested_reviewers + actual_reviewers
pr_reviewers = reviewers.map { |u| u['login'] }
# Official reviewer
official_reviewer = %w[mpermperpisang mmpisang mpermper321]

# If reviewer not include official reviewer
unless official_reviewer.any? { |x| pr_reviewers.include?(x) }
  official_reviewer.delete(github.pr_author)
  review_requests.request(official_reviewer.sample(1))
end

# Make sure one of the approval is from official reviewer
list_approval = []

reviews.map { |u| list_approval.push(u['user']['login']) if u['state'] == 'APPROVED' }

unless official_reviewer.any? { |x| list_approval.include?(x) }
  failure 'Please get an approval from mpermperpisang or mmpisang or mpermper321'
end

# Provide automation running screenshot
formats = ['.png', '.jpg', '.gif', '.jpeg']

unless formats.any? { |x| github.pr_body.include?(x) }
  failure 'Modifying file must provide automation result screenshot!' if git.modified_files.include? 'features/*'
  failure 'Adding file must provide automation result screenshot!' if git.added_files.include? 'features/*'
end

# Make it more obvious that a PR is a work in progress and shouldn't be merged yet
warn 'PR is classed as Work in Progress' if github.pr_labels.include? 'WIP'
warn 'PR is classed as Work in Progress' if github.pr_title.include? '[WIP]'

# Warn when there is a big PR
warn('Big PR') if git.insertions > 500

# Provide PR summary
failure 'Please provide a summary in the Pull Request description' if github.pr_body.length < 5

# Warning to not fill in Todos description
if github.pr_body.include?('What things you do through this pull request')
  warn "Please provide a 'Todos' in the Pull Request description"
end

# Warning to not fill in Screenshot Link description
if github.pr_body.include?('Format file in png, jpg, gif, or jpeg')
  warn "Please provide a 'Screenshot' in the Pull Request description"
end

# Waning mergeable PR
can_merge = github.pr_json['mergeable']
warn('This PR cannot be merged yet.', sticky: false) unless can_merge

# Warning to assign reviewer if change a file
warn 'Please assign @mpermperpisang or @mmpisang or @mpermper321 as reviewer' if git.modified_files.include? 'main.go'

# Looks Good To Me
lgtm.check_lgtm

# Add specific label if approved and scored by official reviewer
label = 'to be crawled'
repo_label_list = github.api.labels(repo)
repo_label_name = repo_label_list.map { |u| u['name'] }
pr_label_list = github.api.labels_for_issue(repo, pr_num)
pr_label_name = pr_label_list.map { |u| u['name'] }
pr_comment_list = github.api.issue_comments(repo, pr_num)
pr_comment_body = pr_comment_list.map { |u| u['body'] }
info = 'After the PR merged, attach the run result within the pipeline/jenkins'\
" job and the link to it using comment with this syntax:\n"\
"```\n"\
"run_result\n"\
"link: <link>\n"\
"[!attach your screenshot]\n"\
'````'

# if official_reviewer.any? { |x| list_approval.include?(x) }
github.api.add_label(repo, label, 'C05472') unless repo_label_name.include?(label)

pr_comment_list.map do |u|
  github.api.delete_comment(repo, u['id']) if u['body'].include? 'After the PR merged'
end

if pr_comment_body.map(&:downcase).find { |e| /pr score/ =~ e }
  github.api.add_labels_to_an_issue(repo, pr_num, [label]) unless pr_label_name.include?(label)
  github.api.add_comment(repo, pr_num, info)
elsif pr_label_name.include?(label)
  github.api.remove_label(repo, pr_num, label)
end
# end
