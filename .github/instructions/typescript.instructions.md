---
applyTo: "typescript/**"
---

# TypeScript sample conventions

## File structure

- `package.json` with dependencies and scripts (`start`, `build`, `test`)
- `tsconfig.json` with ES2022 target, Node module resolution, strict mode
- Source `.ts` file as entry point
- Use `tsx` for direct TypeScript execution (no compile step to run)
- `.gitignore` for `node_modules/`, `dist/`, `.env`

## Dependencies

- Use caret ranges (`^`) in `package.json` — do not pin exact versions
- Primary SDK: `@azure/cosmos`
- Node.js 20+

## Authentication

- Key-based: `COSMOS_ENDPOINT` + `COSMOS_KEY` environment variables
- Read with `process.env.COSMOS_ENDPOINT` — throw if missing
- Do **not** use `dotenv` — put env var setup instructions in the readme

## Code style

- Use `interface` for the `Product` type (not `type` alias or `class`)
- Use `async`/`await` throughout
- Use ES module syntax (`import`/`export`)
