# Create a good release tag.
Ask a user for the release type: major, minor, patch, pre-release etc.
Examine all the git tags.
Create a new **unsigned** tag according to the desired bump type.

(Important!) Ask my confirmation.
If I confirm, then create the tag.

Next you should push the commits and then tags.
(Important!) Ask my confirmation.
If I confirm, then push the commits and tags. Prefer `git push --tags` over `git push --follow-tags`, because the tags are unsigned.