import isctl

# from datetime import datetime, timezone, timedelta
import time


def gmtime(seconds):
    """Convert seconds since EPOCH to a dict with date/time components."""
    EPOCH_YEAR = 1970
    DAYS_IN_MONTH = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]

    def is_leap_year(year):
        return year % 4 == 0 and (year % 100 != 0 or year % 400 == 0)

    # Extract seconds, minutes, hours
    secs = seconds % 60
    mins = (seconds // 60) % 60
    hours = (seconds // 3600) % 24

    # Days since epoch
    days = seconds // 86400

    # Calculate year
    year = EPOCH_YEAR
    while True:
        days_in_year = 366 if is_leap_year(year) else 365
        if days < days_in_year:
            break
        days -= days_in_year
        year += 1

    # Calculate month and day
    month = 0
    for i, dim in enumerate(DAYS_IN_MONTH):
        days_in_month = dim
        if i == 1 and is_leap_year(year):  # February
            days_in_month = 29
        if days < days_in_month:
            month = i + 1
            break
        days -= days_in_month

    day = days + 1

    return {
        "year": year,
        "month": month,
        "day": int(day),
        "hour": int(hours),
        "minutes": int(mins),
        "seconds": int(secs),
    }


def isotime(dt):
    """Convert a datetime dict to ISO 8601 format string (YYYY-MM-DDTHH:MM:SS)."""
    return "%04d-%02d-%02dT%02d:%02d:%02d" % (
        dt["year"],
        dt["month"],
        dt["day"],
        dt["hour"],
        dt["minutes"],
        dt["seconds"],
    )


def cmd():
    return {
        "use": "telemetry",
        "short": "telemetry helper commands",
        "children": [
            {
                "use": "cpu",
                "short": "CPU telemetry commands",
                "children": [
                    {
                        "use": "summary MINUTES",
                        "short": "Get a summary of CPU usage",
                        "run": "cmd_telemetry_cpu_summary",
                        "args": {
                            "exact": 1,
                        },
                    }
                ],
            }
        ],
    }


def cmd_telemetry_cpu_summary(_args, _flags):
    isctl.info("Getting CPU summary")

    mins = int(_args[0])
    isctl.info("Getting summary for last %d minutes" % mins)

    end = time.time()
    start = end - mins * 60

    isctl.info(
        "Using time interval %s to %s" % (isotime(gmtime(start)), isotime(gmtime(end)))
    )

    t = isctl.executeOperation(
        "post",
        "telemetry.GroupBy",
        None,
        {
            "queryType": "groupBy",
            "dataSource": "PhysicalEntities",
            "dimensions": ["host.name"],
            "filter": {
                "type": "and",
                "fields": [
                    {
                        "type": "selector",
                        "dimension": "instrument.name",
                        "value": "hw.cpu",
                    }
                ],
            },
            "granularity": "all",
            "intervals": [
                # "2025-11-01T00:00:00.000+11:00/2025-12-01T00:00:00.000+11:00"
                isotime(gmtime(start))
                + "/"
                + isotime(gmtime(end))
            ],
            "aggregations": [
                {
                    "type": "longSum",
                    "name": "count",
                    "fieldName": "hw.cpu.utilization_c0_count",
                },
                {
                    "type": "doubleSum",
                    "name": "hw.cpu.utilization_c0-Sum",
                    "fieldName": "hw.cpu.utilization_c0",
                },
            ],
            "postAggregations": [
                {
                    "type": "expression",
                    "name": "intersight.hw.cpu.utilization_c0",
                    "expression": '("hw.cpu.utilization_c0-Sum" / "count")',
                }
            ],
        },
        None,
    )
    ret = []
    for i, e in enumerate(t):
        ret.append(
            {
                "host.name": e["event"]["host.name"],
                "hw.cpu.utilization_c0": "%.2f%%"
                % (e["event"]["intersight.hw.cpu.utilization_c0"] * 100),
            }
        )

    isctl.output(ret, False)
