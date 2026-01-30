# Scripts Directory

This directory contains utility scripts for managing the LeetCode project.

## 📁 Directory Structure

```
scripts/
├── README.md              # This file
├── reorganization/        # Reorganization scripts
├── fixes/                 # Fix scripts
└── analysis/              # Analysis scripts (future)
```

## 🔧 Script Categories

### Reorganization Scripts (`reorganization/`)
Scripts used for the recent project reorganization (Jan 2026).

| Script | Purpose |
|--------|---------|
| `categorize_problems.py` | Categorize LeetCode problems by algorithm type |
| `correct_categories.py` | Correct problem categorizations |
| `move_files.py` | Move files to categorized directories |
| `review_categories.py` | Review and validate categories |
| `verify_categories.py` | Verify file placements |

### Fix Scripts (`fixes/`)
Scripts for fixing various issues in the codebase.

| Script | Purpose |
|--------|---------|
| `fix_func_syntax.py` | Fix function syntax issues |
| `fix_imports.py` | Fix import statements |
| `fix_remaining_issues.py` | Fix remaining issues after reorganization |
| `fix_test_files.py` | Fix test file issues |
| `fix_test_files_v2.py` | Updated test file fixes |
| `fix_utils_imports.py` | Fix utils import issues |
| `update_imports.py` | Update import paths |
| `update_packages.py` | Update package declarations |

## 🚀 Usage

Most scripts can be run directly:

```bash
# Run a reorganization script
python scripts/reorganization/categorize_problems.py

# Run a fix script
python scripts/fixes/fix_imports.py
```

## 📝 Notes

- These scripts were primarily used during the Jan 2026 reorganization
- Some scripts may be project-specific and not general purpose
- Always review script output before applying changes

## 🔄 Maintenance

When adding new scripts:
1. Place them in the appropriate category directory
2. Update this README with script description
3. Add usage examples if applicable

---

*Last updated: January 2026*