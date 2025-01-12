# Create a good release tag.

Step 1. Examine all the git tags.

Step 2. Ask a user for the release type (showing the new version number with each bump type):
 1. patch - VERSION_FOR_PATCH
 2. minor - VERSION_FOR_MINOR
 3. major - VERSION_FOR_MAJOR
 4. pre-release - VERSION_FOR_PRERELEASE
 5. other (please specify in the comment)

Create a new **unsigned** tag according to the desired bump type.

(Important!) Ask my confirmation.
If I confirm, then create the tag.

Next you should push the commits and then tags.
(Important!) Ask my confirmation.
If I confirm, then push the commits and tags. Prefer `git push --tags` over `git push --follow-tags`, because the tags are unsigned.