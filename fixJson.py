import json

with open('items.json', 'r', encoding='utf-8') as file:
    data = json.load(file)

def convert_price_to_number(price_str):
    return int(price_str.replace('₽', '').replace(' ', ''))

for item in data:
    item["price"] = convert_price_to_number(item["price"])

with open('fixed_items.json', 'w', encoding='utf-8') as file:
    json.dump(data, file, ensure_ascii=False, indent=2)
