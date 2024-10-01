import base64
import json
import subprocess
import argparse

def decode_base64(data):
    # Add padding if necessary
    padding = '=' * (4 - len(data) % 4)
    return base64.urlsafe_b64decode(data + padding)

def get_jwt_token(token):
    if token:
        return token
    else:
        return subprocess.run(['powershell', '-command', 'Get-Clipboard'], capture_output=True, text=True).stdout.strip()

def decode_jwt(token):
    # Split the token into header, payload, and signature
    parts = token.split('.')
    if len(parts) != 3:
        print("Invalid JWT token")
        return
    
    # Decode the payload
    payload = parts[1]
    decoded_payload = decode_base64(payload)
    
    # Convert the decoded payload to a JSON object
    try:
        decoded_json = json.loads(decoded_payload)
        print(json.dumps(decoded_json, indent=4))
    except json.JSONDecodeError as e:
        print(f"Failed to decode JWT token: {e}")

# Set up argument parser
parser = argparse.ArgumentParser(description="Decode JWT tokens.")
parser.add_argument("token", type=str, nargs='?', help="JWT token to decode")
args = parser.parse_args()


token = get_jwt_token(args.token)
decode_jwt(token)