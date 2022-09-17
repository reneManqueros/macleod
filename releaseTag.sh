git tag -a v1.0.0 -m "Initial public version"
git push origin v1.0.0
curl -sfL https://goreleaser.com/static/run | bash -s -- release