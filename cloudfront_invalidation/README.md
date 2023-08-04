# CloudFront Invalidation Action

This is a custom action for GitHub Actions that allows you to trigger a CloudFront invalidation for the specified paths in a CloudFront distribution.

## Usage

To use this custom action, you can include it in your GitHub Actions workflows in your repository. Here's an example of how you can do it:

``` yml
- name: Run CloudFront Invalidation Action
  uses: ./path/to/cloudfront-invalidation-action
  with:
    cloudfront_distribution_id: your-distribution-id
    invalidation_paths: /path1,/path2
    aws_access_key_id: ${{ secrets.AWS_ACCESS_KEY_ID }}
    aws_secret_access_key: ${{ secrets.AWS_SECRET_ACCESS_KEY }}

```

Make sure to replace your-distribution-id with the ID of your CloudFront distribution, and provide your AWS access keys in the repository secrets.

## Inputs
* cloudfront_distribution_id (required): The ID of the CloudFront distribution you want to invalidate.
* invalidation_paths (required): The paths to invalidate, separated by commas (e.g., /path1,/path2).
* aws_access_key_id (required): AWS access key ID for authentication.
* aws_secret_access_key (required): AWS secret access key for authentication.