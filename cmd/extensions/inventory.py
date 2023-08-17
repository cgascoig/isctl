import isctl


def cmd():
    return {
        "use": "inventory",
        "short": "Inventory helper commands",
        "children": [
            {
                "use": "compute",
                "short": "Display a summary of compute inventory",
                "run": "cmd_inventory_compute",
            }
        ],
    }


def cmd_inventory_compute(_args, _flags):
    isctl.info("Collecting compute inventory")

    ret = {}
    class_ids = [
        "compute.PhysicalSummary",
        "processor.Unit",
        "memory.Unit",
        "storage.Controller",
        "storage.PhysicalDisk",
        "pci.Device",
    ]

    for class_id in class_ids:
        isctl.info("Collecting %s inventory" % class_id)
        res = isctl.executeOperation("get", class_id, None, None, None, 1000)
        if res is None or "Results" not in res:
            isctl.error("Error getting %s(s)" % class_id)
            return
        ret[class_id] = res["Results"]

    # Create a map of (class_id, moid)->Serial for all the compute nodes
    # ...
    serial_map = {}
    for comp in ret["compute.PhysicalSummary"]:
        # gpython doesn't support tuple keys
        # serial_map[(comp["ClassId"], comp["Moid"])] = comp["Serial"]
        serial_map["%s:%s" % (comp["SourceObjectType"], comp["Moid"])] = comp["Serial"]

    # Go through all the results for other class_ids and find any ancestors that match and add the compute serial number
    for class_id in class_ids[1:]:
        for i, mo in enumerate(ret[class_id]):
            for ancestor in mo["Ancestors"]:
                serial = serial_map.get(
                    "%s:%s" % (ancestor["ObjectType"], ancestor["Moid"])
                )
                if serial is not None:
                    ret[class_id][i]["<Compute_Serial>"] = serial

    isctl.output(ret, True)
