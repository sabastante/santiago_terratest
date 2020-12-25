# Technical Evaluation

## Documentation

This automation will allow the Engineer to automatically test terraform modules created in the repo.

In order to do that, they should:

1) Create a new folder within the S3_sample folder.
2) Create a new Terratest unittest within the test directory, targeting the module created in the previous task.

With that information, the Github Actions pipeline will:

1) Run a ubuntu instance that will launch Terratest and run the unittest against all existing modules.
2) If a test fail, the pipeline will fail and the merge will not be possible.

##References:
- https://terratest.gruntwork.io/docs/getting-started/quick-start/
- https://medium.com/@petriautero/automate-terraform-testing-with-github-actions-and-terratest-78d74331fdf8
