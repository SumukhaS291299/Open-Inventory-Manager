import base64
import datetime
import requests
from faker import Faker
import random
import time

fake = Faker()

HOST = "localhost"
PORT = 8080
URL = f"http://{HOST}:{PORT}/additem"


def fake_base64_image():
    return base64.b64encode(fake.binary(length=64)).decode()


def iso_rfc3339():
    return datetime.datetime.now(datetime.timezone.utc).isoformat()


def generate_mock_payload():
    payload = {
        "attributes": {
            "name": fake.word(),
            "description": fake.sentence(),
            "color": fake.color_name(),
            "category": fake.word(),
            "unit_price": round(random.uniform(1.0, 500.0), 2),
            "stock_level": random.randint(0, 500),
            "location": f"{fake.random_uppercase_letter()}{random.randint(1, 99)}",
            "is_active": random.choice([True, False]),
            "is_available": random.choice([True, False]),
            "photo_base64": fake_base64_image(),
            "tags": [
                {"id": fake.uuid4(), "name": fake.word()}
                for _ in range(random.randint(1, 3))
            ],
        },
        "time_meta": {
            "bought": iso_rfc3339(),
            "expires": iso_rfc3339(),
            "modified": iso_rfc3339(),
        },
    }

    # Optional supplier
    if random.choice([True, False]):
        payload["supplier"] = {
            "supplier_type": fake.word(),
            "name": fake.company(),
            "online": random.choice([True, False]),
            "address": fake.address(),
        }

    # Optional comments
    if random.choice([True, False]):
        payload["comments"] = [
            {
                "id": fake.uuid4(),
                "miscellaneous": {
                    "note": fake.sentence(),
                    "severity": random.randint(1, 5),
                },
                "created_at": iso_rfc3339(),
                "created_by": fake.user_name(),
            }
        ]

    # Optional global tags
    if random.choice([True, False]):
        payload["tags"] = [
            {"id": fake.uuid4(), "name": fake.word()}
            for _ in range(random.randint(1, 3))
        ]

    return payload


def send_mock_item():
    payload = generate_mock_payload()
    print("Sending:", payload)

    response = requests.post(URL, json=payload)
    print("Status:", response.status_code)
    print("Response:", response.text)
    print("-" * 40)


if __name__ == "__main__":
    for _ in range(20000):
        send_mock_item()
        time.sleep(0.1)
