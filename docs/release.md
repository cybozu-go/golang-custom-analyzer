# Release procedure

This document describes how to release a new version.

## Labeling

Release notes are automatically generated based on PRs included in the release.
Those PRs are categorized based on the label assigned to them.

## Versioning

Follow [semantic versioning 2.0.0][semver] to choose the new version number.

## Bump version


1. Cut a new release branch from `main` branch.

    ```sh
    # Fetch the latest changes from main branch.
    git fetch origin main

    # Set VERSION and confirm it. It should NOT have "v" prefix.
    # The "v" prefix will be added automatically in the release tag.
    VERSION=x.y.z
    git switch -c release-v$VERSION origin/main
    ```

2. Update the version in `VERSION` file.

    ```sh
    echo v$VERSION > VERSION
    ```

3. Commit the change and push the release branch.

    ```sh
    git add VERSION
    git commit -m "Bump version to v$VERSION"
    git push origin release-v$VERSION
    ```

4. Create a pull request from the release branch to `main` branch and merge it.

    ```sh
    gh pr create --base main --head release-v$VERSION --title "Release v$VERSION" --body "Release v$VERSION"
    ```

5. After the pull request is merged, the CI workflow will automatically create the git tag and GitHub release.

[semver]: https://semver.org/spec/v2.0.0.html
