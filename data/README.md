# Data Files

This directory contains data files used by the LeetCode project.

## 📁 Files

### `problem_categories.csv`
Original problem categorization data used during the Jan 2026 reorganization.

**Columns:**
- `problem_number`: LeetCode problem number (e.g., "0001")
- `problem_name`: Problem name in snake_case (e.g., "two_sum")
- `category`: Algorithm category (e.g., "arrays", "strings", "dp")
- `difficulty`: Problem difficulty (Easy, Medium, Hard)
- `notes`: Additional notes

### `problem_categories_corrected.csv`
Corrected problem categorization data after review and validation.

**Columns:** Same as `problem_categories.csv` with corrected categories.

## 🔧 Usage

These CSV files were used by the reorganization scripts to:
1. Categorize problems by algorithm type
2. Move files to appropriate directories
3. Validate the reorganization

## 📊 Data Sources

The categorization was based on:
- LeetCode problem tags
- Algorithm patterns
- Common problem classifications
- Manual review and correction

## 🔄 Maintenance

When updating problem categorizations:
1. Update the CSV files
2. Run the reorganization scripts if needed
3. Verify file placements match categories

---

*Last updated: January 2026*