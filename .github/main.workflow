workflow "CI" {
  on = "push"
  resolves = ["Test"]
}

action "GolangCI-Lint" {
  uses = "./.github/actions/ci"
  args = "lint"
}

action "Test" {
  needs = ["GolangCI-Lint"]
  uses = "./.github/actions/ci"
  args = "test"
}