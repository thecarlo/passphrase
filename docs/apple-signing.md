# Apple Signing & Notarization Setup

This guide walks through getting every value needed to sign and notarize the `passphrase` binary for the Alfred Gallery.

**Prerequisites:** An active Apple Developer account at [developer.apple.com](https://developer.apple.com) ($99/year).

---

## 1. Find your Team ID → `APPLE_TEAM_ID`

1. Sign in to [developer.apple.com](https://developer.apple.com)
2. Click your name in the top-right → **Account**
3. Under **Membership Details**, copy the **Team ID** — it's a 10-character alphanumeric string (e.g. `AB12CD34EF`)

---

## 2. Create a Developer ID Application Certificate

This is the certificate used to sign the binary.

### Generate a Certificate Signing Request (CSR)

1. Open **Keychain Access** (Applications → Utilities → Keychain Access)
2. Menu bar: **Keychain Access → Certificate Assistant → Request a Certificate from a Certificate Authority**
3. Fill in:
   - **User Email Address:** your Apple ID email
   - **Common Name:** anything (e.g. `passphrase signing`)
   - **CA Email Address:** leave blank
   - Select **Saved to disk**
4. Click **Continue** and save the `.certSigningRequest` file somewhere handy

### Create the certificate on Apple's portal

1. Go to [developer.apple.com/account/resources/certificates/list](https://developer.apple.com/account/resources/certificates/list)
2. Click **+** to create a new certificate
3. Under **Software**, select **Developer ID Application**
4. Click **Continue**, upload your `.certSigningRequest` file
5. Download the resulting `.cer` file
6. Double-click it — this installs it into Keychain Access

---

## 3. Export the certificate as `.p12` → `APPLE_CERTIFICATE` + `APPLE_CERTIFICATE_PASSWORD`

1. Open **Keychain Access**
2. In the sidebar select **My Certificates**
3. Find the certificate named **Developer ID Application: Your Name (TEAMID)**
4. Right-click it → **Export**
5. Choose format **Personal Information Exchange (.p12)**
6. Save it as `certificate.p12`
7. Set a strong password when prompted — this becomes `APPLE_CERTIFICATE_PASSWORD`

Now base64-encode it for the GitHub secret:

```sh
base64 -i certificate.p12 | pbcopy
```

Paste the clipboard value as `APPLE_CERTIFICATE`.

---

## 4. Create an App Store Connect API Key → `APPLE_API_KEY_ID`, `APPLE_API_KEY_ISSUER_ID`, `APPLE_API_PRIVATE_KEY`

This is used by `notarytool` to authenticate with Apple.

1. Go to [appstoreconnect.apple.com](https://appstoreconnect.apple.com)
2. Navigate to **Users and Access → Integrations → App Store Connect API**
3. If this is your first key, click **Request Access** and agree to the terms
4. Click **+** to generate a new key
5. Give it a name (e.g. `notarization`) and set the role to **Developer**
6. Click **Generate**

From this page, copy:
- **Key ID** → `APPLE_API_KEY_ID` (e.g. `ABC123DEFG`)
- **Issuer ID** → `APPLE_API_KEY_ISSUER_ID` — shown at the top of the page above the key list (e.g. `xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx`)

Then download the `.p8` file by clicking **Download API Key**. 

> **Important:** You can only download the `.p8` file once. Store it somewhere safe.

Base64-encode it for the GitHub secret:

```sh
base64 -i AuthKey_XXXXXXXXXX.p8 | pbcopy
```

Paste the clipboard value as `APPLE_API_PRIVATE_KEY`.

---

## 5. Add the secrets to GitHub

1. Go to your repository on GitHub
2. **Settings → Secrets and variables → Actions → New repository secret**
3. Add each secret:

| Secret name | Where it came from |
|---|---|
| `APPLE_TEAM_ID` | Step 1 |
| `APPLE_CERTIFICATE` | Step 3 — base64 `.p12` |
| `APPLE_CERTIFICATE_PASSWORD` | Step 3 — password set on export |
| `APPLE_API_KEY_ID` | Step 4 — Key ID |
| `APPLE_API_KEY_ISSUER_ID` | Step 4 — Issuer ID |
| `APPLE_API_PRIVATE_KEY` | Step 4 — base64 `.p8` |

---

## 6. Verify signing locally (optional)

After a release build, you can confirm the binary is correctly signed:

```sh
# Check signature
codesign --verify --verbose workflow/passphrase

# Check it passes Gatekeeper
spctl --assess --verbose workflow/passphrase
```

Both should exit without errors. `spctl` should print `accepted`.
