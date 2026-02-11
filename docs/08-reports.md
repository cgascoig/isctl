# Reports

The `isctl report` command provides a set of pre-built reports that query Intersight for common operational data.

## HCL Status Report

The `isctl report hcl` command provides a consolidated Hardware Compatibility List (HCL) report for all servers managed by Intersight.

### Usage

```
isctl report hcl
```

The command queries the Intersight API for `cond.HclStatus` and `cond.HclStatusDetail` managed objects, correlates the details back to each server, and produces a single flat report.

### Output Fields

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

### Examples

```
isctl report hcl
```

To export as an Excel spreadsheet:
```
isctl report hcl -o xlsx=hcl-report.xlsx
```

---

## Contract Status Report

The `isctl report contract-status` command displays the contract status for all devices managed by Intersight.

### Usage

```
isctl report contract-status
```

The command queries the Intersight API for `asset.DeviceContractInformation` managed objects and produces a report of contract coverage.

### Output Fields

| Field | Description |
|---|---|
| `ContractStatus` | Current contract status (e.g. `Active`, `Expired`) |
| `ContractStatusReason` | Reason for the current contract status |
| `DeviceId` | Device identifier (e.g. serial number) |
| `DeviceType` | Type of device |
| `ItemType` | Item type classification |
| `PlatformType` | Platform type (e.g. `UCSFI`, `IMC`) |
| `ProductNumber` | Product number (PID) |

### Examples

```
isctl report contract-status
```

To export as an Excel spreadsheet:
```
isctl report contract-status -o xlsx=contract-status.xlsx
```
