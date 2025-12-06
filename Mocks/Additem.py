import requests
from faker import Faker
import random
import time

fake = Faker()

HOST = "localhost"
PORT = 8080
URL = f"http://{HOST}:{PORT}/additem"


def generate_mock_payload():
    payload = {
        "attributes": {
            "name": fake.word(),
            "description": fake.sentence(),
            "color": fake.color_name(),  # faker color
            "category": fake.word(),  # faker category
            "location": f"{fake.random_uppercase_letter()}{fake.random_digit()}{fake.random_digit()}",  # warehouse-like
            "unit_price": round(random.uniform(1.0, 500.0), 2),
            "stock_level": fake.random_int(min=0, max=200),
        },
        "supplier": {"name": fake.company(), "supplier_type": fake.word()},
    }

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
