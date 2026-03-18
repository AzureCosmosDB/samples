# Cosmos DB samples

One stop repo for all Cosmos DB samples.

## Repository Ruleset Recommendations

The repository uses a [GitHub Actions validation workflow](.github/workflows/validate.yml) that runs per-language linting, building, and testing on every PR and push to `main`. The sections below explain how to lock that workflow in as a required gate so that no code reaches `main` without passing CI.

---

### 1. Required Status Checks — workflow job names

The workflow defines the following job IDs. Use the **exact display names** (the `name:` field) when configuring required status checks:

| Job ID | Display name (use this in the ruleset) |
|---|---|
| `validate-structure` | `Validate Sample Structure` |
| `validate-python` | `Validate Python Samples` |
| `validate-javascript` | `Validate JavaScript/TypeScript Samples` |
| `validate-java` | `Validate Java Samples` |
| `validate-dotnet` | `Validate .NET Samples` |
| `validate-go` | `Validate Go Samples` |

At minimum, require **`Validate Sample Structure`** on every PR. Require the language-specific jobs for any language whose samples are considered stable/required.

---

### 2. Configuring a Repository Ruleset (recommended — GitHub Rulesets UI)

GitHub Rulesets are the modern replacement for classic Branch Protection Rules and support more granular controls.

1. Go to **Settings → Rules → Rulesets** in your repository.
2. Click **New ruleset → New branch ruleset**.
3. **Ruleset name**: `Protect main`
4. **Enforcement status**: `Active`
5. **Target branches**: Add rule → `Include by pattern` → enter `main`
6. Under **Rules**, enable the following:

   | Rule | Recommended setting |
   |---|---|
   | **Restrict deletions** | ✅ Enable — prevent `main` from being deleted |
   | **Require linear history** | ✅ Enable — enforces rebase/squash merge strategy, keeps history clean |
   | **Require a pull request before merging** | ✅ Enable |
   | ↳ Required approvals | `1` (increase to `2` for higher-risk repos) |
   | ↳ Dismiss stale pull request approvals when new commits are pushed | ✅ Enable |
   | ↳ Require review from Code Owners | ✅ Enable (add a `CODEOWNERS` file per language subfolder) |
   | ↳ Require approval of the most recent reviewable push | ✅ Enable |
   | **Require status checks to pass** | ✅ Enable |
   | ↳ Require branches to be up to date before merging | ✅ Enable |
   | ↳ Status checks that are required | Add each job name from the table above |
   | **Block force pushes** | ✅ Enable |

7. Click **Save changes**.

---

### 3. Classic Branch Protection Rules (alternative)

If your plan does not support Rulesets, use the classic Branch Protection Rules:

1. Go to **Settings → Branches**.
2. Under **Branch protection rules**, click **Add rule**.
3. **Branch name pattern**: `main`
4. Check the following options:

   - ✅ **Require a pull request before merging**
     - ✅ Require approvals: `1`
     - ✅ Dismiss stale pull request approvals when new commits are pushed
     - ✅ Require review from Code Owners
   - ✅ **Require status checks to pass before merging**
     - ✅ Require branches to be up to date before merging
     - Search for and add each status check name from the table in §1
   - ✅ **Require linear history**
   - ✅ **Do not allow bypassing the above settings** (prevents admins from force-merging)
   - ✅ **Restrict who can push to matching branches** (restrict to release automation if needed)

5. Click **Create** / **Save changes**.

---

### 4. `CODEOWNERS` recommendation

Create a `.github/CODEOWNERS` file to automatically request reviews from the correct team when language-specific files change. **Replace the team handles below with your actual GitHub team or user handles.**

```
# Global fallback
*                   @AzureCosmosDB/samples-maintainers

# Language owners
/python/            @AzureCosmosDB/python-sdk-team
/javascript/        @AzureCosmosDB/javascript-sdk-team
/java/              @AzureCosmosDB/java-sdk-team
/dotnet/            @AzureCosmosDB/dotnet-sdk-team
/go/                @AzureCosmosDB/go-sdk-team
```

---

### 5. Merge strategy recommendation

To keep history linear (required by the **Require linear history** ruleset rule above), configure the repository's allowed merge methods:

1. Go to **Settings → General → Pull Requests**.
2. Uncheck **Allow merge commits**.
3. Keep **Allow squash merging** ✅ (preferred for sample contributions) and/or **Allow rebase merging** ✅.

This ensures every commit on `main` maps directly to a reviewed PR.
