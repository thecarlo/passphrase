# Development

## File structure

```
passphrase/
├── main.go                              # entry point — calls wf.Run
├── workflow.go                          # package-level wf var + init
├── run.go                               # Alfred handler — reads config, builds items, sends feedback
├── suggestions.go                       # assembles the 5 passphrase variants
├── words.go                             # picks and capitalizes N random words
├── passphrase.go                        # joins words with hyphens
├── insert.go                            # inserts a number at an encoded position
├── number.go                            # returns a random int 1–100
├── capitalize.go                        # uppercases the first character
└── internal/
    └── diceware/
        ├── wordlist.go                  # embeds wordlist, exposes RandomWords(n)
        └── eff_large_wordlist.txt       # EFF large wordlist (7,776 words), embedded at compile time
```

`main` package files live at the root since this is a single binary with no exportable packages. Reusable logic is isolated in `internal/diceware/`.

## Running tests

Run all tests:

```sh
go test ./...
```

Run with verbose output to see each test name:

```sh
go test -v ./...
```

Force a fresh run (bypass cache):

```sh
go test -count=1 -v ./...
```

## Releasing a new version

1. Ensure all changes are committed and pushed to `master`.

2. Create a version tag:
   ```sh
   git tag v1.2.3
   ```

3. Push the tag to GitHub:
   ```sh
   git push origin v1.2.3
   ```

This triggers the [release workflow](.github/workflows/release.yml), which:
- Runs tests
- Builds a universal binary (amd64 + arm64 via `lipo`)
- Signs and notarizes the binary using your Apple Developer credentials
- Packages it as `passphrase.alfredworkflow`
- Creates a GitHub Release with auto-generated release notes and the workflow attached

## Setting up the workflow in Alfred

1. Build the binary:
   ```sh
   go build -o passphrase .
   ```

2. In Alfred Preferences, create a new blank workflow, then add a **Script Filter** input with these settings:
   - **Keyword:** `pw`
   - **Argument:** Argument Optional, with input as argv
   - **Script:** `./passphrase "$1"`

3. Connect the Script Filter output to a **Copy to Clipboard** action.

4. Place the compiled `passphrase` binary in the workflow's folder.

## Updating the binary

After making code changes, run `build.sh` to compile a universal binary (Intel + Apple Silicon):

```sh
./build.sh
```

This builds for both `amd64` and `arm64` and combines them into a single `passphrase` binary using `lipo`.

Then copy the binary into the Alfred workflow folder:

```sh
cp passphrase "/path/to/Alfred.alfredpreferences/workflows/user.workflow.XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX/passphrase"
```

To find your workflow folder, right-click the workflow in Alfred Preferences and select **Open in Finder**.
