# RELEASE

**Table of Content**

1. Preparations
2. Changelog
3. Version
4. Tests
5. Git Tag

## Preparations

Close the milestone on [GitHub](https://github.com/mcktr/check_fritz/milestones).

Specify the release version.

```
VERSION=1.1.0
```

## Changelog

Update the `CHANGELOG.md` file.

## Version

Update the version variable in `main.go`.

## Tests

* monitor CI tests in the `master` branch and ensure that they are passing.
* test locally if every check method works correctly.

## Git Tag

Ensure your local git repository is current and you're on `master` branch.

```
git pull
git checkout master
```

Commit these changes to the `master` branch.

```
git commit -v -a -m "Release version $VERSION"
```

Create a signed tag on the `master` branch.

```
git tag -s -m "Version $VERSION" v$VERSION
```

Push the tag.

```
git push --tags
```