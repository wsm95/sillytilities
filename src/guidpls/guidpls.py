import uuid
import subprocess
import argparse

# Set up argument parser
parser = argparse.ArgumentParser(description="Generate and copy GUIDs to clipboard.")
parser.add_argument("count", type=int, nargs='?', default=1, help="Number of GUIDs to generate (default is 1)")
parser.add_argument("-n", "--nocopy", action="store_true", help="Do not copy GUIDs to clipboard")
args = parser.parse_args()

# Generate the specified number of GUIDs
guids = [str(uuid.uuid4()) for _ in range(args.count)]

# Concatenate the GUIDs into a comma-separated list
guids_str = ",".join(guids)

# Copy the GUIDs to the clipboard using the clip command if --nocopy is not specified
if not args.nocopy:
    subprocess.run("clip", universal_newlines=True, input=guids_str)

# Print each GUID on a separate line
for guid in guids:
    print(guid)