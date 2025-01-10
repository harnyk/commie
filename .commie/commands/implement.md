# Task: Implement a github actions workflow for goreleaser

Files:
 - SOURCE: .goreleaser.yml
 - TARGET .github/workflows/goreleaser.yml
 - GORELEASER_DOCUMENTATION: ../goreleaser/


Act step by step:
 - Examine the GORELEASER_DOCUMENTATION
 - Create the workflow file, which would trigger by a new tag and run a dedicated goreleaser job at TARGET, getting Github token from the corresponding secret