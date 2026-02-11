import isctl


def cmd():
    return {
        "use": "report",
        "short": "Report commands",
        "children": [
            {
                "use": "hcl",
                "short": "Display HCL status for all servers",
                "run": "cmd_report_hcl",
            },
            {
                "use": "contract-status",
                "short": "Display contract status for all devices",
                "run": "cmd_report_contract_status",
            },
        ],
    }


def get_str(d, key):
    """Get a string value from a dict, returning empty string if missing or None."""
    if key not in d:
        return ""
    val = d[key]
    if val is None:
        return ""
    return val


def get_nested(d, key1, key2):
    """Get a nested value from a dict, e.g. d[key1][key2], returning empty string if missing."""
    if key1 not in d:
        return ""
    val = d[key1]
    if val is None:
        return ""
    if key2 not in val:
        return ""
    nested = val[key2]
    if nested is None:
        return ""
    return nested


def join_parts(parts, sep):
    """Join a list of strings with a separator (gpython doesn't support str.join)."""
    result = ""
    for i in range(len(parts)):
        if i > 0:
            result = result + sep
        result = result + parts[i]
    return result


# ---- HCL Status Report ----


def cmd_report_hcl(_args, _flags):
    isctl.info("Collecting HCL status information")

    # Get all cond.HclStatus MOs
    isctl.info("Querying cond.HclStatus")
    hcl_statuses_res = isctl.executeOperation(
        "get", "cond.HclStatus", None, None, None, 1000
    )
    if hcl_statuses_res is None or "Results" not in hcl_statuses_res:
        isctl.error("Error getting cond.HclStatus")
        return
    hcl_statuses = hcl_statuses_res["Results"]
    isctl.info("Got %s HclStatus entries" % len(hcl_statuses))

    # Get all cond.HclStatusDetail MOs
    isctl.info("Querying cond.HclStatusDetail")
    hcl_details_res = isctl.executeOperation(
        "get", "cond.HclStatusDetail", None, None, None, 1000
    )
    if hcl_details_res is None or "Results" not in hcl_details_res:
        isctl.error("Error getting cond.HclStatusDetail")
        return
    hcl_details = hcl_details_res["Results"]
    isctl.info("Got %s HclStatusDetail entries" % len(hcl_details))

    # Index all HclStatusDetail entries by their Moid
    details_by_moid = {}
    for i in range(len(hcl_details)):
        detail = hcl_details[i]
        detail_moid = get_str(detail, "Moid")
        if detail_moid != "":
            details_by_moid[detail_moid] = detail
    isctl.info("Indexed %s detail entries" % len(details_by_moid))

    # Build the output list
    result = []
    for i in range(len(hcl_statuses)):
        status = hcl_statuses[i]

        # Build details summary string by looking up the Details MoRef list
        detail_parts = []
        if "Details" in status:
            refs = status["Details"]
            if refs is not None:
                for j in range(len(refs)):
                    ref = refs[j]
                    ref_moid = get_str(ref, "Moid")
                    if ref_moid != "" and ref_moid in details_by_moid:
                        d = details_by_moid[ref_moid]
                        # Component is a MoRef, extract ObjectType as the label
                        component = ""
                        if "Component" in d and d["Component"] is not None:
                            comp = d["Component"]
                            if "ObjectType" in comp:
                                component = comp["ObjectType"]
                        d_status = get_str(d, "Status")
                        hw = get_str(d, "HardwareStatus")
                        sw = get_str(d, "SoftwareStatus")
                        reason = get_str(d, "Reason")
                        s = "%s: status=%s" % (component, d_status)
                        if hw != "":
                            s = s + " hw=%s" % hw
                        if sw != "":
                            s = s + " sw=%s" % sw
                        if reason != "" and reason != "Missing-Os-Info":
                            s = s + " (%s)" % reason
                        detail_parts.append(s)

        details_str = join_parts(detail_parts, "; ")

        entry = {}
        entry["Status"] = get_str(status, "Status")
        entry["ServerReason"] = get_str(status, "ServerReason")
        entry["HclFirmwareVersion"] = get_str(status, "HclFirmwareVersion")
        entry["HclOsVendor"] = get_str(status, "HclOsVendor")
        entry["HclOsVersion"] = get_str(status, "HclOsVersion")
        entry["HclModel"] = get_str(status, "HclModel")
        entry["HclProcessor"] = get_str(status, "HclProcessor")
        entry["ComponentStatus"] = get_str(status, "ComponentStatus")
        entry["HardwareStatus"] = get_str(status, "HardwareStatus")
        entry["SoftwareStatus"] = get_str(status, "SoftwareStatus")
        entry["Details"] = details_str
        result.append(entry)

    isctl.info("Built %s result entries" % len(result))
    isctl.output(result, False)


# ---- Contract Status Report ----


def cmd_report_contract_status(_args, _flags):
    isctl.info("Collecting contract status information")

    res = isctl.executeOperation(
        "get", "asset.DeviceContractInformation", None, None, None, 1000
    )
    if res is None or "Results" not in res:
        isctl.error("Error getting asset.DeviceContractInformation")
        return

    contracts = res["Results"]
    isctl.info("Got %s contract entries" % len(contracts))

    result = []
    for i in range(len(contracts)):
        contract = contracts[i]
        entry = {}
        entry["ContractStatus"] = get_str(contract, "ContractStatus")
        entry["ContractStatusReason"] = get_str(contract, "ContractStatusReason")
        entry["DeviceId"] = get_str(contract, "DeviceId")
        entry["DeviceType"] = get_str(contract, "DeviceType")
        entry["ItemType"] = get_str(contract, "ItemType")
        entry["PlatformType"] = get_str(contract, "PlatformType")
        entry["ProductNumber"] = get_nested(contract, "Product", "Number")
        result.append(entry)

    isctl.output(result, False)
