# minimizer.py
import hashlib

def get_file_hash(content):
    # Create a unique fingerprint (MD5) for the content
    return hashlib.md5(content.encode()).hexdigest()

def minimize_corpus(files_list):
    unique_fingerprints = set()
    clean_suitcase = []

    print(f"📦 Inspecting {len(files_list)} items...")

    for file_content in files_list:
        fingerprint = get_file_hash(file_content)

        if fingerprint in unique_fingerprints:
            print(f"  🗑️  Throwing away duplicate (Hash: {fingerprint[:6]}...)")
        else:
            print(f"  ✅ Keeping unique item (Hash: {fingerprint[:6]}...)")
            unique_fingerprints.add(fingerprint)
            clean_suitcase.append(file_content)

    return clean_suitcase

if __name__ == "__main__":
    # Let's try it!
    messy_room = ["Red Shirt", "Blue Shirt", "Red Shirt", "Green Shirt", "Red Shirt"]
    clean_pile = minimize_corpus(messy_room)

    print(f"\n✨ Final Suitcase: {clean_pile}")
