# triage.py

def triage_crashes(crash_reports):
    buckets = {}

    print(f"🚑 Triaging {len(crash_reports)} accidents...")

    for report in crash_reports:
        # We assume the "cause" is the last line of the error log
        # e.g., "Error: Segmentation fault at address 0x1234"
        cause = report.split(":")[-1].strip()

        if cause not in buckets:
            buckets[cause] = 0
            print(f"  🆕 Found NEW Bug cause: '{cause}'")
        
        buckets[cause] += 1

    print("\n📋 Final Report:")
    for cause, count in buckets.items():
        print(f"  - Bug '{cause}' happened {count} times.")

if __name__ == "__main__":
    # Let's try it!
    hospital_log = [
        "Crash 1: SegFault at 0x001", # The Banana
        "Crash 2: SegFault at 0x001", # The Banana
        "Crash 3: BufferOverflow at 0x999", # A different issue (maybe a wet floor?)
        "Crash 4: SegFault at 0x001"  # The Banana again
    ]

    triage_crashes(hospital_log)
