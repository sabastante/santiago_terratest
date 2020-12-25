# Technical Evaluation

## TL;DR

## Documentation

Nowadays it is critical to test infrastructure deployments. This might not be as intuitive as first as it is with software developments, but when the team is involved in high intensive agile developments, ensuring quality through automated testing **enables maintaining delivery speed**.

Besides that, it has been proved that the cost of not testing quickly becomes higher than the cost of implementing (automated) tests for infrastructure.

### Tools:

Terraform:Tool for building, changing, and versioning infrastructure safely and efficiently.
Terratest:  A swiss army knife for testing infrastructure code.
GitHub Actions:Create custom software development life cycle (SDLC) workflows directly in your GitHub repository.

### Usage

This automation will allow the Engineer to automatically test terraform modules created in the repo.

In order to do that, they should:

1) Create a new folder within the S3_sample folder.
2) Create a new Terratest unittest within the test directory, targeting the module created in the previous task.

With that information, the Github Actions pipeline will:

1) Run a ubuntu instance that will launch Terratest and run the unittest against all existing modules.
2) If a test fails, the pipeline will fail and the merge will not be possible.


## References:
- https://terratest.gruntwork.io/docs/getting-started/quick-start/
- https://medium.com/@petriautero/automate-terraform-testing-with-github-actions-and-terratest-78d74331fdf8
