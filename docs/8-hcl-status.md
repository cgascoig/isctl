# HCL Status Report

The `isctl hcl status` command provides a consolidated Hardware Compatibility List (HCL) report for all servers managed by Intersight.

## Usage

```
isctl hcl status
```

The command queries the Intersight API for `cond.HclStatus` and `cond.HclStatusDetail` managed objects, correlates the details back to each server, and produces a single flat report.

## Output Fields

| Field | Description |
|---|---|
| `Status` | Overall HCL validation status (e.g. `Validated`, `Not-Listed`, `Incomplete`) |
| `ServerReason` | Reason for the current status |
| `HclModel` | Server model (e.g. `UCSC-C220-M5SX`) |
| `HclFirmwareVersion` | Firmware version validated against the HCL |
| `HclOsVendor` | OS vendor (e.g. `VMware`) |
| `HclOsVersion` | OS version (e.g. `ESXi 7.0 U3`) |
| `HclProcessor` | Processor family |
| `ComponentStatus` | Aggregate component validation status |
| `HardwareStatus` | Hardware validation status |
| `SoftwareStatus` | Software/driver validation status |
| `Details` | Per-component detail summary (semicolon-separated) |

## Examples

### Default output
```
isctl hcl status
```

The default output displays the report as a table with all fields. Because the table can be quite wide (especially the `Details` column), you may want to pipe it through a pager or use an alternative output format.

### Excel export
To export the HCL status report as an Excel spreadsheet:
```
isctl hcl status -o xlsx=hcl-report.xlsx
```

This creates an `hcl-report.xlsx` file that can be opened in Excel, Google Sheets, or any other spreadsheet application for further analysis and filtering.
